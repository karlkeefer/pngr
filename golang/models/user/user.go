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
	Pass         string
	Status       int
	Verification string
	Created      time.Time
}

// MarshalJSON here protects "private" fields from ever being sent *out*
// it also makes Name return "" instead of null
func (u User) MarshalJSON() ([]byte, error) {
	var tmp struct {
		ID      int64      `json:"id"`
		Name    string     `json:"name,omitempty"`
		Email   string     `json:"email,omitempty"`
		Status  int        `json:"status,omitempty"`
		Created *time.Time `json:"created,omitempty"`
	}

	tmp.ID = u.ID

	// pick a name
	if u.Name != nil {
		tmp.Name = *u.Name
	}

	tmp.Email = u.Email
	tmp.Status = u.Status
	if !u.Created.IsZero() {
		tmp.Created = &u.Created
	}
	return json.Marshal(&tmp)
}

const (
	UserStatusUnverified = 0
	UserStatusActive     = 1
	UserStatusDisabled   = 2
)

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
	signup, err := db.PrepareNamed(`INSERT INTO users (email, pass, status, verification) VALUES (:email, :pass, :status, :verification) RETURNING *`)
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

	// hash the password
	u.Pass, err = hashPassword(u.Pass)
	if err != nil {
		return nil, err
	}

	fromDB := &User{}
	err = r.signup.Get(fromDB, u)
	if err != nil {
		return nil, err
	}

	return fromDB, nil
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

	if u.Status != UserStatusUnverified {
		return nil, errors.VerificationExpired
	}

	// update status
	u.Status = UserStatusActive
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
