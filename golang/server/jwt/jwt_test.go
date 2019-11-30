package jwt

import (
	"testing"

	"github.com/karlkeefer/pngr/golang/errors"
	"github.com/karlkeefer/pngr/golang/models/user"
	"github.com/stretchr/testify/assert"
)

// Note: more thorough testing of encode/decode stuff requires us to mock out the
// time package, which would be a big PITA
func TestDecodeUser(t *testing.T) {
	u := &user.User{
		ID:     30,
		Status: user.StatusActive,
	}

	token := encodeUser(u)
	assert.NotEmpty(t, token)

	du, err := decodeUser(token)
	assert.NoError(t, err)
	assert.Equal(t, int64(30), du.ID)
	assert.Equal(t, user.Status(1), du.Status)

	// test old token
	oldToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjozMCwibmFtZSI6IiIsInN0YXR1cyI6MX0sImV4cCI6MTU3NTA4Njg5MCwiaWF0IjoxNTc1MDg2ODgzfQ.8HA0d-yOQNl6JIuQUfnuZFVxOgb6799KgmEGcugqFy0"

	_, err = decodeUser(oldToken)
	assert.Equal(t, errors.ExpiredToken, err)

	// test garbage token
	_, err = decodeUser("not_a_real_token")
	assert.Equal(t, "token contains an invalid number of segments", err.Error())
}
