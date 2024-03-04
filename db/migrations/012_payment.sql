-- +goose Up
create table payment(
       id serial primary key,
       customer_id smallint references customer(id) on update cascade on delete restrict,
       staff_id smallint references staff(id) on update cascade on delete restrict,
       rental_id smallint not null references rental(id) on update cascade on delete restrict,
       amount numeric(5,2) not null,
       payment_date timestamptz not null
);

-- +goose Down
drop table payment;