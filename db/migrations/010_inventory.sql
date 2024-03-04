-- +goose Up
create table inventory(
      id serial primary key,
      film_id smallint not null references film(id) on update cascade on delete restrict,
      store_id smallint not null references store(id) on update cascade on delete restrict,
      updated_at timestamptz not null default now(),
      unique(film_id, store_id)
);

-- +goose Down
drop table inventory;