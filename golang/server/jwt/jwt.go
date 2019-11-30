package jwt

import (
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models/user"
	"github.com/karlkeefer/pngr/golang/server/write"
)

// jwt-cookie building and parsing
const cookieName = "pngr-jwt"

// tokens auto-refresh at the end of their lifetime,
// so long as the user hasn't been disabled in the interim
const tokenLifetime = time.Hour * 6

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
func RequireAuth(minStatus user.Status, env *env.Env, w http.ResponseWriter, r *http.Request, fn func(*user.User) http.HandlerFunc) http.HandlerFunc {
	u, err := HandleUserCookie(env, w, r)
	if err != nil {
		return write.Error(err)
	}

	if u.Status < minStatus {
		return write.Error(errors.RouteUnauthorized)
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

// HandleUserCookie builds a user object from a JWT, if it's valid
// Also attempts to refresh and reset the cookie if the user submits an expired token
func HandleUserCookie(env *env.Env, w http.ResponseWriter, r *http.Request) (*user.User, error) {
	cookie, _ := r.Cookie(cookieName)
	var tokenString string
	if cookie != nil {
		tokenString = cookie.Value
	}

	if tokenString == "" {
		return &user.User{}, nil
	}

	u, err := decodeUser(tokenString)

	// attempt refresh of expired token:
	if err == errors.ExpiredToken && u.Status > 0 {
		user, fetchError := env.UserRepo().FindByEmail(u.Email)
		if fetchError != nil {
			return nil, err
		}
		if user.Status > 0 {
			WriteUserCookie(w, user)
			return user, nil
		}
	}

	return u, err
}

// encodeUser convert a user struct into a jwt
func encodeUser(u *user.User) (tokenString string) {
	claims := claims{
		u,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Add(-time.Second).Unix(),
			ExpiresAt: time.Now().Add(tokenLifetime).Unix(),
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
		// check for expired token
		if verr, ok := err.(*jwt.ValidationError); ok {
			if verr.Errors&jwt.ValidationErrorExpired != 0 {
				return getUserFromToken(token), errors.ExpiredToken
			}
		}
	}

	if !token.Valid || err != nil {
		return nil, errors.InvalidToken
	}

	return getUserFromToken(token), nil
}

func getUserFromToken(token *jwt.Token) *user.User {
	if claims, ok := token.Claims.(*claims); ok {
		return claims.User
	}

	return nil
}
