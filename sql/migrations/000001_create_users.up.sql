CREATE TABLE users (
    id uuid primary key, 
    email varchar(120) unique,
    password varchar(32)
);