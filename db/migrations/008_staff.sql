-- +goose Up
create table staff(
     id serial primary key,
     first_name varchar(45) not null,
     last_name varchar(45) not null,
     address_id smallint references address(id) on update cascade on delete set null,
     email varchar(50) unique,
     store_id smallint not null,
     active bool not null default true,
     username varchar(25) not null,
     password varchar(40),
     picture bytea,
     updated_at timestamptz not null default now(),
     unique (first_name, last_name)
);

-- +goose Down
drop table staff;