package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/sssjuing/media-vault/internal/app/server/model"
	"github.com/sssjuing/media-vault/internal/app/server/repository"
	"github.com/sssjuing/media-vault/internal/app/server/utils"
	"github.com/sssjuing/media-vault/internal/pkg/config"
)

var publicUrl string

func init() {
	publicUrl = config.GetMinioPublicUrl()
}

// func newVideoResponse(v *model.Video) *videoResponse {
// 	if v == nil {
// 		return nil
// 	}
// 	var resp videoResponse
// 	_ = copier.Copy(&resp, v)
// 	resp.CoverUrl = publicUrl + v.CoverPath
// 	if v.VideoPath != nil && *v.VideoPath != "" {
// 		videoUrl := publicUrl + *v.VideoPath
// 		resp.VideoUrl = &videoUrl
// 	}

// 	if v.ActressNameSets != nil {
// 		resp.ActressNameSets = lo.Map(v.ActressNameSets, func(ns *model.ActressNameSet, _ int) *nameSetInVideoResponse {
// 			nsResp := &nameSetInVideoResponse{
// 				ID: ns.ID,
// 			}
// 			if ns.Names != nil {
// 				nsResp.Names = lo.Map(ns.Names, func(n *model.ActressName, _ int) *nameResponse {
// 					return &nameResponse{
// 						Name:     n.Name,
// 						NameType: string(n.NameType),
// 					}
// 				})
// 			}
// 			if ns.Actress != nil {
// 				nsResp.Actress = newActressResponse(ns.Actress)
// 			}
// 			return nsResp
// 		})
// 	}

// 	resp.Tags = (*json.RawMessage)(v.Tags)
// 	return &resp
// }

// type actressNameSetIdRequest struct {
// 	ID uint `json:"id" validate:"required"`
// }

// type videoCreateRequest struct {
// 	SerialNumber    string                     `json:"serial_number" validate:"required"`
// 	CoverPath       string                     `json:"cover_path" validate:"required"`
// 	Title           *string                    `json:"title"`
// 	ChineseTitle    *string                    `json:"chinese_title"`
// 	ActressNameSets []*actressNameSetIdRequest `json:"actress_name_sets"`
// 	ReleaseDate     *time.Time                 `json:"release_date"`
// 	VideoPath       *string                    `json:"video_path"`
// 	Mosaic          *bool                      `json:"mosaic"`
// 	Tags            *json.RawMessage           `json:"tags"`
// 	M3u8Url         *string                    `json:"m3u8_url"`
// 	M3u8Referer     *string                    `json:"m3u8_referer"`
// 	Synopsis        *string                    `json:"synopsis"`
// }

// type videoUpdateRequest struct {
// 	SerialNumber    string                     `json:"serial_number"`
// 	CoverPath       string                     `json:"cover_path"`
// 	Title           *string                    `json:"title"`
// 	ChineseTitle    *string                    `json:"chinese_title"`
// 	ActressNameSets []*actressNameSetIdRequest `json:"actress_name_sets"`
// 	ReleaseDate     *time.Time                 `json:"release_date"`
// 	VideoPath       *string                    `json:"video_path"`
// 	Mosaic          *bool                      `json:"mosaic"`
// 	Tags            *json.RawMessage           `json:"tags"`
// 	M3u8Url         *string                    `json:"m3u8_url"`
// 	M3u8Referer     *string                    `json:"m3u8_referer"`
// 	Synopsis        *string                    `json:"synopsis"`
// }

// func (r *videoCreateRequest) toModel() *model.Video {
// 	var m model.Video
// 	_ = copier.Copy(&m, r)

// 	if r.ActressNameSets != nil {
// 		m.ActressNameSets = lo.Map(r.ActressNameSets, func(nsr *actressNameSetIdRequest, _ int) *model.ActressNameSet {
// 			ns := &model.ActressNameSet{}
// 			ns.ID = nsr.ID
// 			return ns
// 		})
// 	}

// 	m.Tags = (*datatypes.JSON)(r.Tags)
// 	return &m
// }

// func (r *videoUpdateRequest) updateModel(v *model.Video) {
// 	if r.SerialNumber != "" {
// 		v.SerialNumber = r.SerialNumber
// 	}
// 	if r.CoverPath != "" {
// 		v.CoverPath = r.CoverPath
// 	}
// 	_ = copier.Copy(v, r)

// 	if r.ActressNameSets != nil {
// 		v.ActressNameSets = lo.Map(r.ActressNameSets, func(nsr *actressNameSetIdRequest, _ int) *model.ActressNameSet {
// 			ns := &model.ActressNameSet{}
// 			ns.ID = nsr.ID
// 			return ns
// 		})
// 	}

// 	v.Tags = (*datatypes.JSON)(r.Tags)
// }

type paginateVideosResponse struct {
	Data  []model.Video `json:"data"`
	Total int64         `json:"total"`
}

func (h *Handler) SearchVideos(c *echo.Context) error {
	var opts repository.VidoesQueryOptions
	if err := validateRequest(c, &opts); err != nil {
		return err
	}
	videos, err := h.videoRepo.PaginateAll(opts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	count, err := h.videoRepo.Count(opts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, paginateVideosResponse{Data: videos, Total: count})
}

func (h *Handler) CreateVideo(c *echo.Context) error {
	var req videoCreateRequest
	if err := validateRequest(c, &req); err != nil {
		return err
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
	if err := validateRequest(c, &req); err != nil {
		return err
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
