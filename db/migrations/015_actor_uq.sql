-- +goose Up
alter table actor add constraint unique_first_name_last_name
    unique (first_name,last_name);

-- +goose Down
alter table actor drop constraint unique_first_name_last_name;