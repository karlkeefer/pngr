package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/karlkeefer/pngr/golang/errors"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"encoding/json"
	"math/rand"
	"os"
	"time"
)

var hmacSecret []byte

func init() {
	hmacSecret = []byte(os.Getenv("TOKEN_SECRET"))
	if hmacSecret == nil {
		panic("No TOKEN_SECRET environment variable was found")
	}
}

type User struct {
	ID           int64
	Name         *string // nullable
	Email        string
	Pass         string
	Status       int
	Verification string
	Created      time.Time
}

// protect private fields from being sent *out*
func (u User) MarshalJSON() ([]byte, error) {
	var tmp struct {
		ID      int64
		Name    *string // nullable
		Email   string
		Status  int
		Created time.Time
	}

	tmp.ID = u.ID
	tmp.Name = u.Name
	tmp.Email = u.Email
	tmp.Status = u.Status
	tmp.Created = u.Created
	return json.Marshal(&tmp)
}

const (
	UserStatusUnverified = 0
	UserStatusActive     = 1
	UserStatusDisabled   = 2
)

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

type Repo struct {
	db *sqlx.DB
}

// generate random verification codes
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
var src = rand.NewSource(time.Now().UnixNano())

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

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

func (r *Repo) Signup(u *User) error {
	err, _ := r.FindByEmail(u.Email)
	if err != errors.UserNotFound {
		return errors.InvalidEmail
	}

	// set verification code
	u.Verification = generateRandomString(32)

	// hash the password
	u.Pass, err = hashPassword(u.Pass)

	if err != nil {
		return err
	}

	_, err = r.db.Exec(`INSERT INTO users (email, pass, status, verification) VALUES ($1, $2, $3, $4)`,
		u.Email, u.Pass, UserStatusUnverified, u.Verification)

	return err
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type Auth struct {
	JWT string
}

func (r *Repo) Authenticate(u *User) (error, *Auth) {
	err, fromDB := r.FindByEmail(u.Email)
	if err != nil {
		return err, nil
	}

	if !checkPasswordHash(u.Pass, fromDB.Pass) {
		return errors.FailedLogin, nil
	}

	u = fromDB

	return buildAuth(u)
}

func buildAuth(u *User) (error, *Auth) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": jwt.MapClaims{
			"ID":      u.ID,
			"Name":    u.Name,
			"Email":   u.Email,
			"Status":  u.Status,
			"Created": u.Created,
		},
		// TODO: setup appropriate JWT values time-related claims
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return errors.InternalError, nil
	}

	return nil, &Auth{
		JWT: tokenString,
	}
}

func (r *Repo) FindByEmail(e string) (error, *User) {
	var u User

	err := r.db.Get(&u, `SELECT * FROM users WHERE email = $1 LIMIT 1`, e)
	if err != nil {
		return errors.UserNotFound, nil
	}

	return nil, &u
}

func (r *Repo) Verify(code string) (error, *Auth) {
	var u User

	err := r.db.Get(&u, `SELECT * FROM users WHERE verification = $1 LIMIT 1`, code)
	if err != nil {
		return errors.VerificationNotFound, nil
	}

	if u.Status != UserStatusUnverified {
		return errors.VerificationExpired, nil
	}

	// update status
	u.Status = UserStatusActive

	err = r.UpdateStatus(&u)
	if err != nil {
		return err, nil
	}

	return buildAuth(&u)
}

func (r *Repo) UpdateStatus(u *User) (err error) {
	_, err = r.db.Exec(`UPDATE users SET status = $1 WHERE id = $2`,
		u.Status, u.ID)

	return
}
