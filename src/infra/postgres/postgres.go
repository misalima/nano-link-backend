package postgres

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDatabase(connStr string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatal("Error creating database connection: ", err)
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		log.Fatal("Failed to ping to database: ", err)
		return pool, err
	}
	log.Println("Successfully connected to database.")
	return pool, nil

}
