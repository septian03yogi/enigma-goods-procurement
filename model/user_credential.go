package model

import "time"

type UserCredential struct {
	Id         string    `json:"id"`
	Password   string    `json:"password"`
	RoleId     string    `json:"role_id"`
	EmployeeId string    `json:"employee_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	RoleUser   RoleUser  `gorm:"foreignKey:RoleId"`
	Employee   Employee  `gorm:"foreignKey:EmployeeId"`
}

func (UserCredential) TableName() string {
	return "user_credential"
}
