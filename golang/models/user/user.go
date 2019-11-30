package user

import (
	"encoding/json"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int64
	Name         *string // nullable
	Email        string
	Salt         string
	Pass         string
	Status       Status
	Verification string
	Created      time.Time
}

// Status is like a role, including unverified
type Status int

const (
	StatusDisabled   Status = -1
	StatusUnverified        = 0
	StatusActive            = 1
	StatusAdmin             = 10
)

// MarshalJSON here protects "private" fields from ever being sent *out*
// it also makes Name return "" instead of null
func (u User) MarshalJSON() ([]byte, error) {
	var tmp struct {
		ID      int64      `json:"id"`
		Name    string     `json:"name"`
		Email   string     `json:"email,omitempty"`
		Status  Status     `json:"status"`
		Created *time.Time `json:"created,omitempty"`
	}

	tmp.ID = u.ID

	// pick a name
	if u.Name != nil {
		tmp.Name = *u.Name
	} else {
		tmp.Name = ""
	}

	tmp.Email = u.Email
	tmp.Status = u.Status
	if !u.Created.IsZero() {
		tmp.Created = &u.Created
	}
	return json.Marshal(&tmp)
}

// REPO stuff
// TODO: consider moving repo to separate package
type Repo struct {
	signup       *sqlx.NamedStmt
	updateStatus *sqlx.NamedStmt
	findByEmail  *sqlx.Stmt
	verify       *sqlx.Stmt
}

func NewRepo(db *sqlx.DB) *Repo {
	// TODO: create a helper to prepare named and regular statements
	signup, err := db.PrepareNamed(`INSERT INTO users (email, salt, pass, status, verification) VALUES (:email, :salt, :pass, :status, :verification) RETURNING *`)
	if err != nil {
		panic(err)
	}
	updateStatus, err := db.PrepareNamed(`UPDATE users SET status = :status WHERE id = :id`)
	if err != nil {
		panic(err)
	}
	findByEmail, err := db.Preparex(`SELECT * FROM users WHERE email = $1 LIMIT 1`)
	if err != nil {
		panic(err)
	}
	verify, err := db.Preparex(`SELECT * FROM users WHERE verification = $1 LIMIT 1`)
	if err != nil {
		panic(err)
	}
	return &Repo{
		signup,
		updateStatus,
		findByEmail,
		verify,
	}
}

func (r *Repo) Signup(u *User) (*User, error) {
	_, err := r.FindByEmail(u.Email)
	if err != errors.UserNotFound {
		return nil, errors.InvalidEmail
	}

	// set verification code
	u.Verification = utils.GenerateRandomString(32)
	u.Salt = utils.GenerateRandomString(16)

	// hash the password
	u.Pass, err = hashPassword(u.Pass, u.Salt)
	if err != nil {
		return nil, err
	}

	returedUser := &User{}
	err = r.signup.Get(returedUser, u)
	if err != nil {
		return nil, err
	}

	return returedUser, nil
}

func hashPassword(password, salt string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password+salt), 14)
	return string(bytes), err
}

func checkPasswordHash(password, salt, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt))
	return err == nil
}

func (r *Repo) Authenticate(u *User) (fromDB *User, err error) {
	fromDB, err = r.FindByEmail(u.Email)
	if err != nil {
		return
	}

	if !checkPasswordHash(u.Pass, fromDB.Salt, fromDB.Pass) {
		err = errors.FailedLogin
	}

	return
}

func (r *Repo) FindByEmail(e string) (*User, error) {
	u := &User{}
	err := r.findByEmail.Get(u, e)
	if err != nil {
		return nil, errors.UserNotFound
	}

	return u, nil
}

func (r *Repo) Verify(code string) (*User, error) {
	u := &User{}
	err := r.verify.Get(u, code)
	if err != nil {
		return nil, errors.VerificationNotFound
	}

	if u.Status != StatusUnverified {
		return nil, errors.VerificationExpired
	}

	// update status
	u.Status = StatusActive
	err = r.UpdateStatus(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *Repo) UpdateStatus(u *User) (err error) {
	_, err = r.updateStatus.Exec(u)

	return
}
