package database

import (
	"fmt"
)

var users = []User{
	{
		ID:       "1",
		Name:     "Michael",
		Username: "MD",
		Password: "123456",
		Type:     "ADMIN",
	},
	{
		ID:       "2",
		Name:     "Isabela",
		Username: "Isa",
		Password: "123123",
		Type:     "BASIC",
	},
}

type User struct {
	ID       string
	Name     string
	Username string
	Password string
	Type     string
}

func NewUser() *User {
	return &User{}
}
func (u *User) ListAll() ([]User, error) {
	return users, nil
}

func (u *User) FindById(id string) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}

	return User{}, fmt.Errorf("User not found with ID %s", id)
}
