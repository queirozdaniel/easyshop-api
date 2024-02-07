package entity

import (
	"database/sql"
	"easyshop-api/pkg/entity"
)

type Resource struct {
	ID          entity.ID
	Name        string
	Description sql.NullString
}

type Role struct {
	ID          entity.ID
	Name        string
	Description sql.NullString
}

type RolesResource struct {
	ID         entity.ID
	RoleID     entity.ID
	ResourceID entity.ID
}
