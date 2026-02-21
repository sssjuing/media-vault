package setup

import (
	"fmt"
	"log/slog"

	"github.com/sssjuing/media-vault/internal/app/server/db"
	"github.com/sssjuing/media-vault/internal/pkg/config"
)

func Run() {
	logger := slog.Default()
	dsn := config.GetPostgresDsn()
	d := db.NewDB(dsn, config.GetPostgresReplicas())
	if err := db.AutoMigrate(d); err != nil {
		logger.Error(fmt.Sprintf("auto migrate failed: %s", err.Error()))
	} else {
		logger.Info("setup successfully")
	}
}
