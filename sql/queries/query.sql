-- name: ListUsers :many
select * from users;

-- name: ListPersons :many
select * from persons;

-- name: GetUserPerson :one
select * from users u
join persons p on p.user_id = u.id
where u.id = $1;