-- +goose Up
create table country(
    id serial primary key,
    country varchar(50) unique not null,
    updated_at timestamptz not null default now()
);

create table city(
    id serial primary key,
    city varchar(50) not null,
    country_id smallint references country(id) on update cascade on delete set null,
    updated_at timestamptz not null default now(),
    unique(city, country_id)
);

create table address(
      id serial primary key,
      address varchar(50) not null,
      address2 varchar(50),
      district varchar(30) not null,
      city_id smallint references city(id) on update cascade on delete set null,
      postal_code smallint unique,
      phone varchar(20) unique not null,
      updated_at timestamptz not null default now(),
      unique(address, city_id)
);

-- +goose Down
drop table address;
drop table city;
drop table country;
