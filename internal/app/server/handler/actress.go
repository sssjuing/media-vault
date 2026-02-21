package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/samber/lo"

	"github.com/sssjuing/media-vault/internal/app/server/model"
	"github.com/sssjuing/media-vault/internal/app/server/types"
	"github.com/sssjuing/media-vault/internal/app/server/utils"
)

func (h *Handler) ListActresses(c *echo.Context) error {
	list, err := h.actressSvc.GetList()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, list)
}

func (h *Handler) CreateActress(c *echo.Context) error {
	var cad types.CreateActressDTO
	if err := c.Bind(&cad); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	ad, err := h.actressSvc.Create(&cad)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, ad)
}

func (h *Handler) GetActressById(c *echo.Context) error {
	a := c.Get("actress").(*model.Actress)
	return c.JSON(http.StatusOK, a.DTO())
}

func (h *Handler) UpdateActress(c *echo.Context) error {
	a := c.Get("actress").(*model.Actress)
	var uad types.UpdateActressDTO
	if err := c.Bind(&uad); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	ad, err := h.actressSvc.Update(a, &uad)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, ad)
}

func (h *Handler) RemoveActress(c *echo.Context) error {
	a := c.Get("actress").(*model.Actress)
	if err := h.actressRepo.Delete(a); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *Handler) GetVideosByActressId(c *echo.Context) error {
	a := c.Get("actress").(*model.Actress)
	videos, err := h.actressRepo.FindVideos(a.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	list := lo.Map(videos, func(v *model.Video, _ int) types.VideoDTO {
		return *v.DTO()
	})
	return c.JSON(http.StatusOK, list)
}
