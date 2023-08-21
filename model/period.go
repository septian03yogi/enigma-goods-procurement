package model

import "time"

type Period struct {
	ID         string    `json:"id"`
	PeriodName string    `json:"period_name"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	IsActive   bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Period) TableName() string {
	return "period"
}
