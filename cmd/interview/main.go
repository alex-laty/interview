package main

import (
	"interview/internal/config"
)

func main() {
	// TODO: init config: cleanenv
	cfg := config.LoadConfig()

	println(cfg)

	// TODO: init logger: slog

	// TODO: init storage: pgsql

	// TODO: init router: chi

	// TODO: init server
}
