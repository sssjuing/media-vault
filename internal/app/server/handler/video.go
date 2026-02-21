package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/samber/lo"
	"github.com/sssjuing/media-vault/internal/app/server/model"
	"github.com/sssjuing/media-vault/internal/app/server/repository"
	"github.com/sssjuing/media-vault/internal/app/server/types"
	"github.com/sssjuing/media-vault/internal/app/server/utils"
)

// func (h *Handler) ListVideos(c echo.Context) error {
// 	tags := c.QueryParams()["tags"]
// 	videos, err := h.videoRepo.FindAll(repository.VidoesQueryOptions{Tags: tags})
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
// 	}
// 	list := commonUtils.Map(videos, func(v model.Video) types.VideoDTO {
// 		return *v.DTO()
// 	})
// 	return c.JSON(http.StatusOK, list)
// }

type paginateVideosRespone struct {
	Data  []types.VideoDTO `json:"data"`
	Total int64            `json:"total"`
}

func (h *Handler) SearchVideos(c *echo.Context) error {
	var opts repository.VidoesQueryOptions
	if code, err := validateRequest(c, &opts); err != nil {
		return c.JSON(code, utils.NewError(err))
	}
	videos, err := h.videoRepo.PaginateAll(opts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	count, err := h.videoRepo.Count(opts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	list := lo.Map(videos, func(v model.Video, _ int) types.VideoDTO {
		return *v.DTO()
	})
	return c.JSON(http.StatusOK, paginateVideosRespone{Data: list, Total: count})
}

func (h *Handler) CreateVideo(c *echo.Context) error {
	var cvd types.CreateVideoDTO
	if code, err := validateRequest(c, &cvd); err != nil {
		return c.JSON(code, utils.NewError(err))
	}
	vd, err := h.videoSvc.Create(&cvd)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, vd)
}

func (h *Handler) GetVideoById(c *echo.Context) error {
	v := c.Get("video").(*model.Video)
	return c.JSON(http.StatusOK, v.DTO())
}

func (h *Handler) UpdateVideo(c *echo.Context) error {
	v := c.Get("video").(*model.Video)
	var uvd types.UpdateVideoDTO
	if code, err := validateRequest(c, &uvd); err != nil {
		return c.JSON(code, utils.NewError(err))
	}
	vd, err := h.videoSvc.Update(v, &uvd)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, vd)
}

func (h *Handler) RemoveVideo(c *echo.Context) error {
	v := c.Get("video").(*model.Video)
	if err := h.videoRepo.Delete(v); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.NoContent(http.StatusNoContent)
}
