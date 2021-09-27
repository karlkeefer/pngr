package jwt

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/karlkeefer/pngr/golang/db"
	"github.com/karlkeefer/pngr/golang/db/wrapper"
	"github.com/karlkeefer/pngr/golang/env"
	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/stretchr/testify/assert"
)

func TestHandleUserCookie(t *testing.T) {

	expired := encodeUser(&db.User{
		ID:     30,
		Email:  "j@j.com",
		Status: db.UserStatusActive,
	}, time.Now().Add(-2*tokenLifetime))

	tests := []struct {
		// setup
		name     string
		mockUser db.User // what will FindByEmail return?
		mockErr  error
		token    string
		// expectations
		shouldSetCookie bool
		expectedError   error
		expectedUser    *db.User
	}{
		{
			"user marked disabled since token expired (should log them out)",
			db.User{
				ID:     30,
				Status: db.UserStatusDisabled,
			},
			nil,
			expired,
			//result
			true,
			nil,
			&db.User{},
		},
		{
			"token is not expired, no action",
			db.User{},
			nil,
			encodeUser(&db.User{
				ID:     30,
				Status: db.UserStatusActive,
			}, time.Now()),
			false,
			nil,
			&db.User{
				ID:     30,
				Status: db.UserStatusActive,
			},
		},
		{
			"token refresh when token is expired and user is still valid",
			db.User{
				ID:     30,
				Status: db.UserStatusActive,
			},
			nil,
			expired,
			true,
			nil,
			&db.User{
				ID:     30,
				Status: db.UserStatusActive,
			},
		},
	}

	for _, test := range tests {
		test := test // capture range variable
		t.Run(test.name, func(t *testing.T) {

			// setup mock
			ctrl := gomock.NewController(t)
			mockDB := wrapper.NewMockQuerier(ctrl)
			env := env.Mock(mockDB)

			mockDB.EXPECT().FindUserByEmail(gomock.Any(), gomock.Any()).Return(test.mockUser, test.mockErr).AnyTimes()

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
			assert.Equal(t, test.expectedUser, u)

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
		Value: encodeUser(&db.User{
			ID:     30,
			Status: db.UserStatusActive,
		}, time.Now()),
	})

	u, err = userFromCookie(r)
	assert.Equal(t, int64(30), u.ID)
	assert.NoError(t, err)
}

func TestDecodeUser(t *testing.T) {
	token := encodeUser(&db.User{
		ID:     30,
		Email:  "j@j.com",
		Status: db.UserStatusActive,
	}, time.Now())
	assert.NotEmpty(t, token)

	du, err := decodeUser(token)
	assert.NoError(t, err)
	assert.Equal(t, int64(30), du.ID)
	assert.Equal(t, "j@j.com", du.Email)
	assert.Equal(t, db.UserStatusActive, du.Status)
}

func TestDecodeUser_OldToken(t *testing.T) {
	token := encodeUser(&db.User{
		ID:     30,
		Email:  "j@j.com",
		Status: db.UserStatusActive,
	}, time.Now().Add(-2*tokenLifetime))

	assert.NotEmpty(t, token)

	_, err := decodeUser(token)
	assert.Equal(t, errors.ExpiredToken, err)
}

func TestDecodeUser_Invalid(t *testing.T) {
	_, err := decodeUser("garbage_token")
	assert.Equal(t, errors.InvalidToken, err)
}
