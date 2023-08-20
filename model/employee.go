package model

import "time"

type Employee struct {
	Id           string `json:"id"`
	EmployeeName string `json:"employee_name"`
	PhoneNumber  string `json:"phone_number"`
	IsDelete     bool
	DepartmentId string     `json:"department_id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Department   Department `gorm:"foreignKey:DepartmentId"`
}

func (Employee) TableName() string {
	return "employee"
}
