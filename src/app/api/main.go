package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/misalima/nano-link-backend/src/app/api/config"
	"github.com/misalima/nano-link-backend/src/app/api/container"
	"github.com/misalima/nano-link-backend/src/app/api/router"
	"github.com/misalima/nano-link-backend/src/infra/postgres"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	cfg := config.LoadConfig()
	connStr := cfg.GetConnString()

	db, err := postgres.ConnectDatabase(connStr)
	if err != nil {
		log.Fatal("Couldn't connect to database")
	}

	c := container.New(db)

	startServer(c)
}

func startServer(c *container.Container) {
	e := router.NewRouter(c)
	err := e.Start(":8080")
	if err != nil {
		return
	}
	fmt.Println("Server started")
}
