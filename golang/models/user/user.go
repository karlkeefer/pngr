package user

import (
	"github.com/jmoiron/sqlx"
	"github.com/karlkeefer/pngr/golang/errors"

	"golang.org/x/crypto/bcrypt"

	"encoding/json"
	"math/rand"
	"time"
)

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
	_, err := r.FindByEmail(u.Email)
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

func (r *Repo) Authenticate(u *User) (fromDB *User, err error) {
	fromDB, err = r.FindByEmail(u.Email)
	if err != nil {
		return
	}

	if !checkPasswordHash(u.Pass, fromDB.Pass) {
		err = errors.FailedLogin
	}

	return
}

func (r *Repo) FindByEmail(e string) (*User, error) {
	var u User

	err := r.db.Get(&u, `SELECT * FROM users WHERE email = $1 LIMIT 1`, e)
	if err != nil {
		return nil, errors.UserNotFound
	}

	return &u, nil
}

func (r *Repo) Verify(code string) (u *User, err error) {

	err = r.db.Get(&u, `SELECT * FROM users WHERE verification = $1 LIMIT 1`, code)
	if err != nil {
		err = errors.VerificationNotFound
		return
	}

	if u.Status != UserStatusUnverified {
		err = errors.VerificationExpired
		return
	}

	// update status
	u.Status = UserStatusActive

	err = r.UpdateStatus(u)
	if err != nil {
		return
	}

	return
}

func (r *Repo) UpdateStatus(u *User) (err error) {
	_, err = r.db.Exec(`UPDATE users SET status = $1 WHERE id = $2`,
		u.Status, u.ID)

	return
}
