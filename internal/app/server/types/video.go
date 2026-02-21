package types

import (
	"encoding/json"
	"time"
)

type VideoDTO struct {
	ID           uint             `json:"id"`
	SerialNumber string           `json:"serial_number"`
	CoverUrl     string           `json:"cover_url"` // 封面完整 URL
	Title        *string          `json:"title"`
	ChineseTitle *string          `json:"chinese_title"`
	Actresses    []*ActressDTO    `json:"actresses"`
	ReleaseDate  *time.Time       `json:"release_date"`
	VideoUrl     *string          `json:"video_url"` // 视频完整 URL
	Mosaic       *bool            `json:"mosaic"`
	Tags         *json.RawMessage `json:"tags"`
	M3U8Url      *string          `json:"m3u8_url"`
	Synopsis     *string          `json:"synopsis"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
}

type CreateVideoDTO struct {
	SerialNumber string           `json:"serial_number" validate:"required"`
	CoverPath    string           `json:"cover_path" validate:"required"`
	Title        *string          `json:"title"`
	ChineseTitle *string          `json:"chinese_title"`
	Actresses    []*ActressDTO    `json:"actresses"`
	ReleaseDate  *time.Time       `json:"release_date"`
	VideoPath    *string          `json:"video_path"`
	Mosaic       *bool            `json:"mosaic"`
	Tags         *json.RawMessage `json:"tags"`
	M3U8Url      *string          `json:"m3u8_url"`
	Synopsis     *string          `json:"synopsis"`
}

type UpdateVideoDTO struct {
	SerialNumber string           `json:"serial_number"`
	CoverPath    string           `json:"cover_path"`
	Title        *string          `json:"title"`
	ChineseTitle *string          `json:"chinese_title"`
	Actresses    []*ActressDTO    `json:"actresses"`
	ReleaseDate  *time.Time       `json:"release_date"`
	VideoPath    *string          `json:"video_path"`
	Mosaic       *bool            `json:"mosaic"`
	Tags         *json.RawMessage `json:"tags"`
	M3U8Url      *string          `json:"m3u8_url"`
	Synopsis     *string          `json:"synopsis"`
}
