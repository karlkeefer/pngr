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
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
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
		ID        int64      `json:"id"`
		Name      string     `json:"name"`
		Email     string     `json:"email,omitempty"`
		Status    Status     `json:"status"`
		CreatedAt *time.Time `json:"created_at,omitempty"`
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
	if !u.CreatedAt.IsZero() {
		tmp.CreatedAt = &u.CreatedAt
	}
	return json.Marshal(&tmp)
}

// REPO stuff
// TODO: consider moving repo to separate package
type Repo interface {
	Signup(u *User) (*User, error)
	Authenticate(u *User) (*User, error)
	FindByEmail(e string) (*User, error)
	Verify(code string) (*User, error)
	UpdateStatus(u *User) error
}

type repo struct {
	signup       *sqlx.NamedStmt
	updateStatus *sqlx.NamedStmt
	findByEmail  *sqlx.Stmt
	verify       *sqlx.Stmt
}

func NewRepo(db *sqlx.DB) Repo {
	// TODO: create a helper to prepare named and regular statements
	signup, err := db.PrepareNamed(`INSERT INTO users (email, salt, pass, status, verification) VALUES (:email, :salt, :pass, :status, :verification) RETURNING *`)
	if err != nil {
		panic(err)
	}
	updateStatus, err := db.PrepareNamed(`UPDATE users SET status = :status, updated_at = :updated_at WHERE id = :id`)
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
	return &repo{
		signup,
		updateStatus,
		findByEmail,
		verify,
	}
}

func (r *repo) Signup(u *User) (*User, error) {
	_, err := r.FindByEmail(u.Email)
	if err != errors.UserNotFound {
		return nil, errors.InvalidEmail
	}

	// set verification code
	u.Verification = utils.GenerateRandomString(32)
	u.Salt = utils.GenerateRandomString(32)

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

func (r *repo) Authenticate(u *User) (dbUser *User, err error) {
	dbUser, err = r.FindByEmail(u.Email)
	if err != nil {
		return
	}

	if !checkPasswordHash(u.Pass, dbUser.Salt, dbUser.Pass) {
		err = errors.FailedLogin
	}

	return
}

func (r *repo) FindByEmail(e string) (*User, error) {
	u := &User{}
	err := r.findByEmail.Get(u, e)
	if err != nil {
		return nil, errors.UserNotFound
	}

	return u, nil
}

func (r *repo) Verify(code string) (*User, error) {
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

func (r *repo) UpdateStatus(u *User) (err error) {
	u.UpdatedAt = time.Now()
	_, err = r.updateStatus.Exec(u)

	return
}

// mock this repository!
type mock struct {
	user *User
	err  error
}

func Mock(u *User, e error) Repo {
	return &mock{
		u,
		e,
	}
}

func (m *mock) Signup(u *User) (*User, error) {
	return m.user, m.err
}
func (m *mock) Authenticate(u *User) (*User, error) {
	return m.user, m.err
}
func (m *mock) FindByEmail(e string) (*User, error) {
	return m.user, m.err
}
func (m *mock) Verify(code string) (*User, error) {
	return m.user, m.err
}
func (m *mock) UpdateStatus(u *User) error {
	return m.err
}
