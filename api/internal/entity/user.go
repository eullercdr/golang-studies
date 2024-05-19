package entity

import (
	"github.com/eullercdr/go/api/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, error := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if error != nil {
		return nil, error
	}
	return &User{
		ID:       entity.NewID().String(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
