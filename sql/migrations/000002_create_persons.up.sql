CREATE TABLE persons (
    id uuid primary key, 
    name varchar(120),
    age integer,
    cpf character varying(11),
    address text,
    user_id uuid,
    foreign key(user_id) references users(id) 
);