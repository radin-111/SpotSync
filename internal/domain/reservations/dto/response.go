package dto

import "time"

type CreateReservationResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    CreateReservation `json:"data"`
}
type CreateReservation struct {
	ID           uint      `json:"id"`
	UserID       uint      `json:"user_id"`
	ZoneID       uint      `json:"zone_id"`
	LicensePlate string    `json:"license_plate"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetAllReservationsResponse struct {
	Success bool                `json:"success"`
	Message string              `json:"message"`
	Data    []GetAllReservation `json:"data"`
}

type GetAllReservation struct {
	ID uint `json:"id"`

	LicensePlate string      `json:"license_plate"`
	Status       string      `json:"status"`
	Zone         ParkingZone `json:"zone"`
	CreatedAt    time.Time   `json:"created_at"`
}

type ParkingZone struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type DeleteReservationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
