package domain

import (
	"time"
)

type CategoryStore interface {
	Create(m CategoryInput) (*Category, error)
	GetMany() ([]*Category, error)
	GetOne(id int) (*Category, error)
	Update(m CategoryInput, id int) (*Category, error)
	Delete(id int) (*Category, error)
}

type Category struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CategoryInput struct {
	Name string `json:"name" validate:"required"`
}
