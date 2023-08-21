package model

import "time"

type Department struct {
	ID             string    `json:"id"`
	DepartmentName string    `json:"department_name"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (Department) TableName() string {
	return "department"
}
