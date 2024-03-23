package domain

import (
	"time"
)

type StoreStore interface {
	Create(m StoreInput) (*StoreDB, error)
	GetMany() ([]*StoreDBSecondVer, error)
	GetOne(id int) (*StoreDB, error)
	Update(m StoreInput, id int) (*StoreDB, error)
	Delete(id int) (*StoreDB, error)
}

type Store struct {
	ID        int       `json:"id"`
	Manager   Staff     `json:"manager"`
	Address   Address   `json:"address"`
	UpdatedAt time.Time `json:"updated_at"`
}

type StoreDB struct {
	ID        int       `db:"id"`
	ManagerID int       `db:"manager_id"`
	AddressID int       `db:"address_id"`
	UpdatedAt time.Time `db:"updated_at"`
}

type StoreDBSecondVer struct {
	ID                int       `db:"id"`
	UpdatedAt         time.Time `db:"updated_at"`
	AddressID         int       `db:"address_id"`
	AddressAddress    string    `db:"address_address"`
	AddressAddress2   *string   `db:"address_address2"`
	AddressDistrict   string    `db:"address_district"`
	AddressCityID     *int      `db:"address_city_id"`
	AddressPostalCode *int      `db:"address_postal_code"`
	AddressPhone      string    `db:"address_phone"`
	AddressUpdatedAt  time.Time `db:"address_updated_at"`
	StaffID           int       `db:"staff_id"`
	StaffFirstName    string    `db:"staff_first_name"`
	StaffLastName     string    `db:"staff_last_name"`
	StaffEmail        *string   `db:"staff_email"`
	StaffActive       bool      `db:"staff_active"`
	StaffUsername     string    `db:"staff_username"`
	StaffPassword     *string   `db:"staff_password"`
	StaffPicture      *[]byte   `db:"staff_picture"`
	StaffUpdatedAt    time.Time `db:"staff_updated_at"`
}

type StoreInput struct {
	ManagerID int `json:"manager_id" validate:"required"`
	AddressID int `json:"address_id" validate:"required"`
}

func StoreDBtoStore(storeDB *StoreDB, addr *Address, stf Staff) Store {
	var address Address
	if addr != nil {
		address = *addr
	}

	return Store{
		ID:        storeDB.ID,
		Address:   address,
		Manager:   stf,
		UpdatedAt: storeDB.UpdatedAt,
	}
}

func StoreDBSecondVerToStore(storeSV []*StoreDBSecondVer) []Store {
	var resStore []Store

	for _, st := range storeSV {
		resStore = append(resStore, Store{
			ID: st.ID,
			Manager: Staff{
				ID:        st.StaffID,
				FirstName: st.StaffFirstName,
				LastName:  st.StaffLastName,
				Email:     st.StaffEmail,
				StoreID:   nil,
				Active:    st.StaffActive,
				Username:  st.StaffUsername,
				Password:  st.StaffPassword,
				Picture:   st.StaffPicture,
				Address:   nil,
				UpdatedAt: st.StaffUpdatedAt,
			},
			Address: Address{
				ID:         st.AddressID,
				Address:    st.AddressAddress,
				Address2:   st.AddressAddress2,
				District:   st.AddressDistrict,
				City:       nil,
				Country:    nil,
				PostalCode: st.AddressPostalCode,
				Phone:      st.AddressPhone,
				UpdatedAt:  st.AddressUpdatedAt,
			},
			UpdatedAt: st.UpdatedAt,
		},
		)
	}

	return resStore
}
