-- +goose Up
create table film_category(
     film_id smallint not null references film(id) on update cascade on delete restrict,
     category_id smallint not null references category(id) on update cascade on delete restrict,
     updated_at timestamptz not null default now()
);

-- +goose Down
drop table film_category;