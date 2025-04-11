package logger

import (
	"log/slog"
	"os"
)

func SetupLogger(env string) *slog.Logger {
	// there may be different types of handlers or levels depending on the enviiroment
	// Example: if env == prod info
	logger := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	return logger
}
