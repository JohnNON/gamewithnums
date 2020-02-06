CREATE TABLE records (
    id bigserial not null primary key,
    userid bigint not null,
    difficulty smallint not null,
    roundcount smallint not null,
    gametime integer not null,
    FOREIGN KEY (userid) REFERENCES users (id) ON DELETE CASCADE
);