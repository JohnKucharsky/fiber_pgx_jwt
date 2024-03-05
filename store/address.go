package store

import (
	"context"
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
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
		ctx, `
		select 
			address.id,
			address.address,
			address.address2,
			address.district,
			city.id as city_id,
			city.city as city_city,
			city.updated_at as city_updated_at,
			country.id as country_id,
			country.country as country_country,
			country.updated_at as country_updated_at,
			address.postal_code,
			address.phone,
			address.updated_at from address 
		left join city on address.city_id = city.id
		left join country on city.country_id = country.id;
     `,
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectRows(
		rows, domain.ScanAddress,
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
		`select 
			address.id,
			address.address,
			address.address2,
			address.district,
			city.id as city_id,
			city.city as city_city,
			city.updated_at as city_updated_at,
			country.id as country_id,
			country.country as country_country,
			country.updated_at as country_updated_at,
			address.postal_code,
			address.phone,
			address.updated_at from address 
		left join city on address.city_id = city.id
		left join country on city.country_id = country.id 
		where address.id = @id`,
		pgx.NamedArgs{"id": id},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, domain.ScanAddress,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (as *AddressStore) Update(m domain.AddressInput, id int) error {
	ctx := context.Background()

	_, err := as.db.Exec(
		ctx,
		`UPDATE address SET 
			address = @address,
			address2 = @address2,
			district = @district,
			city_id = @city_id,
			postal_code = @postal_code,
			phone = @phone
             WHERE id = @id`,
		pgx.NamedArgs{
			"id":          id,
			"address":     m.Address,
			"address2":    m.Address2,
			"district":    m.District,
			"city_id":     m.CityID,
			"postal_code": m.PostalCode,
			"phone":       m.Phone,
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func (as *AddressStore) Delete(id int) (int, error) {
	ctx := context.Background()

	rows, err := as.db.Query(
		ctx,
		`delete from address where id = @id 
        returning id`,
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
