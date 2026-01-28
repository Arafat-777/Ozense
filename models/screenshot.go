package models

import "gorm.io/gorm"

type Screenshot struct {
	gorm.Model
	MovieID uint   `json:"movie_id"`
	URL     string `json:"url"`
}
