package user

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models"
	"github.com/karlkeefer/pngr/golang/utils"
	"golang.org/x/crypto/bcrypt"
)

type Repo interface {
	Signup(u *models.User) (*models.User, error)
	Authenticate(u *models.User) (*models.User, error)
	FindByEmail(e string) (*models.User, error)
	Verify(code string) (*models.User, error)
	UpdateStatus(u *models.User) error
}

type repo struct {
	signup       *sqlx.NamedStmt
	updateStatus *sqlx.NamedStmt
	findByEmail  *sqlx.Stmt
	verify       *sqlx.Stmt
}

func NewRepo(db *sqlx.DB) Repo {
	// TODO: create a helper to prepare named and regular statements
	signup, err := db.PrepareNamed(`INSERT INTO users (email, salt, pass, status, verification) VALUES (LOWER(:email), :salt, :pass, :status, :verification) RETURNING *`)
	if err != nil {
		panic(err)
	}
	updateStatus, err := db.PrepareNamed(`UPDATE users SET status = :status, updated_at = :updated_at WHERE id = :id`)
	if err != nil {
		panic(err)
	}
	findByEmail, err := db.Preparex(`SELECT * FROM users WHERE email = LOWER($1) LIMIT 1`)
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

func (r *repo) Signup(u *models.User) (*models.User, error) {
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

	returedUser := &models.User{}
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

func (r *repo) Authenticate(u *models.User) (dbUser *models.User, err error) {
	dbUser, err = r.FindByEmail(u.Email)
	if err != nil {
		return
	}

	if !checkPasswordHash(u.Pass, dbUser.Salt, dbUser.Pass) {
		err = errors.FailedLogin
	}

	return
}

func (r *repo) FindByEmail(e string) (*models.User, error) {
	u := &models.User{}
	err := r.findByEmail.Get(u, e)
	if err != nil {
		return nil, errors.UserNotFound
	}

	return u, nil
}

func (r *repo) Verify(code string) (*models.User, error) {
	u := &models.User{}
	err := r.verify.Get(u, code)
	if err != nil {
		return nil, errors.VerificationNotFound
	}

	if u.Status != models.UserStatusUnverified {
		return nil, errors.VerificationExpired
	}

	// update status
	u.Status = models.UserStatusActive
	err = r.UpdateStatus(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *repo) UpdateStatus(u *models.User) (err error) {
	u.UpdatedAt = time.Now()
	_, err = r.updateStatus.Exec(u)

	return
}

// mock this repository!
type mock struct {
	user *models.User
	err  error
}

func Mock(u *models.User, e error) Repo {
	return &mock{
		u,
		e,
	}
}

func (m *mock) Signup(u *models.User) (*models.User, error) {
	return m.user, m.err
}
func (m *mock) Authenticate(u *models.User) (*models.User, error) {
	return m.user, m.err
}
func (m *mock) FindByEmail(e string) (*models.User, error) {
	return m.user, m.err
}
func (m *mock) Verify(code string) (*models.User, error) {
	return m.user, m.err
}
func (m *mock) UpdateStatus(u *models.User) error {
	return m.err
}
