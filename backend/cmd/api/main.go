package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"flag"
	"time"

	_ "github.com/lib/pq"

	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN            string
	Domain         string
	DB             repository.DatabaseRepo
	auth           Auth
	JWTSecret      string
	JWTIssuer      string
	JWTAudience    string
	CookieDomain   string
	FrontendDomain string
}

func main() {
	// set application config - type that stores info that app needs (ex. for database, or jwt secret, etc.)
	var app application

	// read from command line - we might pass flag that says this api is only for domain example.com
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=polls sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "verysecret", "signing secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "poll.com", "signing issuer")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "poll.com", "signing audience")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	flag.StringVar(&app.Domain, "domain", "poll.com", "domain")
	flag.StringVar(&app.FrontendDomain, "frontend-domain", "http://localhost:3000", "domain")
	flag.Parse()

	log.Println("Starting application on port", port)

	// connect to the database
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

	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookiePath:    "/",
		CookieName:    "refresh_token",
		CookieDomain:  app.CookieDomain,
	}

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
