package main

import (
	"interview/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	println(cfg)
}
