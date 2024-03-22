-- +goose Up
alter table address alter
column postal_code type integer;
-- +goose Down
alter table address alter
    column postal_code type smallint;