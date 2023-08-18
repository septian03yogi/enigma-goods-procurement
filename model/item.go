package model

import "time"

type Item struct {
	Id        string    `json:"id"`
	ItemName  string    `json:"item_mame"`
	Stock     int       `json:"stock"`
	UomId     string    `json:"uom_id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Uom       Uom       `gorm:"foreignKey:UomId"`
}

func (Item) TableName() string {
	return "item"
}
