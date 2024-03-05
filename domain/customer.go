package domain

import (
	"time"
)

type CustomerStore interface {
	Create(m CustomerInput) (*CustomerDB, error)
	GetMany() ([]*CustomerDB, error)
	GetOne(id int) (*CustomerDB, error)
	Update(m CustomerInput, id int) (*CustomerDB, error)
	Delete(id int) (*CustomerDB, error)
}

type Customer struct {
	ID         int       `json:"id"`
	StoreID    int       `json:"store_id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      *string   `json:"email"`
	Address    *Address  `json:"address"`
	Active     bool      `json:"active"`
	CreateDate string    `json:"create_date"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CustomerDB struct {
	ID         int       `db:"id"`
	StoreID    int       `db:"store_id"`
	FirstName  string    `db:"first_name"`
	LastName   string    `db:"last_name"`
	Email      *string   `db:"email"`
	AddressID  *int      `db:"address_id"`
	Active     bool      `db:"active"`
	CreateDate string    `db:"create_date"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type CustomerInput struct {
	StoreID    int     `json:"store_id" validate:"required"`
	FirstName  string  `json:"first_name" validate:"required"`
	LastName   string  `json:"last_name" validate:"required"`
	Email      *string `json:"email" validate:"omitempty,email"`
	AddressID  *int    `json:"address_id"`
	Active     bool    `json:"active"`
	CreateDate string  `json:"create_date" validate:"required"`
}

func CustomerDBtoCustomer(cusDB *CustomerDB, addr *Address) Customer {
	var address *Address

	if cusDB.AddressID != nil {
		address = &Address{
			ID:         addr.ID,
			Address:    addr.Address,
			Address2:   addr.Address2,
			District:   addr.District,
			City:       addr.City,
			Country:    addr.Country,
			PostalCode: addr.PostalCode,
			Phone:      addr.Phone,
			UpdatedAt:  addr.UpdatedAt,
		}
	}

	return Customer{
		ID:         cusDB.ID,
		StoreID:    cusDB.StoreID,
		FirstName:  cusDB.FirstName,
		LastName:   cusDB.LastName,
		Email:      cusDB.Email,
		Address:    address,
		Active:     cusDB.Active,
		CreateDate: cusDB.CreateDate,
		UpdatedAt:  cusDB.UpdatedAt,
	}
}
