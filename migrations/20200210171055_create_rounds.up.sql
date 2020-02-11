CREATE TABLE rounds (
    id bigserial not null primary key,
    userid bigint not null,
    difficulty smallint not null,
    gamenumber varchar not null,
    gametime varchar not null,
	inpt      varchar not null,
	outpt     varchar not null,
    FOREIGN KEY (userid) REFERENCES users (id) ON DELETE CASCADE
);