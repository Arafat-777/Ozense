package models

import "gorm.io/gorm"

type Season struct {
	gorm.Model
	MovieID  uint      `json:"movie_id"`
	Number   int       `json:"number"`
	Name     string    `json:"name"`
	Episodes []Episode `gorm:"foreignKey:SeasonID" json:"episodes"`
}
