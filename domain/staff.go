package domain

import (
	"time"
)

type StaffStore interface {
	Create(m StaffInput) (*StaffDB, error)
	GetMany() ([]*StaffDB, error)
	GetOne(id int) (*StaffDB, error)
	Update(m StaffInput, id int) (*StaffDB, error)
	Delete(id int) (*StaffDB, error)
}

type Staff struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     *string   `json:"email"`
	StoreID   *int      `json:"store_id"`
	Active    bool      `json:"active"`
	Username  string    `json:"username"`
	Password  *string   `json:"password"`
	Picture   *[]byte   `json:"picture"`
	Address   *Address  `json:"address"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StaffDB struct {
	ID        int       `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     *string   `db:"email"`
	StoreID   *int      `db:"store_id"`
	Active    bool      `db:"active"`
	Username  string    `db:"username"`
	Password  *string   `db:"password"`
	Picture   *[]byte   `db:"picture"`
	AddressID *int      `db:"address_id"`
	UpdatedAt time.Time `db:"updated_at"`
}

type StaffInput struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     *string `json:"email"`
	StoreID   *int    `json:"store_id"`
	Active    bool    `json:"active"`
	Username  string  `json:"username"`
	Password  *string `json:"password"`
	Picture   *[]byte `json:"picture"`
	AddressID *int    `json:"address_id"`
}

func StaffDBtoStaff(stfDB *StaffDB, addr *Address) Staff {
	var address *Address

	if stfDB.AddressID != nil {
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

	return Staff{
		ID:        stfDB.ID,
		FirstName: stfDB.FirstName,
		LastName:  stfDB.LastName,
		Email:     stfDB.Email,
		StoreID:   stfDB.StoreID,
		Active:    stfDB.Active,
		Username:  stfDB.Username,
		Password:  stfDB.Password,
		Picture:   stfDB.Picture,
		Address:   address,
		UpdatedAt: stfDB.UpdatedAt,
	}
}
