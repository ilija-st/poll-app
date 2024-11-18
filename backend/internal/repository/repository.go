package repository

import (
	"backend/ent"
)

type DatabaseRepo interface {
	Connection() *ent.Client
	AllPolls() ([]*ent.Poll, error)
	AllUsers() ([]*ent.User, error)
	SeedData() error
}
