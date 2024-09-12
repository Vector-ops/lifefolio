package database

import (
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/vector-ops/lifefolio/configs"
	"github.com/vector-ops/lifefolio/ent"
)

type Database struct {
	EntClient *ent.Client
	RDB       *redis.Client
}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) SetupEntClient() {
	pgdb := configs.NewPostgresDB()
	drv := entsql.OpenDB("pgx", pgdb)

	db.EntClient = ent.NewClient(ent.Driver(drv))
}

func (db *Database) SetupRedis() {
	db.RDB = configs.NewRedisClient()
}
