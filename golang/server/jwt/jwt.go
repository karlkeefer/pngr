package jwt

import (
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/karlkeefer/pngr/golang/models/user"
)

// jwt-cookie building and parsing

var hmacSecret []byte

func init() {
	hmacSecret = []byte(os.Getenv("TOKEN_SECRET"))
	if hmacSecret == nil {
		panic("No TOKEN_SECRET environment variable was found")
	}
}

type claims struct {
	User *user.User
	jwt.StandardClaims
}

func WriteUserCookie(w http.ResponseWriter, u *user.User) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    encodeUser(u),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})
}

func ParseUser(tokenString string) (*user.User, error) {
	if tokenString == "" {
		return &user.User{}, nil
	}
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*claims); ok && token.Valid {
		return claims.User, nil
	}

	// anon
	return &user.User{}, nil
}

func encodeUser(u *user.User) (tokenString string) {
	claims := claims{
		u,
		jwt.StandardClaims{
			// TODO: setup appropriate JWT values time-related claims
			NotBefore: time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// unhandled err here
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		log.Println("Error signing token", err)
	}
	return
}
