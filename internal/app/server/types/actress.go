package types

import (
	"encoding/json"
	"time"
)

type ActressDTO struct {
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

type CreateActressDTO struct {
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

type UpdateActressDTO struct {
	CreateActressDTO
}
