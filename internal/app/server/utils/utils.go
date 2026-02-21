package utils

// import (
// 	"fmt"
// 	"log/slog"

// 	"github.com/labstack/gommon/log"
// 	"github.com/sssjuing/media-vault/internal/pkg/config"
// )

// func GetLogLevel() *slog.Level {
// 	cfg := config.GetConfig()
// 	level := cfg.GetString("server.log_level")
// 	fmt.Println("log level is", level)

// 	switch level {
// 	case "debug":
// 		return log.DEBUG
// 	case "info":
// 		return log.INFO
// 	case "warn":
// 		return log.WARN
// 	case "error":
// 		return log.ERROR
// 	default:
// 		return log.INFO
// 	}
// }
