package store

import (
	"context"
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CountryStore struct {
	db *pgxpool.Pool
}

func NewCountryStore(db *pgxpool.Pool) *CountryStore {
	return &CountryStore{
		db: db,
	}
}

func (s *CountryStore) Create(m domain.CountryInput) (
	*domain.Country,
	error,
) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx, `
        INSERT INTO country (country)
        VALUES (@country)
        RETURNING id, country, updated_at`,
		pgx.NamedArgs{
			"country": m.Country,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows,
		pgx.RowToStructByName[domain.Country],
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *CountryStore) GetMany() ([]*domain.Country, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx, `select * from country`,
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectRows(
		rows, pgx.RowToAddrOfStructByName[domain.Country],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CountryStore) GetOne(id int) (*domain.Country, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx,
		`select * from country where id = @id`,
		pgx.NamedArgs{"id": id},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.Country],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CountryStore) Update(m domain.CountryInput, id int) (
	*domain.Country,
	error,
) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx,
		`UPDATE country SET 
                country = @country
             WHERE id = @id 
        returning id, country, updated_at`,
		pgx.NamedArgs{
			"id":      id,
			"country": m.Country,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.Country],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CountryStore) Delete(id int) (*domain.Country, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx,
		`delete from country where id = @id 
        returning id, country, updated_at`,
		pgx.NamedArgs{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.Country],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
