package jwt

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models"
	"github.com/karlkeefer/pngr/golang/repos/user"
	"github.com/stretchr/testify/assert"
)

// expiredToken has uid of 30
const expiredToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjozMCwibmFtZSI6IiIsInN0YXR1cyI6MX0sImV4cCI6MTU3NTA4Njg5MCwiaWF0IjoxNTc1MDg2ODgzfQ.8HA0d-yOQNl6JIuQUfnuZFVxOgb6799KgmEGcugqFy0"

func TestHandleUserCookie(t *testing.T) {
	tests := []struct {
		// setup
		name     string
		mockUser *models.User // what will FindByEmail return?
		token    string
		// expectations
		shouldSetCookie bool
		expectedError   error
		expectedUid     int64
	}{
		{
			"user marked disabled since token expired (should log them out)",
			&models.User{
				ID:     30,
				Status: models.UserStatusDisabled,
			},
			expiredToken,
			true,
			nil,
			0,
		},
		{
			"token is not expired, no action",
			nil,
			encodeUser(&models.User{
				ID:     30,
				Status: models.UserStatusActive,
			}),
			false,
			nil,
			30,
		},
		{
			"token refresh when token is expired and user is still valid",
			&models.User{
				ID:     30,
				Status: models.UserStatusActive,
			},
			expiredToken,
			true,
			nil,
			30,
		},
	}

	for _, test := range tests {
		test := test // capture range variable
		t.Run(test.name, func(t *testing.T) {

			// setup mock
			ur := user.Mock(test.mockUser, nil)
			env := env.Mock(nil, ur, nil)

			// build a fake request with an expired token
			r := httptest.NewRequest(http.MethodGet, "/path", nil)
			r.AddCookie(&http.Cookie{
				Name:  cookieName,
				Value: test.token,
			})
			w := httptest.NewRecorder()

			u, err := HandleUserCookie(env, w, r)

			// make sure it returns the expected user and error
			assert.Equal(t, test.expectedError, err)
			assert.Equal(t, test.expectedUid, u.ID)

			// make sure it sets a cookie if the cookie is expired and the user is still valid
			resp := w.Result()
			if test.shouldSetCookie {
				assert.Len(t, resp.Cookies(), 1)
			} else {
				assert.Len(t, resp.Cookies(), 0)
			}
		})
	}
}

func TestUserFromCookie(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/path", nil)

	// make sure empty cookie returns "anon user"
	u, err := userFromCookie(r)
	assert.Equal(t, int64(0), u.ID)
	assert.NoError(t, err)

	// now give the request a valid cookie
	r.AddCookie(&http.Cookie{
		Name: cookieName,
		Value: encodeUser(&models.User{
			ID:     30,
			Status: models.UserStatusActive,
		}),
	})

	u, err = userFromCookie(r)
	assert.Equal(t, int64(30), u.ID)
	assert.NoError(t, err)
}

// Note: more thorough testing of encode/decode stuff requires us to mock out the
// time package, which would be a big PITA
func TestDecodeUser(t *testing.T) {
	token := encodeUser(&models.User{
		ID:     30,
		Status: models.UserStatusActive,
	})
	assert.NotEmpty(t, token)

	du, err := decodeUser(token)
	assert.NoError(t, err)
	assert.Equal(t, int64(30), du.ID)
	assert.Equal(t, models.UserStatusActive, du.Status)

	// test old token
	_, err = decodeUser(expiredToken)
	assert.Equal(t, errors.ExpiredToken, err)

	// test garbage token
	_, err = decodeUser("garbage_token")
	assert.Equal(t, errors.InvalidToken, err)
}
