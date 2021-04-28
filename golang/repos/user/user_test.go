package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	pass := "MyPassword"
	salt := "MySalt"

	hash, err := hashPassword(pass, salt)
	assert.NoError(t, err)

	result := checkPasswordHash("badPass", salt, hash)
	assert.False(t, result)

	result = checkPasswordHash(pass, salt, hash)
	assert.True(t, result)
}
