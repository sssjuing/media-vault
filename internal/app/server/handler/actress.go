package handler

import (
	"net/http"
	"time"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v5"
	"github.com/samber/lo"

	"github.com/sssjuing/media-vault/internal/app/server/model"
	"github.com/sssjuing/media-vault/internal/app/server/utils"
)

type nameResponse struct {
	Name     string `json:"name"`
	NameType string `json:"name_type"`
}

type nameSetResponse struct {
	ID    uint            `json:"id"`
	Names []*nameResponse `json:"names"`
}

type actressResponse struct {
	ID         uint               `json:"id"`
	UniqueName string             `json:"unique_name"`
	NameSets   []*nameSetResponse `json:"name_sets"`
	BirthDate  *time.Time         `json:"birth_date"`
	BirthPlace *string            `json:"birth_place"`
	Height     *int               `json:"height"`
	Weight     *int               `json:"weight"`
	BWH        *model.BWH         `json:"bwh"`
	Cup        *string            `json:"cup"`
	BloodGroup *string            `json:"blood_group"`
	DebutDate  *time.Time         `json:"debut_date"`
	Hobbies    *string            `json:"hobbies"`
	Notes      *string            `json:"notes"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

func newActressResponse(a *model.Actress) *actressResponse {
	if a == nil {
		return nil
	}
	var resp actressResponse
	_ = copier.Copy(&resp, a)

	if a.NameSets != nil {
		resp.NameSets = lo.Map(a.NameSets, func(ns *model.ActressNameSet, _ int) *nameSetResponse {
			nsResp := &nameSetResponse{
				ID: ns.ID,
			}
			if ns.Names != nil {
				nsResp.Names = lo.Map(ns.Names, func(n *model.ActressName, _ int) *nameResponse {
					return &nameResponse{
						Name:     n.Name,
						NameType: string(n.NameType),
					}
				})
			}
			return nsResp
		})
	}

	return &resp
}

type nameRequest struct {
	Name     string `json:"name" validate:"required"`
	NameType string `json:"name_type" validate:"required"`
}

type nameSetRequest struct {
	Names []*nameRequest `json:"names" validate:"required,dive"`
}

type bwhRequest struct {
	Bust  *int `json:"bust"`
	Waist *int `json:"waist"`
	Hips  *int `json:"hips"`
}

type actressCreateRequest struct {
	UniqueName string            `json:"unique_name" validate:"required"`
	NameSets   []*nameSetRequest `json:"name_sets" validate:"required,dive"`
	BirthDate  *time.Time        `json:"birth_date"`
	BirthPlace *string           `json:"birth_place"`
	Height     *int              `json:"height"`
	Weight     *int              `json:"weight"`
	BWH        *bwhRequest       `json:"bwh"`
	Cup        *string           `json:"cup"`
	BloodGroup *string           `json:"blood_group"`
	DebutDate  *time.Time        `json:"debut_date"`
	Hobbies    *string           `json:"hobbies"`
	Notes      *string           `json:"notes"`
}

type actressUpdateRequest struct {
	ID uint `json:"id"`
	actressCreateRequest
}

func (r *actressCreateRequest) toModel() *model.Actress {
	var m model.Actress
	_ = copier.Copy(&m, r)

	if r.NameSets != nil {
		m.NameSets = lo.Map(r.NameSets, func(nsr *nameSetRequest, _ int) *model.ActressNameSet {
			ns := &model.ActressNameSet{}
			if nsr.Names != nil {
				ns.Names = lo.Map(nsr.Names, func(nr *nameRequest, _ int) *model.ActressName {
					return &model.ActressName{
						Name:     nr.Name,
						NameType: model.NameType(nr.NameType),
					}
				})
			}
			return ns
		})
	}

	if r.BWH != nil {
		m.BWH = &model.BWH{
			Bust:  r.BWH.Bust,
			Waist: r.BWH.Waist,
			Hips:  r.BWH.Hips,
		}
	}

	return &m
}

func (r *actressUpdateRequest) updateModel(a *model.Actress) {
	_ = copier.Copy(a, r)

	if r.NameSets != nil {
		a.NameSets = lo.Map(r.NameSets, func(nsr *nameSetRequest, _ int) *model.ActressNameSet {
			ns := &model.ActressNameSet{}
			if nsr.Names != nil {
				ns.Names = lo.Map(nsr.Names, func(nr *nameRequest, _ int) *model.ActressName {
					return &model.ActressName{
						Name:     nr.Name,
						NameType: model.NameType(nr.NameType),
					}
				})
			}
			return ns
		})
	}

	if r.BWH != nil {
		a.BWH = &model.BWH{
			Bust:  r.BWH.Bust,
			Waist: r.BWH.Waist,
			Hips:  r.BWH.Hips,
		}
	}
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
