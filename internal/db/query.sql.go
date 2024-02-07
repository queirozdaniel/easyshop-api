// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const getUserPerson = `-- name: GetUserPerson :one
select u.id, email, password, p.id, name, age, cpf, address, user_id from users u
join persons p on p.user_id = u.id
where u.id = $1
`

type GetUserPersonRow struct {
	ID       uuid.UUID
	Email    sql.NullString
	Password sql.NullString
	ID_2     uuid.UUID
	Name     sql.NullString
	Age      sql.NullInt32
	Cpf      sql.NullString
	Address  sql.NullString
	UserID   uuid.NullUUID
}

func (q *Queries) GetUserPerson(ctx context.Context, id uuid.UUID) (GetUserPersonRow, error) {
	row := q.db.QueryRowContext(ctx, getUserPerson, id)
	var i GetUserPersonRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.ID_2,
		&i.Name,
		&i.Age,
		&i.Cpf,
		&i.Address,
		&i.UserID,
	)
	return i, err
}

const listPersons = `-- name: ListPersons :many
select id, name, age, cpf, address, user_id from persons
`

func (q *Queries) ListPersons(ctx context.Context) ([]Person, error) {
	rows, err := q.db.QueryContext(ctx, listPersons)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Person
	for rows.Next() {
		var i Person
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Age,
			&i.Cpf,
			&i.Address,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
select id, email, password from users
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Email, &i.Password); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}