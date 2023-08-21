package model

import "time"

type UserCredential struct {
	ID         string    `json:"id"`
	Password   string    `json:"password"`
	RoleId     string    `json:"role_id"`
	EmployeeId string    `json:"employee_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (UserCredential) TableName() string {
	return "user_credential"
}
