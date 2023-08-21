package model

import (
	"time"
)

type Employee struct {
	ID           string `json:"id"`
	EmployeeName string `json:"employee_name"`
	PhoneNumber  string `json:"phone_number"`
	IsDelete     bool
	DepartmentId string    `json:"department_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (Employee) TableName() string {
	return "employee"
}
