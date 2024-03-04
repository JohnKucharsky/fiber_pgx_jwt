package domain

import (
	"time"
)

type CityStore interface {
	Create(m CityInput) (int, error)
	GetMany() ([]*City, error)
	GetOne(id int) (*City, error)
	Update(m CityInput, id int) (int, error)
	Delete(id int) (int, error)
}

type City struct {
	ID        int       `json:"id" db:"id"`
	City      string    `json:"city" db:"city"`
	Country   *Country  `json:"country" db:"country"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CityInput struct {
	City      string `json:"city" validate:"required"`
	CountryID int    `json:"country_id" validate:"required"`
}
