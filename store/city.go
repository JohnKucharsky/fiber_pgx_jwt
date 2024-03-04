package store

import (
	"context"
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
	"github.com/induzo/gocom/database/pginit/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CityStore struct {
	db *pgxpool.Pool
}

func NewCityStore(db *pgxpool.Pool) *CityStore {
	return &CityStore{
		db: db,
	}
}

func (s *CityStore) Create(m domain.CityInput) (
	int,
	error,
) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx, `
        INSERT INTO city (city, country_id)
        VALUES (@city, @country_id)
        RETURNING id`,
		pgx.NamedArgs{
			"city":       m.City,
			"country_id": m.CountryID,
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

func (s *CityStore) GetMany() ([]*domain.City, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx,
		`select json_build_object(
       'id', city.id,
       'city', city.city,
       'country', CASE 
            WHEN city.country_id IS NULL THEN NULL
            ELSE json_build_object(
           'id', country.id,
           'country', country.country,
           'updated_at', country.updated_at
       ) 
       END,
       'updated_at', city.updated_at
     ) from city left join country on city.country_id = country.id`,
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectRows(
		rows, pginit.JSONRowToAddrOfStruct[domain.City],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CityStore) GetOne(id int) (*domain.City, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx,
		`select json_build_object(
       'id',city.id,
       'city',city.city,
       'country', CASE 
            WHEN city.country_id IS NULL THEN NULL
            ELSE json_build_object(
           'id', country.id,
           'country', country.country,
           'updated_at', country.updated_at
       ) 
       END,
       'updated_at',city.updated_at
     ) from city left join country on city.country_id = country.id where city.id = @id`,
		pgx.NamedArgs{"id": id},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pginit.JSONRowToAddrOfStruct[domain.City],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CityStore) Update(m domain.CityInput, id int) (
	int,
	error,
) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx,
		`UPDATE city SET 
                city = @city,
                country_id = @country_id
             WHERE id = @id 
        returning id`,
		pgx.NamedArgs{
			"id":         id,
			"country_id": m.CountryID,
			"city":       m.City,
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

func (s *CityStore) Delete(id int) (int, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx,
		`delete from city where id = @id 
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
