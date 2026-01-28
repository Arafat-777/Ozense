package models

import "gorm.io/gorm"

type Episode struct {
	gorm.Model
	SeasonID uint   `json:"season_id"`
	Title    string `json:"title"`
	VideoURL string `json:"video_url"`
	Duration string `json:"duration"`
}
