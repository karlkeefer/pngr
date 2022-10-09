package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/karlkeefer/pngr/golang/db"
	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/server/jwt"
	"github.com/karlkeefer/pngr/golang/server/write"
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

func Signup(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	u := &db.User{}
	err := decoder.Decode(&u)
	if err != nil || u == nil {
		return write.Error(errors.NoJSONBody)
	}

	// salt and hash it
	u.Salt = generateRandomString(32)
	u.Pass, err = hashPassword(u.Pass, u.Salt)
	if err != nil {
		return write.Error(err)
	}

	dbUser, err := env.DB().CreateUser(r.Context(), db.CreateUserParams{
		Email:        u.Email,
		Pass:         u.Pass,
		Salt:         u.Salt,
		Status:       db.UserStatusUnverified,
		Verification: generateRandomString(32),
	})

	if err != nil {
		if isDupe(err) {
			return write.Error(errors.AlreadyRegistered)
		}
		return write.Error(err)
	}

	err = env.Mailer().VerifyEmail(dbUser.Email, dbUser.Verification)
	if err != nil {
		return write.Error(err)
	}

	return write.Success()
}

func UpdatePassword(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	if user.Status != db.UserStatusActive {
		return write.Error(errors.RouteUnauthorized)
	}

	decoder := json.NewDecoder(r.Body)
	u := &db.User{}
	err := decoder.Decode(&u)
	if err != nil || u == nil {
		return write.Error(errors.NoJSONBody)
	}

	// salt and hash it
	u.Salt = generateRandomString(32)
	u.Pass, err = hashPassword(u.Pass, u.Salt)
	if err != nil {
		return write.Error(err)
	}

	err = env.DB().UpdateUserPassword(r.Context(), db.UpdateUserPasswordParams{
		ID:   user.ID,
		Pass: u.Pass,
		Salt: u.Salt,
	})

	if err != nil {
		return write.Error(err)
	}

	return write.Success()
}

func Whoami(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	return write.JSON(user)
}

type verifyRequest struct {
	Code string
}

func Verify(env env.Env, user *db.User, w http.ResponseWriter, r *http.Request) http.HandlerFunc {
	decoder := json.NewDecoder(r.Body)
	req := &verifyRequest{}

	err := decoder.Decode(&req)
	if err != nil || req == nil || req.Code == "" {
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
		ID:     u.ID,
		Status: u.Status,
	})

	if err != nil {
		return write.Error(err)
	}

	jwt.WriteUserCookie(w, &u)
	return write.JSON(&u)
}
