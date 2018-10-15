package jwt

import (
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models/user"
)

// jwt-cookie building and parsing
const cookieName = "pngr-jwt"

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

// RequireAuth middleware makes sure the user exists based on their JWT
func RequireAuth(r *http.Request, fn func(*user.User) http.HandlerFunc) http.HandlerFunc {
	u, err := ParseUserCookie(r)
	if u.ID == 0 || err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			errors.Write(w, errors.RouteUnauthorized)
		}
	}
	return fn(u)
}

// WriteUserCookie encodes a user's JWT and sets it as an httpOnly & Secure cookie
func WriteUserCookie(w http.ResponseWriter, u *user.User) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    encodeUser(u),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})
}

// ParseUserCookie builds a user object from a JWT, if it's valid
func ParseUserCookie(r *http.Request) (*user.User, error) {
	cookie, _ := r.Cookie(cookieName)
	var tokenString string
	if cookie != nil {
		tokenString = cookie.Value
	}

	if tokenString == "" {
		return &user.User{}, nil
	}

	return decodeUser(tokenString)
}

// encodeUser convert a user struct into a jwt
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

// decodeUser converts a jwt into a user struct (or returns a zero-value user)
func decodeUser(tokenString string) (*user.User, error) {
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
