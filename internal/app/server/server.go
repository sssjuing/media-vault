package server

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/sssjuing/media-vault/internal/app/server/db"
	"github.com/sssjuing/media-vault/internal/app/server/download"
	"github.com/sssjuing/media-vault/internal/app/server/handler"
	"github.com/sssjuing/media-vault/internal/app/server/repository"
	"github.com/sssjuing/media-vault/internal/app/server/router"
	"github.com/sssjuing/media-vault/internal/pkg/config"
)

func Run() {
	r := router.New()
	// r.GET("/swagger/*", echoSwagger.WrapHandler)

	download.Init(r.Logger)

	dsn := config.GetPostgresDsn()
	replicas := config.GetPostgresReplicas()
	d := db.NewDB(dsn, replicas)

	ar := repository.NewActressRepositoryImpl(d)
	vr := repository.NewVideoRepositoryImpl(d)
	vtr := repository.NewVideoTagRepositoryImpl(d)

	h := handler.NewHandler(ar, vr, vtr)
	v1 := r.Group("/api")
	h.Register(v1)

	r.GET("/", func(c *echo.Context) error {
		ua := strings.ToLower(c.Request().UserAgent())
		isMobile := strings.Contains(ua, "mobile") ||
			strings.Contains(ua, "android") ||
			strings.Contains(ua, "iphone") ||
			strings.Contains(ua, "ipad")
		if isMobile {
			return c.Redirect(http.StatusSeeOther, "/h5")
		}
		return c.Redirect(http.StatusSeeOther, "/console")
	})

	if err := r.Start(":1323"); err != nil {
		r.Logger.Error("failed to start server", "error", err)
	}
}
