package model

import "time"

type SubmissionDetail struct {
	ID            string    `json:"id"`
	SubmissionId  string    `json:"submission_id"`
	ItemId        string    `json:"item_id"`
	AmountSubmit  int       `json:"amount_submit"`
	StatusDetail  string    `json:"status_detail"`
	AmountApprove int       `json:"amount_approve"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (SubmissionDetail) TableName() string {
	return "submission_detail"
}
