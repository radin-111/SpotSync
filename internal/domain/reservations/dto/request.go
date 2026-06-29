package dto

type CreateReservationRequest struct {
	ZoneID       uint   `json:"zone_id"`
	LicensePlate string `json:"license_plate"`
}
