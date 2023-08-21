package model

import (
	"time"
)

type Item struct {
	ID        string    `json:"id"`
	ItemName  string    `json:"item_name"`
	Stock     int       `json:"stock"`
	UomID     string    `json:"uom_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time
}

func (Item) TableName() string {
	return "item"
}
