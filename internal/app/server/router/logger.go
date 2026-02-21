package router

import (
	"log/slog"
	"os"

	"github.com/sssjuing/media-vault/internal/pkg/config"
)

func newLogger() *slog.Logger {
	var level slog.Level
	cfg := config.GetConfig()
	switch cfg.GetString("server.log_level") {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})
	return slog.New(handler)
}
