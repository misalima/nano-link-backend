package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/misalima/nano-link-backend/src/infra/logger"
)

func ConnectDatabase(connStr string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		logger.Error("Error creating database connection: ", err)
		return nil, err
	}
	err = pool.Ping(ctx)
	if err != nil {
		logger.Error("Failed to ping to database: ", err)
		return pool, err
	}
	logger.Info("Successfully connected to database.")
	return pool, nil

}
