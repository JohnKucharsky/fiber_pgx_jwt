-- +goose Up
create table actor(
      id serial primary key,
      first_name varchar(45) not null,
      last_name varchar(45) not null,
      updated_at timestamptz not null default now(),
      unique(first_name,last_name)
);

-- +goose Down
drop table actor;