package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/sssjuing/media-vault/internal/app/server/middleware"
)

func (h *Handler) Register(g *echo.Group) {
	g.POST("/user/login", h.Login)

	g.Use(middleware.JWT())
	g.GET("/user/whoami", h.CurrentUser)

	files := g.Group("/files")
	files.GET("", h.ListFiles)
	files.GET("/:folder_name", h.ListFiles)
	files.POST("/upload-image", h.UploadImage)
	files.POST("/upload-local-file", h.UploadLocalFile)

	g.GET("/actresses", h.ListActresses)
	g.POST("/actresses", h.CreateActress)
	actress := g.Group("/actresses/:actress_id", middleware.ActressChecker(h.actressRepo))
	actress.GET("", h.GetActressById)
	actress.PUT("", h.UpdateActress)
	actress.DELETE("", h.DeleteActress)
	actress.GET("/videos", h.GetVideosByActressId)

	// g.GET("/videos", h.ListVideos)
	g.POST("/videos\\:search", h.SearchVideos)
	g.POST("/videos", h.CreateVideo)
	video := g.Group("/videos/:video_id", middleware.VideoChecker(h.videoRepo))
	video.GET("", h.GetVideoById)
	video.PUT("", h.UpdateVideo)
	video.DELETE("", h.RemoveVideo)

	g.GET("/resources", h.ListResources)
	g.DELETE("/resources/:resource_id", h.DeleteResource)
	g.GET("/resources/:resource_id/segments", h.ListSegments)
	g.POST("/resources/:resource_id/download", h.DownloadResource)
	g.POST("/download", h.SubmitDownload)
	g.GET("/download/files", h.ListDownloadedFiles)

	g.GET("/*", func(c *echo.Context) error {
		return c.String(http.StatusNotFound, "not found")
	})
}
