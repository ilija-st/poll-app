package repository

import (
	"backend/ent"
)

type DatabaseRepo interface {
	Connection() *ent.Client
	AllPolls() ([]*ent.Poll, error)
	AllUsers() ([]*ent.User, error)
	GetUserByEmail(email string) (*ent.User, error)
	GetUserById(id int) (*ent.User, error)
	SeedData() error
}
