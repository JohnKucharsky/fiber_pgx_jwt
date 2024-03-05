package store

import (
	"context"
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryStore struct {
	db *pgxpool.Pool
}

func NewCategoryStore(db *pgxpool.Pool) *CategoryStore {
	return &CategoryStore{
		db: db,
	}
}

func (s *CategoryStore) Create(m domain.CategoryInput) (
	*domain.Category,
	error,
) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx, `
        INSERT INTO category (name)
        VALUES (@name)
        RETURNING id, name, updated_at`,
		pgx.NamedArgs{
			"name": m.Name,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows,
		pgx.RowToStructByName[domain.Category],
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *CategoryStore) GetMany() ([]*domain.Category, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx, `select * from category`,
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectRows(
		rows, pgx.RowToAddrOfStructByName[domain.Category],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CategoryStore) GetOne(id int) (*domain.Category, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx,
		`select * from category where id = @id`,
		pgx.NamedArgs{"id": id},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.Category],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CategoryStore) Update(m domain.CategoryInput, id int) (
	*domain.Category,
	error,
) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx,
		`UPDATE category SET 
                name = @name
             WHERE id = @id 
        returning id, name, updated_at`,
		pgx.NamedArgs{
			"id":   id,
			"name": m.Name,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.Category],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CategoryStore) Delete(id int) (*domain.Category, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx,
		`delete from category where id = @id 
        returning id, name, updated_at`,
		pgx.NamedArgs{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.Category],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
