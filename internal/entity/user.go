package entity

import (
	"easyshop-api/pkg/entity"
)

type User struct {
	ID       entity.ID
	Email    string
	Password string
}

type UsersResource struct {
	ID             entity.ID
	UserID         entity.ID
	RoleResourceID entity.ID
}
