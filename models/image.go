package models

import (
	"time"
)

type Image struct {
	// ID          string    `json:"image_id,omitempty"`
	URL         string    `json:"image_url" validate:"required"`
	W           int       `json:"width" validate:"required"`
	H           int       `json:"height" validate:"required"`
	AspectRatio float64   `json:"aspect_ratio" validate:"required"`
	SizeKB      int       `json:"size_kb" validate:"required"`
	UploadTime  time.Time `json:"upload_time" validate:"required"`
	UploadBy    string    `json:"upload_by" validate:"required"`
}
