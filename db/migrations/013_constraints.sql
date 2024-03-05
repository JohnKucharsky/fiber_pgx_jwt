-- +goose Up
alter table customer add constraint customer_store_fk
foreign key (store_id) references store(id)
    on update cascade on delete cascade;
alter table inventory add constraint inventory_store_fk
foreign key (store_id) references store(id)
    on update cascade on delete cascade ;

-- +goose Down
alter table customer drop constraint customer_store_fk;
alter table inventory drop constraint inventory_store_fk;