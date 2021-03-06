package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// make sure we don't marshal the salt/pass/verification fields
func TestMarshalJSON(t *testing.T) {
	u := &User{
		ID:           2,
		Name:         nil,
		Email:        "test@test.com",
		Salt:         "saltyboi",
		Pass:         "myHashGoesHere",
		Status:       UserStatusActive,
		Verification: "VerificationCode",
		CreatedAt:    time.Unix(0, 0),
	}

	out, err := u.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, `{"id":2,"name":"","email":"test@test.com","status":1,"created_at":"1970-01-01T00:00:00Z"}`, string(out))
}
