package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/karlkeefer/pngr/golang/db"
	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/handlers/jwt"
	"github.com/karlkeefer/pngr/golang/handlers/write"
	"golang.org/x/crypto/bcrypt"
)

var src = rand.NewSource(time.Now().UnixNano())

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// generateRandomString can be used to create verification codes or something this
// this implementation comes from stackoverflow
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func generateRandomString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func hashPassword(password, salt string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+salt), 14)
	return string(bytes), err
}

func checkPasswordHash(password, salt, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
	return err == nil
}

type signupResponse struct {
	URL string
}

func Signup(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	var u db.User
	err := decoder.Decode(&u)
	if err != nil || &u == nil {
		return write.Error(errors.NoJSONBody)
	}

	// salt and hash it
	u.Salt = generateRandomString(32)
	u.Pass, err = hashPassword(u.Pass, u.Salt)
	if err != nil {
		return write.Error(err)
	}

	dbUser, err := env.DB().CreateUser(r.Context(), db.CreateUserParams{
		Lower:        u.Email,
		Pass:         u.Pass,
		Salt:         u.Salt,
		Status:       db.UserStatusUnverified,
		Verification: generateRandomString(32),
	})
	if err != nil {
		return write.Error(err)
	}

	// TODO: this is where we should actually email the code with mailgun or something
	// for now we just pass verification code back in the response...
	return write.JSON(&signupResponse{
		URL: fmt.Sprintf("%s/verify/%s", os.Getenv("APP_ROOT"), dbUser.Verification),
	})
}

func Whoami(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return write.JSON(user)
}

type verifyRequest struct {
	Code string
}

func Verify(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	var req verifyRequest

	err := decoder.Decode(&req)
	if err != nil || &req == nil || req.Code == "" {
		return write.Error(errors.NoJSONBody)
	}

	u, err := env.DB().FindUserByVerificationCode(r.Context(), req.Code)
	if err != nil {
		return write.Error(err)
	}

	if u.Status != db.UserStatusUnverified {
		return write.Error(errors.VerificationExpired)
	}

	u.Status = db.UserStatusActive

	err = env.DB().UpdateUserStatus(r.Context(), db.UpdateUserStatusParams{
		ID:        u.ID,
		Status:    u.Status,
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return write.Error(err)
	}

	jwt.WriteUserCookie(w, &u)
	return write.JSON(u)
}
