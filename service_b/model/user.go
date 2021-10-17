package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID string
}

func NewUser() (*User, error) {
	return &User{
		ID: uuid.New().String(),
	}, nil
}
