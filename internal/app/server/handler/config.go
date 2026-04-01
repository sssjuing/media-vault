package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/sssjuing/media-vault/internal/app/server/model"
)

func (h *Handler) GetVideoTags(c *echo.Context) error {
	var globalSetting model.GlobalSetting
	h.db.Where("id = ?", true).First(&globalSetting)
	if globalSetting.VideoTags == nil {
		return c.JSON(http.StatusOK, []string{})
	}
	if globalSetting.VideoTags != nil {
		return c.JSON(http.StatusOK, globalSetting.VideoTags)
	}
	return c.JSON(http.StatusOK, globalSetting.VideoTags)
}

func (h *Handler) SetVideoTags(c *echo.Context) error {
	var globalSetting model.GlobalSetting
	h.db.Where("id = ?", true).First(&globalSetting)
	// globalSetting.VideoTags = datatypes.JSONSlice[string](c.FormValue("video_tags"))
	h.db.Save(&globalSetting)
	return c.JSON(http.StatusOK, globalSetting.VideoTags)
}
