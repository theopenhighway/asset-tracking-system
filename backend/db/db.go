package database

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB() (*sql.DB, error) {
	// "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
