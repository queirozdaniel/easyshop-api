CREATE TABLE users_resources (
    id uuid primary key,
    user_id uuid,
    role_resource_id uuid,
    foreign key(user_id) references users(id),
    foreign key(role_resource_id) references roles_resources(id)
);
