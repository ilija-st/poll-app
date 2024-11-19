package repository

import (
	"backend/ent"
)

type DatabaseRepo interface {
	Connection() *ent.Client
	AllPolls() ([]*ent.Poll, error)
	AllUsers() ([]*ent.User, error)
	CreateUser(firstName string, lastName string, email string, password string) (*ent.User, error)
	GetUserByEmail(email string) (*ent.User, error)
	GetUserById(id int) (*ent.User, error)
	ExistsUserWithEmail(email string) (bool, error)
	GetPollById(id int) (*ent.Poll, error)
	GetPollOptionById(id int) (*ent.PollOption, error)
	VoteOnPollOption(uid int, poid int) error
	SeedData() error
}
