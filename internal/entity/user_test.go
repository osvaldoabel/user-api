package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("user 01", "email@example.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "user 01", user.Name)
	assert.Equal(t, "email@example.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("user 01", "email@example.com", "123456")
	assert.Nil(t, err)
	assert.NotEmpty(t, user.Password)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("123456789"))
	assert.NotEqual(t, "123456789", user.Password)
}

func TestUser_ValidatePasswordWithUnauthorizedPassword(t *testing.T) {
	user, err := NewUser("user 01", "email@example.com", "123")
	assert.Nil(t, err)

	assert.False(t, user.ValidatePassword("123"))
}
