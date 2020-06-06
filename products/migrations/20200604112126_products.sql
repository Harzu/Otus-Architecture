-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table if not exists products (
    id serial primary key,
    name text not null,
    type text not null,
    price integer not null
);

insert into products (name, type, price) values ('iphone 6', 'mobile', 100000);
insert into products (name, type, price) values ('nexus 5', 'mobile', 80000);
insert into products (name, type, price) values ('htc one', 'mobile', 1100000);
insert into products (name, type, price) values ('google phone', 'mobile', 210000);
insert into products (name, type, price) values ('iphone 10', 'mobile', 90000);
insert into products (name, type, price) values ('samsung galaxy 20', 'mobile', 160000);
insert into products (name, type, price) values ('sammsung galaxe a12', 'mobile', 150000);
insert into products (name, type, price) values ('sony Z3', 'mobile', 140000);
insert into products (name, type, price) values ('sony Z4', 'mobile', 130000);
insert into products (name, type, price) values ('motorola X', 'mobile', 120000);
insert into products (name, type, price) values ('acer swift 3', 'laptop', 100000);
insert into products (name, type, price) values ('microsoft surface', 'laptop', 80000);
insert into products (name, type, price) values ('apple macbook pro 15', 'laptop', 1100000);
insert into products (name, type, price) values ('apple macbook pro 16', 'laptop', 210000);
insert into products (name, type, price) values ('apple macbook pro 13', 'laptop', 90000);
insert into products (name, type, price) values ('chrome book', 'laptop', 160000);
insert into products (name, type, price) values ('asus rog zepherus g14', 'laptop', 150000);
insert into products (name, type, price) values ('msi gaming', 'laptop', 140000);
insert into products (name, type, price) values ('dell xps 15', 'laptop', 130000);
insert into products (name, type, price) values ('dell xps 13', 'laptop', 120000);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table products