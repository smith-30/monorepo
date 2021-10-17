package model

type User struct {
	ID uint
}

func NewUser() (*User, error) {
	return &User{
		ID: 1,
	}, nil
}
