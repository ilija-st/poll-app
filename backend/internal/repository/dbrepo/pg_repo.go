package dbrepo

import (
	"backend/ent"
	"backend/ent/user"
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// This type holds our database connections
type PostgresDBRepo struct {
	DB *ent.Client
}

// Giving users 3 seconds to interact with the database, and if it takes more then we will cancel request
const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *ent.Client {
	return m.DB
}

func (m *PostgresDBRepo) ExistsUserWithEmail(email string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	cnt, err := m.DB.User.
		Query().
		Where(user.Email(email)).
		Count(ctx)
	if err != nil {
		return false, err
	}
	if cnt > 0 {
		return true, nil
	}
	return false, nil
}

func (m *PostgresDBRepo) SeedData() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	// check if we need to seed data
	cnt, err := m.DB.User.
		Query().
		Where(user.Email("admin@example.com")).
		Count(ctx)
	if err != nil {
		return err
	}
	if cnt > 0 {
		return nil
	}

	// seed admin user
	hashPassword, err := bcrypt.GenerateFromPassword([]byte("secret"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error generating hashed password")
	}
	u, err := m.DB.User.
		Create().
		SetEmail("admin@example.com").
		SetFirstName("Admin").
		SetLastName("Example").
		SetPassword(string(hashPassword)).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating user: %w", err)
	}

	// seed poll
	poll, err := m.DB.Poll.
		Create().
		SetQuestion("What is your favorite programming language?").
		SetStatus("open").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating poll: %w", err)
	}

	u.Update().AddPolls(poll).Save(ctx)

	fmt.Println("user was updated:", u)

	fmt.Println(poll)

	// seed poll options
	po1 := m.DB.PollOption.Create().
		SetTitle("C++").
		SetPoll(poll).
		SaveX(ctx)

	po2 := m.DB.PollOption.Create().
		SetTitle("Go").
		SetPoll(poll).
		SaveX(ctx)

	m.DB.PollOption.Create().
		SetTitle("Python").
		SetPoll(poll).
		SaveX(ctx)

	m.DB.PollOption.Create().
		SetTitle("Javascript").
		SetPoll(poll).
		SaveX(ctx)

	// seed votes
	m.DB.Vote.Create().SetUser(u).SetPollOption(po1).SaveX(ctx)
	m.DB.Vote.Create().SetUser(u).SetPollOption(po2).SaveX(ctx)

	fmt.Println("Seeding data finished successfuly")

	return nil
}

func (m *PostgresDBRepo) AllPolls() ([]*ent.Poll, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	polls, err := m.DB.Poll.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	for _, p := range polls {
		poll_opts, err := p.QueryPollOptions().All(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed querying poll option %q poll: %w", p.Question, err)

		}
		p.Edges.PollOptions = append(p.Edges.PollOptions, poll_opts...)
	}

	return polls, nil
}

func (m *PostgresDBRepo) AllUsers() ([]*ent.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	users, err := m.DB.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (m *PostgresDBRepo) CreateUser(firstName string, lastName string, email string, password string) (*ent.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u, err := m.DB.User.
		Create().
		SetEmail(email).
		SetFirstName(firstName).
		SetLastName(lastName).
		SetPassword(string(hashPassword)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (m *PostgresDBRepo) GetUserByEmail(email string) (*ent.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	u, err := m.DB.User.Query().Where(user.EmailEQ(email)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (m *PostgresDBRepo) GetUserById(id int) (*ent.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	u, err := m.DB.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}
func (m *PostgresDBRepo) GetPollById(id int) (*ent.Poll, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	poll, err := m.DB.Poll.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	poll_opts, err := poll.QueryPollOptions().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying poll option %q poll: %w", poll.Question, err)

	}
	for _, po := range poll_opts {
		vs, err := po.QueryVotes().All(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed querying votes for %q poll option: %w", po.Title, err)
		}

		for _, v := range vs {
			u, err := v.QueryUser().Only(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed querying user for %q vote: %w", v, err)
			}
			v.Edges.User = u
		}

		po.Edges.Votes = append(po.Edges.Votes, vs...)
	}

	poll.Edges.PollOptions = append(poll.Edges.PollOptions, poll_opts...)

	return poll, nil
}

func (m *PostgresDBRepo) GetPollOptionById(id int) (*ent.PollOption, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	po, err := m.DB.PollOption.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// TODO: Extract into function
	vs, err := po.QueryVotes().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying votes for %q poll option: %w", po.Title, err)
	}

	for _, v := range vs {
		u, err := v.QueryUser().Only(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed querying user for %q vote: %w", v, err)
		}
		v.Edges.User = u
	}

	po.Edges.Votes = append(po.Edges.Votes, vs...)

	return po, nil
}

func (m *PostgresDBRepo) VoteOnPollOption(uid int, poid int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	err := m.DB.Vote.Create().SetPollOptionID(poid).SetUserID(uid).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
