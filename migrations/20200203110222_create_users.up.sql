CREATE TABLE users (
    id bigserial not null primary key,
    nickname varchar not null unique,
    email varchar not null unique,
    encrypted_password varchar not null
);