package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v5"
	"github.com/samber/lo"
	"gorm.io/datatypes"

	"github.com/sssjuing/media-vault/internal/app/server/model"
	"github.com/sssjuing/media-vault/internal/app/server/utils"
)

type actressResponse struct {
	ID           uint             `json:"id"`
	UniqueName   string           `json:"unique_name"`
	ChineseName  string           `json:"chinese_name"`
	EnglishName  *string          `json:"english_name"`
	OtherNames   *json.RawMessage `json:"other_names"`
	BirthDate    *time.Time       `json:"birth_date"`
	BirthPlace   *string          `json:"birth_place"`
	Height       *int             `json:"height"`
	Weight       *int             `json:"weight"`
	Measurements *json.RawMessage `json:"measurements"`
	Cup          *string          `json:"cup"`
	BloodGroup   *string          `json:"blood_group"`
	DebutDate    *time.Time       `json:"debut_date"`
	Hobbies      *string          `json:"hobbies"`
	Notes        *string          `json:"notes"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
}

func newActressResponse(a *model.Actress) *actressResponse {
	if a == nil {
		return nil
	}
	var resp actressResponse
	_ = copier.Copy(&resp, a)
	resp.OtherNames = (*json.RawMessage)(a.OtherNames)
	resp.Measurements = (*json.RawMessage)(a.Measurements)
	return &resp
}

type actressCreateRequest struct {
	UniqueName   string           `json:"unique_name" validate:"required"`
	ChineseName  string           `json:"chinese_name" validate:"required"`
	EnglishName  *string          `json:"english_name"`
	OtherNames   *json.RawMessage `json:"other_names"`
	BirthDate    *time.Time       `json:"birth_date"`
	BirthPlace   *string          `json:"birth_place"`
	Height       *int             `json:"height"`
	Weight       *int             `json:"weight"`
	Measurements *json.RawMessage `json:"measurements"`
	Cup          *string          `json:"cup"`
	BloodGroup   *string          `json:"blood_group"`
	DebutDate    *time.Time       `json:"debut_date"`
	Hobbies      *string          `json:"hobbies"`
	Notes        *string          `json:"notes"`
}

type actressUpdateRequest struct {
	ID uint `json:"id"`
	actressCreateRequest
}

func (r *actressCreateRequest) toModel() *model.Actress {
	var m model.Actress
	_ = copier.Copy(&m, r)
	m.OtherNames = (*datatypes.JSON)(r.OtherNames)
	m.Measurements = (*datatypes.JSON)(r.Measurements)
	return &m
}

func (r *actressUpdateRequest) updateModel(a *model.Actress) {
	_ = copier.Copy(a, r)
	a.OtherNames = (*datatypes.JSON)(r.OtherNames)
	a.Measurements = (*datatypes.JSON)(r.Measurements)
}

func (h *Handler) ListActresses(c *echo.Context) error {
	actresses, err := h.actressRepo.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, lo.Map(actresses, func(a model.Actress, _ int) *actressResponse {
		return newActressResponse(&a)
	}))
}

func (h *Handler) CreateActress(c *echo.Context) error {
	var req actressCreateRequest
	if err := validateRequest(c, &req); err != nil {
		return err
	}
	a := req.toModel()
	err := h.actressRepo.Create(a)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newActressResponse(a))
}

func (h *Handler) GetActressById(c *echo.Context) error {
	a := c.Get("actress").(*model.Actress)
	return c.JSON(http.StatusOK, newActressResponse(a))
}

func (h *Handler) UpdateActress(c *echo.Context) error {
	a := c.Get("actress").(*model.Actress)
	var req actressUpdateRequest
	if err := validateRequest(c, &req); err != nil {
		return err
	}
	req.updateModel(a)
	if err := h.actressRepo.Update(a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newActressResponse(a))
}

func (h *Handler) DeleteActress(c *echo.Context) error {
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
	list := lo.Map(videos, func(v *model.Video, _ int) *videoResponse {
		return newVideoResponse(v)
	})
	return c.JSON(http.StatusOK, list)
}
