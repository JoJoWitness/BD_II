package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return fmt.Errorf("DATABASE_URL no está definida")
	}
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("db.Ping: %w", err)
	}
	DB = db
	return nil
}

func Close() error {
	if DB == nil {
		return nil
	}
	return DB.Close()
}
