package jwt

import (
	"log"
	"net/http"
	"os"
	"time"

	mdb "github.com/QuinnMain/infograph/golang/db"

	"github.com/QuinnMain/infograph/golang/env"
	"github.com/QuinnMain/infograph/golang/errors"
	jwt "github.com/dgrijalva/jwt-go" //  JSON Web Tokens
)

// jwt-cookie building and parsing
const cookieName = "pngr-jwt"
const insecureSecret = "asd973hkalkjhx97asdh"

// tokens auto-refresh at the end of their lifetime,
// so long as the user hasn't been disabled in the interim
const tokenLifetime = time.Hour * 6

var hmacSecret []byte

func init() {
	hmacSecret = []byte(os.Getenv("TOKEN_SECRET"))
	if hmacSecret == nil {
		panic("No TOKEN_SECRET environment variable was found")
	}
	if string(hmacSecret) == insecureSecret {
		log.Print("\n\n*** WARNING ***\nYour JWT isn't secure!\nYou need to change your TOKEN_SECRET variable in .env (and restart your containers).\n\n")
	}
}

type claims struct {
	User *mdb.MUser
	jwt.StandardClaims
}

// WriteUserCookie encodes a user's JWT and sets it as an httpOnly & Secure cookie
func WriteUserCookie(w http.ResponseWriter, u *mdb.MUser) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    encodeUser(u, time.Now()),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   60 * 60 * 24 * 7, // one week
	})
}

// HandleUserCookie attempts to refresh an expired token if the user is still valid
func HandleUserCookie(e env.Env, w http.ResponseWriter, r *http.Request) (*mdb.MUser, error) {
	u, err := userFromCookie(r)

	// attempt refresh of expired token:
	if err == errors.ExpiredToken && *u.Status == mdb.UserStatusActive {
		user, fetchError := e.DB().FindUserByEmail(r.Context(), u.Email)
		if fetchError != nil {
			return wipeCookie(e, w)
		}
		if *user.Status == mdb.UserStatusActive {
			WriteUserCookie(w, &user)
			return &user, nil
		} else {
			// their account isn't verified, log them out
			return wipeCookie(e, w)
		}
	}

	if err != nil {
		return nil, err
	}

	return u, err
}

func wipeCookie(e env.Env, w http.ResponseWriter) (*mdb.MUser, error) {
	u := &mdb.MUser{}
	WriteUserCookie(w, u)
	return u, nil
}

// userFromCookie builds a user object from a JWT, if it's valid
func userFromCookie(r *http.Request) (*mdb.MUser, error) {
	cookie, _ := r.Cookie(cookieName)
	var tokenString string
	if cookie != nil {
		tokenString = cookie.Value
	}

	if tokenString == "" {
		return &mdb.MUser{}, nil
	}

	return decodeUser(tokenString)
}

// encodeUser convert a user struct into a jwt
func encodeUser(u *mdb.MUser, t time.Time) (tokenString string) {
	claims := claims{
		u,
		jwt.StandardClaims{
			IssuedAt:  t.Add(-time.Second).Unix(),
			ExpiresAt: t.Add(tokenLifetime).Unix(),
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
func decodeUser(tokenString string) (*mdb.MUser, error) {
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

	if err != nil || !token.Valid {
		return nil, errors.InvalidToken
	}

	return getUserFromToken(token), nil
}

func getUserFromToken(token *jwt.Token) *mdb.MUser {
	if claims, ok := token.Claims.(*claims); ok {
		return claims.User
	}

	return &mdb.MUser{}
}
