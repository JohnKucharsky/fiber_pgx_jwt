-- +goose Up
create type mpaa_rating as enum ('G', 'PG', 'PG-13','R', 'NC17');
create table film(
     id serial primary key,
     title varchar(225) unique not null,
     description text,
     release_year smallint,
     language_id smallint not null references language(id) on update cascade on delete restrict,
     rental_duration smallint not null default 3,
     rental_rate numeric(4,2) not null default 4.99,
     length smallint,
     replacement_cost numeric(5,2) not null default 19.99,
     rating mpaa_rating default 'G'::mpaa_rating,
     updated_at timestamptz not null default now()
);
create table film_actor(
    actor_id smallint not null references actor(id) on update cascade on delete restrict,
    film_id smallint not null references film(id) on update cascade on delete restrict,
    updated_at timestamptz not null default now(),
    primary key (actor_id,film_id)
);

-- +goose Down
drop table film_actor;
drop table film;
drop type mpaa_rating;
