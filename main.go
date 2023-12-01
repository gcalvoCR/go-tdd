package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gcalvocr/go-tdd/routes"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/gorilla/mux"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func main() {
	// Connection to DB
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error making DB connected: %s", err.Error())
	}

	// We run the migrations if needed
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error making DB driver: %s", err.Error())
	}

	migrator, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("Error making migration engine: %s", err.Error())
	}
	migrator.Steps(2)

	r := mux.NewRouter()

	r.HandleFunc("/healthcheck", routes.HandleHealthCheck)

	s := http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	s.ListenAndServe()
}
