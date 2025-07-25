package db

import (
	"fmt"
	"log"
	"time"

	"github.com/sebenitezg/hotel-service/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewConnection(dbConfig config.DatabaseConfigurations) *bun.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName)

	log.Printf("Connecting to database: %s, pool-size: %d", dsn, dbConfig.PoolSize)

	// Parse config from string
	parsedCfg, err := pgx.ParseConfig(dsn)

	if err != nil {
		log.Fatalf("Error parsing database config: %v", err)
		return nil
	}

	// Init a connection compatible with standard library
	conn := stdlib.OpenDB(*parsedCfg)

	// Connection pool settings
	conn.SetMaxOpenConns(dbConfig.PoolSize)
	conn.SetMaxIdleConns(dbConfig.PoolSize)
	conn.SetConnMaxLifetime(3 * time.Minute)

	// Test connection
	err = conn.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil
	}

	log.Printf("Successfully connected to database")

	db := bun.NewDB(conn, pgdialect.New(), bun.WithDiscardUnknownColumns())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(dbConfig.LogQueries)))

	return db
}
