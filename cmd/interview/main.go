package main

import (
	"interview/internal/config"
	"interview/internal/logger"
)

func main() {
	cfg := config.LoadConfig()

	logger := logger.SetupLogger(cfg.Env)

	logger.Info("starting application...");
}