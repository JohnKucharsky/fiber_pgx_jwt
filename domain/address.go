package domain

import (
	"time"
)

type AddressStore interface {
	Create(m AddressInput) (int, error)
	GetMany() ([]*Address, error)
	GetOne(id int) (*Address, error)
	Update(m AddressInput, id int) (int, error)
	Delete(id int) (int, error)
}

type Address struct {
	ID         int       `json:"id" db:"id"`
	Address    string    `json:"address" db:"address"`
	Address2   *string   `json:"address2" db:"address2"`
	District   string    `json:"district" db:"district"`
	City       *City     `json:"city" db:"city"`
	Country    *Country  `json:"country" db:"country"`
	PostalCode *int      `json:"postal_code" db:"postal_code"`
	Phone      string    `json:"phone" db:"phone"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

type AddressInput struct {
	Address    string  `json:"address" validate:"required"`
	Address2   *string `json:"address2"`
	District   string  `json:"district" validate:"required"`
	CityID     *int    `json:"city_id"`
	PostalCode *int    `json:"postal_code"`
	Phone      string  `json:"phone" validate:"required"`
}
