package utils

import (
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/karlkeefer/pngr/golang/models/user"

	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

// jwt-cookie building and parsing

var hmacSecret []byte

func init() {
	hmacSecret = []byte(os.Getenv("TOKEN_SECRET"))
	if hmacSecret == nil {
		panic("No TOKEN_SECRET environment variable was found")
	}
}

type Claims struct {
	User *user.User
	jwt.StandardClaims
}

func SetCookieForUser(w http.ResponseWriter, u *user.User) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    ToToken(u),
		HttpOnly: true,
		Secure:   true,
	})
}

func FromToken(tokenString string) (*user.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.User, nil
	}

	// anon
	return &user.User{}, nil
}

func ToToken(u *user.User) (tokenString string) {
	claims := Claims{
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
