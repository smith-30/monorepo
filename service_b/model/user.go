package model

type User struct {
	ID string
}

func NewUser() (*User, error) {
	return &User{
		ID: "",
	}, nil
}
