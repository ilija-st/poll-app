package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"flag"

	_ "github.com/lib/pq"

	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN    string
	Domain string
	DB     repository.DatabaseRepo
}

func main() {
	var app application

	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=polls sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")

	app.Domain = "poll.com"
	log.Println("Starting application on port", port)

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	err = app.DB.SeedData()
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
