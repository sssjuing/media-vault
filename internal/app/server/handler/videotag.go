package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/sssjuing/media-vault/internal/app/server/utils"
)

func (h *Handler) ListVideoTags(c *echo.Context) error {
	tags, err := h.videoTagRepo.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, tags)
}
