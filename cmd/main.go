package main

import (
	"log"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/quippv/go-clean/configs"
	_ "github.com/quippv/go-clean/docs"
	"github.com/quippv/go-clean/internal/handler/http/middleware"
)

const (
	defaultTimeout = 30
	defaultAddress = ":9090"
)

func init() {
	configs.EnvLoad(".env")
}

func main() {
	dbConfig := configs.DefaultPostgresConfig()
	dbConn, err := configs.Open(dbConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal("Failed to ping database: ", err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal("Got error when closing the DB connection", err)
		}
	}()

	// prepare echo
	e := echo.New()
	e.Use(middleware.CORS)
	timeoutStr := configs.EnvLookup("CONTEXT_TIMEOUT")
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Println("failed to parse timeout, using default timeout")
		timeout = defaultTimeout
	}
	timeoutContext := time.Duration(timeout) * time.Second
	e.Use(middleware.SetRequestContextWithTimeout(timeoutContext))

	configs.ComposeAppSymphony(dbConn, e)

	// Start Server
	address := configs.EnvLookup("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	log.Fatal(e.Start(address))
}
