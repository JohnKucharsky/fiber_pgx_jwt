package store

import (
	"context"
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CustomerStore struct {
	db *pgxpool.Pool
}

func NewCustomerStore(db *pgxpool.Pool) *CustomerStore {
	return &CustomerStore{
		db: db,
	}
}

func (s *CustomerStore) Create(m domain.CustomerInput) (
	*domain.CustomerDB,
	error,
) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx, `
        INSERT INTO customer (store_id, 
                              first_name, 
                              last_name, 
                              email, 
                              address_id, 
                              active, 
                              create_date)
        VALUES (@store_id,
                @first_name,
                @last_name,
                @email,
                @address_id,
                @active,
                @create_date)
        RETURNING id, store_id, first_name, last_name, email, address_id, 
        active, create_date, updated_at`,
		pgx.NamedArgs{
			"store_id":    m.StoreID,
			"first_name":  m.FirstName,
			"last_name":   m.LastName,
			"email":       m.Email,
			"address_id":  m.AddressID,
			"active":      m.Active,
			"create_date": m.CreateDate,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows,
		pgx.RowToStructByName[domain.CustomerDB],
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *CustomerStore) GetMany() ([]*domain.CustomerDB, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx, `select * from customer`,
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectRows(
		rows, pgx.RowToAddrOfStructByName[domain.CustomerDB],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CustomerStore) GetOne(id int) (*domain.CustomerDB, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx,
		`select * from customer where id = @id`,
		pgx.NamedArgs{"id": id},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.CustomerDB],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CustomerStore) Update(m domain.CustomerInput, id int) (
	*domain.CustomerDB,
	error,
) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx, `
        update customer set 
				store_id = @store_id, 
				  first_name = @first_name, 
				  last_name = @last_name, 
				  email = @email, 
				  address_id = @address_id, 
				  active = @active, 
				  create_date = @create_date
         	where customer.id = @id
        RETURNING id, store_id, first_name, last_name, email, address_id, 
        active, create_date, updated_at`,
		pgx.NamedArgs{
			"store_id":    m.StoreID,
			"first_name":  m.FirstName,
			"last_name":   m.LastName,
			"email":       m.Email,
			"address_id":  m.AddressID,
			"active":      m.Active,
			"create_date": m.CreateDate,
			"id":          id,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.CustomerDB],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *CustomerStore) Delete(id int) (*domain.CustomerDB, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx,
		`delete from customer where id = @id 
        returning id, store_id, first_name, last_name, email, address_id, 
        active, create_date, updated_at`,
		pgx.NamedArgs{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.CustomerDB],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
