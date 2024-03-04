-- +goose Up
alter table city add constraint unique_city_country_id
    unique (city,country_id);

-- +goose Down
alter table city drop constraint unique_city_country_id;