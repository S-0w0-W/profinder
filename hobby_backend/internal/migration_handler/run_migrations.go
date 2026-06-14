package db

import (
	"context"
	"fmt"

	// "log"
	"os"
	// "time"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

type DSN struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func ConnectDB(ctx context.Context) (*pgxpool.Pool, error) {
	godotenv.Load("../.env")
	dsn := DSN{
		host:     os.Getenv("DB_HOST_LOCAL"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
		dbname:   os.Getenv("POSTGRES_DB"),
	}

	dsn_string := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dsn.host, dsn.port, dsn.user, dsn.password, dsn.dbname)
	fmt.Println(dsn_string)

	cfg, err := pgxpool.ParseConfig(dsn_string)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("open pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping: %w", err)
	}

	return pool, nil
}

func RunMigrations(ctx context.Context, pool *pgxpool.Pool) error {
	// Borrow the config from the pool you already have
	db := stdlib.OpenDBFromPool(pool)
	defer db.Close()

	goose.SetDialect("postgres")
	if err := goose.UpContext(ctx, db, "db/migrations"); err != nil {
		return fmt.Errorf("goose up failed: %w", err)
	}

	return nil
}
