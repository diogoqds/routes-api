package entities

import (
	"time"
)

type Client struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Geolocation Point      `json:"geolocation"`
	RouteId     int        `json:"route_id"`
	Route       *Route     `json:"route"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}
