-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id integer primary key,
    city varchar(50),
    created_at timestamp default now()
);

