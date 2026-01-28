package models

type MovieInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
	Trailer     string `json:"trailer"`
	AgeID       uint   `json:"age_id"`
	GenreIDs    []uint `json:"genre_ids"`
	CategoryIDs []uint `json:"category_ids"`
}
