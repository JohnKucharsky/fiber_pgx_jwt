package domain

import (
	"time"
)

type CountryStore interface {
	Create(m CountryInput) (*Country, error)
	GetMany() ([]*Country, error)
	GetOne(id int) (*Country, error)
	Update(m CountryInput, id int) (*Country, error)
	Delete(id int) (*Country, error)
}

type Country struct {
	ID        int       `json:"id" db:"id"`
	Country   string    `json:"country" db:"country"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CountryInput struct {
	Country string `json:"country" validate:"required"`
}
