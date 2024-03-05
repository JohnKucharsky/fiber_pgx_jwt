package domain

import (
	"time"
)

type LanguageStore interface {
	Create(m LanguageInput) (*Language, error)
	GetMany() ([]*Language, error)
	GetOne(id int) (*Language, error)
	Update(m LanguageInput, id int) (*Language, error)
	Delete(id int) (*Language, error)
}

type Language struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type LanguageInput struct {
	Name string `json:"name" validate:"required"`
}
