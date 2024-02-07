package entity

import (
	"database/sql"
	"easyshop-api/pkg/entity"
)

type Person struct {
	ID      entity.ID
	Name    sql.NullString
	Age     sql.NullInt32
	Cpf     sql.NullString
	Address sql.NullString
	UserID  entity.ID
}
