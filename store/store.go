package store

import (
	"context"
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StoreStore struct {
	db *pgxpool.Pool
}

func NewStoreStore(db *pgxpool.Pool) *StoreStore {
	return &StoreStore{
		db: db,
	}
}

func (as *StoreStore) Create(m domain.StoreInput) (
	*domain.StoreDB,
	error,
) {
	ctx := context.Background()

	rows, err := as.db.Query(
		ctx, `
        INSERT INTO store (manager_id, address_id)
        VALUES (@manager_id, @address_id)
        RETURNING id, manager_id, address_id, updated_at`,
		pgx.NamedArgs{
			"manager_id": m.ManagerID,
			"address_id": m.AddressID,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows,
		pgx.RowToStructByName[domain.StoreDB],
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (as *StoreStore) GetMany() ([]*domain.StoreDBSecondVer, error) {
	ctx := context.Background()

	rows, err := as.db.Query(
		ctx, `select 
			store.id,
			store.updated_at,
			address.id as address_id,
			address.address as address_address,
			address.address2 as address_address2,
			address.district as address_district,
			address.city_id  as address_city_id,
			address.postal_code as address_postal_code,
			address.phone as address_phone,
			address.updated_at as address_updated_at,
			staff.id as staff_id,
			staff.first_name as staff_first_name,
			staff.last_name as staff_last_name,
			staff.email as staff_email,
			staff.active as staff_active,
			staff.username as staff_username,
			staff.password as staff_password,
			staff.picture as staff_picture,
			staff.updated_at as staff_updated_at
			from store left join staff on store.manager_id = staff.id
           left join address on store.address_id = address.id`,
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectRows(
		rows, func(row pgx.CollectableRow) (*domain.StoreDBSecondVer, error) {
			var storeDB domain.StoreDBSecondVer
			err := row.Scan(
				&storeDB.ID,
				&storeDB.UpdatedAt,
				&storeDB.AddressID,
				&storeDB.AddressAddress,
				&storeDB.AddressAddress2,
				&storeDB.AddressDistrict,
				&storeDB.AddressCityID,
				&storeDB.AddressPostalCode,
				&storeDB.AddressPhone,
				&storeDB.AddressUpdatedAt,
				&storeDB.StaffID,
				&storeDB.StaffFirstName,
				&storeDB.StaffLastName,
				&storeDB.StaffEmail,
				&storeDB.StaffActive,
				&storeDB.StaffUsername,
				&storeDB.StaffPassword,
				&storeDB.StaffPicture,
				&storeDB.StaffUpdatedAt,
			)
			return &storeDB, err
		},
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (as *StoreStore) GetOne(id int) (*domain.StoreDB, error) {
	ctx := context.Background()

	rows, err := as.db.Query(
		ctx,
		`select * from store where id = @id`,
		pgx.NamedArgs{"id": id},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.StoreDB],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (as *StoreStore) Update(u domain.StoreInput, id int) (
	*domain.StoreDB,
	error,
) {
	ctx := context.Background()

	rows, err := as.db.Query(
		ctx,
		`UPDATE store SET 
                address_id = @address_id,
    			manager_id = @manager_id 
             WHERE id = @id 
        returning id, address_id, manager_id, updated_at`,
		pgx.NamedArgs{
			"id":         id,
			"address_id": u.AddressID,
			"manager_id": u.ManagerID,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.StoreDB],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (as *StoreStore) Delete(id int) (*domain.StoreDB, error) {
	ctx := context.Background()

	rows, err := as.db.Query(
		ctx,
		`delete from store where id = @id 
        returning id, address_id, manager_id, updated_at`,
		pgx.NamedArgs{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.StoreDB],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
