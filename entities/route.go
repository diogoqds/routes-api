package entities

import (
	"time"
)

type Route struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	SellerId  int        `json:"seller_id"`
	Seller    *Seller    `json:"seller"`
	Bounds    *Polygon   `json:"bounds"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}
