package dto

import "time"

type AllParkingZonesResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    []SingleParkingZone `json:"data"`
}

type SingleParkingZone struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	TotalCapacity int       `json:"total_capacity"`
	PricePerHour  float64   `json:"price_per_hour"`
	CreatedAt     time.Time `json:"created_at"`
}

type ParkingZoneByIDResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    SingleParkingZone `json:"data"`
}

type CreateParkingZoneResponse struct {
	Success bool `json:"success"`

	Message string      `json:"message"`
	Data    ParkingZone `json:"data"`
}
type ParkingZone struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	TotalCapacity int       `json:"total_capacity"`
	PricePerHour  float64   `json:"price_per_hour"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
