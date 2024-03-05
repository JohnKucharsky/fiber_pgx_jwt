package store

import (
	"context"
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LanguageStore struct {
	db *pgxpool.Pool
}

func NewLanguageStore(db *pgxpool.Pool) *LanguageStore {
	return &LanguageStore{
		db: db,
	}
}

func (s *LanguageStore) Create(m domain.LanguageInput) (
	*domain.Language,
	error,
) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx, `
        INSERT INTO language (name)
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
		pgx.RowToStructByName[domain.Language],
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *LanguageStore) GetMany() ([]*domain.Language, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx, `select * from language`,
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectRows(
		rows, pgx.RowToAddrOfStructByName[domain.Language],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *LanguageStore) GetOne(id int) (*domain.Language, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx,
		`select * from language where id = @id`,
		pgx.NamedArgs{"id": id},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.Language],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *LanguageStore) Update(m domain.LanguageInput, id int) (
	*domain.Language,
	error,
) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx,
		`UPDATE language SET 
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
		rows, pgx.RowToAddrOfStructByName[domain.Language],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *LanguageStore) Delete(id int) (*domain.Language, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx,
		`delete from language where id = @id 
        returning id, name, updated_at`,
		pgx.NamedArgs{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.Language],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
