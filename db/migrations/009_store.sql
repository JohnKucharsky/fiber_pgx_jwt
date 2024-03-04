-- +goose Up
create table store(
      id serial primary key,
      manager_id smallint not null references staff(id) on update cascade on delete restrict,
      address_id smallint references address(id) on update cascade on delete restrict,
      updated_at timestamptz not null default now()
);

-- +goose Down
drop table store;