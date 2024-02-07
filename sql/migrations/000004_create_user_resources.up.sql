CREATE TABLE users_resources (
    user_id uuid,
    role_resource_id uuid,
    foreign key(user_id) references users(id),
    foreign key(role_resource_id) references roles_resources(id)
);

ALTER TABLE users_resources ADD CONSTRAINT "ID_PKEY" KEY (user_id, role_resource_id);
