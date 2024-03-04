-- +goose Up
alter table customer add constraint customer_store_fk
foreign key (store_id) references store(id);
alter table inventory add constraint inventory_store_fk
foreign key (store_id) references store(id);

-- +goose Down
alter table customer drop constraint customer_store_fk;
alter table inventory drop constraint inventory_store_fk;