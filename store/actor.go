package store

import (
	"context"
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ActorStore struct {
	db *pgxpool.Pool
}

func NewActorStore(db *pgxpool.Pool) *ActorStore {
	return &ActorStore{
		db: db,
	}
}

func (as *ActorStore) Create(m domain.ActorInput) (
	*domain.Actor,
	error,
) {
	ctx := context.Background()

	rows, err := as.db.Query(
		ctx, `
        INSERT INTO actor (first_name, last_name)
        VALUES (@first_name, @last_name)
        RETURNING id, first_name, last_name, updated_at`,
		pgx.NamedArgs{
			"first_name": m.FirstName,
			"last_name":  m.LastName,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows,
		pgx.RowToStructByName[domain.Actor],
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (as *ActorStore) GetMany() ([]*domain.Actor, error) {
	ctx := context.Background()

	rows, err := as.db.Query(
		ctx, `select * from actor`,
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectRows(
		rows, pgx.RowToAddrOfStructByName[domain.Actor],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (as *ActorStore) GetOne(id int) (*domain.Actor, error) {
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
		rows, pgx.RowToAddrOfStructByName[domain.Actor],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (as *ActorStore) Update(u domain.ActorInput, id int) (
	*domain.Actor,
	error,
) {
	ctx := context.Background()

	rows, err := as.db.Query(
		ctx,
		`UPDATE actor SET 
                first_name = @first_name,
    			last_name = @last_name 
             WHERE id = @id 
        returning id, first_name, last_name, updated_at`,
		pgx.NamedArgs{
			"id":         id,
			"first_name": u.FirstName,
			"last_name":  u.LastName,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.Actor],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (as *ActorStore) Delete(id int) (*domain.Actor, error) {
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
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.Actor],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
