package user

import (
	// "github.com/karlkeefer/pngr/golang/db"
	"github.com/karlkeefer/pngr/golang/errors"

	jwt "github.com/dgrijalva/jwt-go"

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
	Id           int64
	Name         string
	Email        string
	Pass         string
	Roles        int
	Status       int
	Verification string
	Created      time.Time
}

func Insert(u *User) error {
	// db.DB().Select()
	u.Verification = "randomCode"
	return nil
}

func FindByEmail(e string) (error, *User) {
	// db.DB().Select()
	return errors.UserNotFound, nil
}

func Verify(e string) (error, *Auth) {
	// TODO: check that verification code exists
	// db.DB().Select()

	// TODO: modify status and verification code for the user
	// u.Verification = ""
	// u.Status |= user.StatusVerified
	u := &User{
		Id: 40,
	}
	return buildAuth(u)
}

type Auth struct {
	JWT string
}

func Authenticate(u *User) (error, *Auth) {
	// TODO: check that email and pass match

	return buildAuth(u)
}

func buildAuth(u *User) (error, *Auth) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": jwt.MapClaims{
			"id": u.Id,
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
