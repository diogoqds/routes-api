package entities

type Point struct {
	Type        string    `json:"type" db:"type"`
	Coordinates []float64 `json:"coordinates" db:"coordinates"`
}
