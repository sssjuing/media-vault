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
	"github.com/sssjuing/media-vault/internal/app/server/repository"
	"github.com/sssjuing/media-vault/internal/app/server/utils"
	"github.com/sssjuing/media-vault/internal/pkg/config"
)

var publicUrl string

func init() {
	publicUrl = config.GetMinioPublicUrl()
}

type videoResponse struct {
	ID           uint               `json:"id"`
	SerialNumber string             `json:"serial_number"`
	CoverUrl     string             `json:"cover_url"`
	Title        *string            `json:"title"`
	ChineseTitle *string            `json:"chinese_title"`
	Actresses    []*actressResponse `json:"actresses"`
	ReleaseDate  *time.Time         `json:"release_date"`
	VideoUrl     *string            `json:"video_url"`
	Mosaic       *bool              `json:"mosaic"`
	Tags         *json.RawMessage   `json:"tags"`
	M3U8Url      *string            `json:"m3u8_url"`
	Synopsis     *string            `json:"synopsis"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}

func newVideoResponse(v *model.Video) *videoResponse {
	if v == nil {
		return nil
	}
	var resp videoResponse
	_ = copier.Copy(&resp, v)
	resp.CoverUrl = publicUrl + v.CoverPath
	if v.VideoPath != nil && *v.VideoPath != "" {
		videoUrl := publicUrl + *v.VideoPath
		resp.VideoUrl = &videoUrl
	}
	resp.Actresses = lo.Map(v.Actresses, func(a *model.Actress, _ int) *actressResponse {
		return newActressResponse(a)
	})
	resp.Tags = (*json.RawMessage)(v.Tags)
	resp.M3U8Url = v.M3U8URL
	return &resp
}

type actressIdRequest struct {
	ID uint `json:"id" validate:"required"`
}

type videoCreateRequest struct {
	SerialNumber string              `json:"serial_number" validate:"required"`
	CoverPath    string              `json:"cover_path" validate:"required"`
	Title        *string             `json:"title"`
	ChineseTitle *string             `json:"chinese_title"`
	Actresses    []*actressIdRequest `json:"actresses"`
	ReleaseDate  *time.Time          `json:"release_date"`
	VideoPath    *string             `json:"video_path"`
	Mosaic       *bool               `json:"mosaic"`
	Tags         *json.RawMessage    `json:"tags"`
	M3U8Url      *string             `json:"m3u8_url"`
	Synopsis     *string             `json:"synopsis"`
}

type videoUpdateRequest struct {
	SerialNumber string              `json:"serial_number"`
	CoverPath    string              `json:"cover_path"`
	Title        *string             `json:"title"`
	ChineseTitle *string             `json:"chinese_title"`
	Actresses    []*actressIdRequest `json:"actresses"`
	ReleaseDate  *time.Time          `json:"release_date"`
	VideoPath    *string             `json:"video_path"`
	Mosaic       *bool               `json:"mosaic"`
	Tags         *json.RawMessage    `json:"tags"`
	M3U8Url      *string             `json:"m3u8_url"`
	Synopsis     *string             `json:"synopsis"`
}

func (r *videoCreateRequest) toModel() *model.Video {
	var m model.Video
	_ = copier.Copy(&m, r)
	m.Actresses = lo.Map(r.Actresses, func(ad *actressIdRequest, _ int) *model.Actress {
		actress := &model.Actress{}
		actress.ID = ad.ID
		return actress
	})
	m.Tags = (*datatypes.JSON)(r.Tags)
	m.M3U8URL = r.M3U8Url
	return &m
}

func (r *videoUpdateRequest) updateModel(v *model.Video) {
	if r.SerialNumber != "" {
		v.SerialNumber = r.SerialNumber
	}
	if r.CoverPath != "" {
		v.CoverPath = r.CoverPath
	}
	_ = copier.Copy(v, r)
	if r.Actresses != nil {
		v.Actresses = lo.Map(r.Actresses, func(ad *actressIdRequest, _ int) *model.Actress {
			actress := &model.Actress{}
			actress.ID = ad.ID
			return actress
		})
	}
	v.Tags = (*datatypes.JSON)(r.Tags)
	v.M3U8URL = r.M3U8Url
}

type paginateVideosResponse struct {
	Data  []*videoResponse `json:"data"`
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
	list := lo.Map(videos, func(v model.Video, _ int) *videoResponse {
		return newVideoResponse(&v)
	})
	return c.JSON(http.StatusOK, paginateVideosResponse{Data: list, Total: count})
}

func (h *Handler) CreateVideo(c *echo.Context) error {
	var req videoCreateRequest
	if code, err := validateRequest(c, &req); err != nil {
		return c.JSON(code, utils.NewError(err))
	}
	v := req.toModel()
	err := h.videoRepo.Create(v)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newVideoResponse(v))
}

func (h *Handler) GetVideoById(c *echo.Context) error {
	v := c.Get("video").(*model.Video)
	return c.JSON(http.StatusOK, newVideoResponse(v))
}

func (h *Handler) UpdateVideo(c *echo.Context) error {
	v := c.Get("video").(*model.Video)
	var req videoUpdateRequest
	if code, err := validateRequest(c, &req); err != nil {
		return c.JSON(code, utils.NewError(err))
	}
	req.updateModel(v)
	if err := h.videoRepo.Update(v); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newVideoResponse(v))
}

func (h *Handler) RemoveVideo(c *echo.Context) error {
	v := c.Get("video").(*model.Video)
	if err := h.videoRepo.Delete(v); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.NoContent(http.StatusNoContent)
}
