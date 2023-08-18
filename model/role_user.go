package model

import "time"

type RoleUser struct {
	Id        string    `json:"id"`
	RoleName  string    `json:"role_name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (RoleUser) TableName() string {
	return "role_user"
}
