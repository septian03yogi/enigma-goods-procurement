package model

import "time"

type SubmisisonStatus struct {
	ID           string    `json:"id"`
	StatusDetail string    `json:"status_detail"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (SubmisisonStatus) TableName() string {
	return "submission_status"
}
