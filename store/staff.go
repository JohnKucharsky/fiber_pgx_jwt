package store

import (
	"context"
	"github.com/JohnKucharsky/fiber_pgx_jwt/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StaffStore struct {
	db *pgxpool.Pool
}

func NewStaffStore(db *pgxpool.Pool) *StaffStore {
	return &StaffStore{
		db: db,
	}
}

func (s *StaffStore) Create(m domain.StaffInput) (
	*domain.StaffDB,
	error,
) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx, `
        INSERT INTO staff (
                            first_name,
							last_name,
							email,
							store_id,
							active,
							username,
							password,
							picture,
							address_id
                              )
        VALUES (@first_name,
                @last_name,
                @email,
                @store_id,
                @active,
                @username,
                @password,
                @picture,
                @address_id
                )
        RETURNING id,
            first_name,
			last_name,
			email,
			store_id,
			active,
			username,
			password,
			picture,
			address_id,
             updated_at`,
		pgx.NamedArgs{
			"first_name": m.FirstName,
			"last_name":  m.LastName,
			"email":      m.Email,
			"store_id":   m.StoreID,
			"active":     m.Active,
			"username":   m.Username,
			"password":   m.Password,
			"picture":    m.Picture,
			"address_id": m.AddressID,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows,
		pgx.RowToStructByName[domain.StaffDB],
	)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (s *StaffStore) GetMany() ([]*domain.StaffDB, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx, `select * from staff`,
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectRows(
		rows, pgx.RowToAddrOfStructByName[domain.StaffDB],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *StaffStore) GetOne(id int) (*domain.StaffDB, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx,
		`select * from staff where id = @id`,
		pgx.NamedArgs{"id": id},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.StaffDB],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *StaffStore) Update(m domain.StaffInput, id int) (
	*domain.StaffDB,
	error,
) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx, `
        update staff set 
				  first_name = @first_name, 
				  last_name = @last_name, 
				  email = @email, 
				  store_id = @store_id, 
				  active = @active, 
				  username = @username, 
				  password = @password,
				  picture = @picture,
				  address_id = @address_id
         	where staff.id = @id
        RETURNING id,
            first_name,
			last_name,
			email,
			store_id,
			active,
			username,
			password,
			picture,
			address_id,
             updated_at`,
		pgx.NamedArgs{
			"first_name": m.FirstName,
			"last_name":  m.LastName,
			"email":      m.Email,
			"store_id":   m.StoreID,
			"active":     m.Active,
			"username":   m.Username,
			"password":   m.Password,
			"picture":    m.Picture,
			"address_id": m.AddressID,
			"id":         id,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.StaffDB],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *StaffStore) Delete(id int) (*domain.StaffDB, error) {
	ctx := context.Background()

	rows, err := s.db.Query(
		ctx,
		`delete from staff where id = @id 
        returning id,
            first_name,
			last_name,
			email,
			store_id,
			active,
			username,
			password,
			picture,
			address_id,
             updated_at`,
		pgx.NamedArgs{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	res, err := pgx.CollectExactlyOneRow(
		rows, pgx.RowToAddrOfStructByName[domain.StaffDB],
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}
