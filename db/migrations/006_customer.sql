-- +goose Up
create table customer(
     id serial primary key,
     store_id smallint not null,
     first_name varchar(45) not null,
     last_name varchar(45) not null,
     email varchar(50) unique,
     address_id smallint references address(id) on update cascade on delete set null,
     active bool not null default true,
     create_date date not null default now()::text::date,
     updated_at timestamptz not null default now(),
     unique (first_name, last_name)
);

-- +goose Down
drop table customer;