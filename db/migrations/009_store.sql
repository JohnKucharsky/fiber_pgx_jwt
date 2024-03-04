-- +goose Up
create table store(
      id serial primary key,
      manager_id smallint not null references staff(id) on update cascade on delete restrict,
      address_id smallint not null references address(id) on update cascade on delete restrict,
      updated_at timestamptz not null default now(),
      unique (manager_id, address_id)
);

-- +goose Down
drop table store;