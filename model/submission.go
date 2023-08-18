package model

import "time"

type Submission struct {
	Id         string    `json:"id"`
	EmployeeId string    `json:"employee_id"`
	PeriodId   string    `json:"period_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Employee   Employee  `gorm:"foreignKey:EmployeeId"`
	Period     Period    `gorm:"foreignKey:PeriodId"`
}

func (Submission) TableName() string {
	return "submission"
}
