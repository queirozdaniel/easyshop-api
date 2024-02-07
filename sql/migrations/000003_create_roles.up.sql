CREATE TABLE roles (
    id uuid primary key, 
    name varchar(40),
    description varchar(40)
);

CREATE TABLE resources (
    id uuid primary key, 
    name varchar(40),
    description varchar(40)
);

CREATE TABLE roles_resources (
    id uuid primary key, 
    role_id uuid,
    resource_id uuid,
    foreign key(role_id) references roles(id),
    foreign key(resource_id) references resources(id)
);
