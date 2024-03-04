package domain

import (
	"time"
)

type ActorStore interface {
	Create(m ActorInput) (*Actor, error)
	GetMany() ([]*Actor, error)
	GetOne(id int) (*Actor, error)
	Update(m ActorInput, id int) (*Actor, error)
	Delete(id int) (*Actor, error)
}

type Actor struct {
	ID        int       `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ActorInput struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}
