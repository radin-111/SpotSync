package reservations

import (
	"SpotSync/internal/domain/parking_zones"
	"SpotSync/internal/domain/users"

	"gorm.io/gorm"
)

const (
	ReservationStatusActive    = "active"
	ReservationStatusCompleted = "completed"
	ReservationStatusCancelled = "cancelled"
)

type Reservation struct {
	gorm.Model
	UserId       uint                      `json:"user_id" gorm:"not null"`
	User         users.User                `json:"user" gorm:"foreignKey:UserId;references:Id;constraint:OnDelete:CASCADE"`
	ZoneId       uint                      `json:"zone_id" gorm:"not null"`
	Zone         parking_zones.ParkingZone `json:"zone" gorm:"foreignKey:ZoneId;references:Id;constraint:OnDelete:CASCADE"`
	LicensePlate string                    `json:"license_plate" gorm:"type:varchar(20);not null;uniqueIndex:license_plate"`
	Status       string                    `json:"status" gorm:"type:varchar(20);default:active"`
}
