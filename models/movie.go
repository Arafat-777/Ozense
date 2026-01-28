package models

type Movie struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Trailer     string `json:"trailer"`
	AgeID       uint   `json:"age_id"`
	Age         Age    `json:"age"`

	Genres     []Genre    `gorm:"many2many:movie_genres" json:"genres"`
	Categories []Category `gorm:"many2many:movie_categories" json:"categories"`

	Seasons     []Season     `json:"seasons"`
	Screenshots []Screenshot `json:"screenshots"`
}
