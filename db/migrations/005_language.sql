-- +goose Up
create table language(
     id serial primary key,
     name varchar(45) not null,
     updated_at timestamptz not null default now()
);

-- +goose Down
drop table language;