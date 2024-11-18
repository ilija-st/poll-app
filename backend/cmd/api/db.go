package main

import (
	"backend/ent"
	"context"
)

func (app *application) connectToDB() (*ent.Client, error) {
	client, err := ent.Open("postgres", app.DSN)
	if err != nil {
		return nil, err
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}
