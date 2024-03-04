-- +goose Up
create table rental(
      id serial primary key,
      rental_date timestamptz not null,
      inventory_id smallint not null references inventory(id) on update cascade on delete restrict,
      customer_id smallint references customer(id) on update cascade on delete restrict,
      staff_id smallint references staff(id) on update cascade on delete restrict,
      return_date timestamptz,
      updated_at timestamptz not null default now()
);

-- +goose Down
drop table rental;