package domain

import (
	"time"
)

type FilmStore interface {
	Create(m FilmInput) (*Film, error)
	GetMany() ([]*Film, error)
	GetOne(id int) (*Film, error)
	Update(m FilmInput, id int) (*Film, error)
	Delete(id int) (*Film, error)
}

type Film struct {
	ID              int       `json:"id" db:"id"`
	Title           string    `json:"title" db:"title"`
	Description     *string   `json:"description" db:"description"`
	ReleaseYear     *int8     `json:"release_year" db:"release_year"`
	Language        Language  `json:"language_id" db:"language_id"`
	RentalDuration  int       `json:"rental_duration" db:"rental_duration"`
	RentalRate      int       `json:"rental_rate" db:"rental_rate"`
	Length          int       `json:"length" db:"length"`
	ReplacementCost int       `json:"replacement_cost" db:"replacement_cost"`
	Rating          string    `json:"rating" db:"rating"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

type FilmInput struct {
	Title           string  `json:"title" validate:"required"`
	Description     *string `json:"description" `
	ReleaseYear     *int8   `json:"release_year" `
	LanguageID      int     `json:"language_id" validate:"required"`
	RentalDuration  *int    `json:"rental_duration" validate:"required"`
	RentalRate      *int    `json:"rental_rate" validate:"required"`
	Length          *int    `json:"length" validate:"length"`
	ReplacementCost *int    `json:"replacement_cost" validate:"replacement_cost"`
	Rating          *string `json:"rating" validate:"rating"`
}
