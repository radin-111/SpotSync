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
