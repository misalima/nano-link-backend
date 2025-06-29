package main

import (
	"github.com/joho/godotenv"
	"github.com/misalima/nano-link-backend/src/app/api/config"
	"github.com/misalima/nano-link-backend/src/app/api/container"
	"github.com/misalima/nano-link-backend/src/app/api/router"
	"github.com/misalima/nano-link-backend/src/infra/logger"
	"github.com/misalima/nano-link-backend/src/infra/postgres"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Fatal("Error loading .env file: ", err)
	}

	cfg := config.LoadConfig()

	logger.Init(cfg.LoggerConfig.Environment)
	logger.Info("Starting application...")

	connStr := cfg.GetConnString()

	db, err := postgres.ConnectDatabase(connStr)
	if err != nil {
		logger.Fatal("Couldn't connect to database: ", err)
	}

	c := container.New(db)

	startServer(c, cfg.ServerConfig)
}

func startServer(c *container.Container, config config.ServerConfig) {
	e := router.NewRouter(c)

	serverAddr := config.Host + ":" + config.Port
	logger.Infof("Starting server on %s", serverAddr)

	err := e.Start(serverAddr)
	if err != nil {
		logger.Errorf("Server error: %v", err)
		return
	}
}
