-- +goose Up
-- SQL in this section is executed when the migration is applied.
create table users (
    email text primary key,
    password_hash text not null,
    first_name text default '',
    last_name text default ''
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table users;