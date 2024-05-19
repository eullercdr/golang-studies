package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "ha@pihego.kr", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "ha@pihego.kr", user.Email)
	assert.NotEmpty(t, user.ID)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, _ := NewUser("John Doe", "ketun@pa.mn", "123456")
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("654321"))
}
