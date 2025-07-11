// in cmd/main.go
package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/pas-platform/identity-service/internal/api"
	"github.com/pas-platform/identity-service/internal/storage"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// Initialize Database Connection
	db := connectToDB()
	log.Println("Database connection established")
	defer db.Close()

	// Run Database Migrations
	runMigrations(os.Getenv("DATABASE_URL"))

	// Inject Dependencies
	store := storage.NewPostgresStore(db)
	handler := api.NewAPIHandler(store)
	router := api.SetupRouter(handler)

	// Start Server
	log.Println("Identity & Access Service starting on port 8081...")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// runMigrations applies database migrations.
func runMigrations(dbURL string) {
	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Database migrations applied successfully")
}



func connectToDB() *sql.DB {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	var db *sql.DB
	var err error
	maxRetries := 5
	retryDelay := 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("pgx", dbURL)
		if err != nil {
			log.Fatalf("Invalid database URL: %v", err)
		}

		err = db.Ping()
		if err == nil {
			return db
		}

		log.Printf("Could not connect to database (attempt %d/%d): %v. Retrying in %v...", i+1, maxRetries, err, retryDelay)
		time.Sleep(retryDelay)
		retryDelay *= 2
	}

	log.Fatalf("Failed to connect to database after %d attempts", maxRetries)
	return nil
}