package store

import (
	"context"
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
	"github.com/induzo/gocom/database/pginit/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AddressStore struct {
	db *pgxpool.Pool
}

func NewAddressStore(db *pgxpool.Pool) *AddressStore {
	return &AddressStore{
		db: db,
	}
}

func (as *AddressStore) Create(m domain.AddressInput) (
	int,
	error,
) {
	ctx := context.Background()

	rows, err := as.db.Query(
		ctx, `
        INSERT INTO address (address, address2, district, city_id, postal_code, phone)
        VALUES (@address, @address2, @district, @city_id, @postal_code, @phone)
        RETURNING id`,
		pgx.NamedArgs{
			"address":     m.Address,
			"address2":    m.Address2,
			"district":    m.District,
			"city_id":     m.CityID,
			"postal_code": m.PostalCode,
			"phone":       m.Phone,
		},
	)
	if err != nil {
		return 0, err
	}

	type idRes struct {
		ID int `db:"id"`
	}

	res, err := pgx.CollectExactlyOneRow(
		rows,
		pgx.RowToStructByName[idRes],
	)
	if err != nil {
		return 0, err
	}

	return res.ID, nil
}

func (as *AddressStore) GetMany() ([]*domain.Address, error) {
	ctx := context.Background()

	rows, err := as.db.Query(
		ctx, `select json_build_object(
       'id', address.id,
       'address', address.address,
       'address2', address.address2,
       'district', address.district,
       'city', CASE 
            WHEN address.city_id IS NULL THEN NULL
            ELSE json_build_object(
           'id', city.id,
           'city', city.city,
           'updated_at', city.updated_at
       ) 
       END,
       'country', CASE 
            WHEN city.country_id IS NULL THEN NULL
            ELSE json_build_object(
           'id', country.id,
           'country', country.country,
           'updated_at', country.updated_at
       ) 
       END,
       'updated_at', address.updated_at
     ) from address left join city on address.city_id = city.id
     left join country on city.country_id = country.id
     `,
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectRows(
		rows, pginit.JSONRowToAddrOfStruct[domain.Address],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (as *AddressStore) GetOne(id int) (*domain.Address, error) {
	ctx := context.Background()

	rows, err := as.db.Query(
		ctx,
		`select * from actor where id = @id`,
		pgx.NamedArgs{"id": id},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.Address],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (as *AddressStore) Update(u domain.AddressInput, id int) (
	int,
	error,
) {
	ctx := context.Background()

	rows, err := as.db.Query(
		ctx,
		`UPDATE actor SET 
                first_name = @first_name,
    			last_name = @last_name 
             WHERE id = @id 
        returning id`,
		pgx.NamedArgs{
			"id":         id,
			"first_name": u.Address2,
			"last_name":  u.District,
		},
	)
	if err != nil {
		return 0, err
	}

	type idRes struct {
		ID int `db:"id"`
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[idRes],
	)
	if err != nil {
		return 0, err
	}

	return res.ID, nil
}

func (as *AddressStore) Delete(id int) (int, error) {
	ctx := context.Background()

	rows, err := as.db.Query(
		ctx,
		`delete from actor where id = @id 
        returning id, first_name, last_name, updated_at`,
		pgx.NamedArgs{
			"id": id,
		},
	)
	if err != nil {
		return 0, err
	}

	type idRes struct {
		ID int `db:"id"`
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[idRes],
	)
	if err != nil {
		return 0, err
	}

	return res.ID, nil
}
