package configs

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgresDB() *sql.DB {
	postgresHost := GetEnv("POSTGRES_HOST", "localhost")
	postgresDbName := GetEnv("POSTGRES_DB", "lifefolio_db")
	postgresUser := GetEnv("POSTGRES_USER", "lifefolio")
	postgresPassword := GetEnv("POSTGRES_PASSWORD", "postgres")

	postgresURL := fmt.Sprintf(
		"host=%s user=%v password=%v dbname=%v sslmode=disable",
		postgresHost, postgresUser, postgresPassword, postgresDbName,
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
