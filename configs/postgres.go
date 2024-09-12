package configs

import (
	"database/sql"
	"fmt"
	"time"
)

func NewPostgresDB() *sql.DB {
	postgresDbName := GetEnv("POSTGRES_DB", "lifefolio_db")
	postgresUser := GetEnv("POSTGRES_USER", "postgres")
	postgresPassword := GetEnv("POSTGRES_PASSWORD", "potsgres")

	postgresURL := fmt.Sprintf(
		"user=%v password=%v dbname=%v sslmode=disable",
		postgresUser, postgresPassword, postgresDbName,
	)

	driverName := "postgres"
	db, err := sql.Open(driverName, postgresURL)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
