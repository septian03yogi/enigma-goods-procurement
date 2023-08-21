package model

import "time"

type Uom struct {
	ID        string    `json:"id"`
	UomName   string    `json:"uom_name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (Uom) TableName() string {
	return "uom"
}
