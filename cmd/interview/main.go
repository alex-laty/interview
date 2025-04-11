package main

import (
	"fmt"
	"interview/internal/config"
	"interview/internal/logger"
	"interview/internal/storage/pgsql"
)

func main() {
	cfg := config.LoadConfig()

	logger := logger.SetupLogger(cfg.Env)

	logger.Info("starting application...")

	storage, err := pgsql.New(cfg.DatabaseUrl)

	fmt.Println(*storage, err)
}
