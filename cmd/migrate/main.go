package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var (
		migrationsPath string
		databaseURL    string
		direction      string
	)

	flag.StringVar(&migrationsPath, "path", "./migrations", "path to migrations folder")
	flag.StringVar(&databaseURL, "db", "", "database connection string (e.g., postgres://user:pass@host:port/db?sslmode=disable)")
	flag.StringVar(&direction, "direction", "up", "migration direction: up or down")
	flag.Parse()

	if databaseURL == "" {
		databaseURL = os.Getenv("DATABASE_URL")
		if databaseURL == "" {
			log.Fatal("database URL is required: use -db flag or DATABASE_URL env var")
		}
	}

	m, err := migrate.New(
		fmt.Sprintf("file://%s", migrationsPath),
		databaseURL,
	)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}
	defer func() {
		srcErr, dbErr := m.Close()
		if srcErr != nil {
			log.Printf("source close error: %v", srcErr)
		}
		if dbErr != nil {
			log.Printf("database close error: %v", dbErr)
		}
	}()

	switch direction {
	case "up":
		if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("failed to run migrations up: %v", err)
		}
		log.Println("migrations applied successfully")
	case "down":
		if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatalf("failed to run migrations down: %v", err)
		}
		log.Println("migrations rolled back successfully")
	default:
		log.Fatalf("unknown direction: %s (use 'up' or 'down')", direction)
	}
}
