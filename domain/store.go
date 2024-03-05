package domain

import (
	"time"
)

type StoreStore interface {
	Create(m StoreInput) (*Store, error)
	GetMany() ([]*Store, error)
	GetOne(id int) (*Store, error)
	Update(m StoreInput, id int) (*Store, error)
	Delete(id int) (*Store, error)
}

type Store struct {
	ID        int       `json:"id"`
	Manager   string    `json:"manager"`
	Address   Address   `json:"address"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StoreInput struct {
	ManagerID string `json:"manager_id" validate:"required"`
	AddressID string `json:"address_id" validate:"required"`
}
