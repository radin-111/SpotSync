package parking_zones

import "gorm.io/gorm"

const (
	ParkingZoneTypeGeneral    = "general"
	ParkingZoneTypeEVCharging = "ev_charging"
	ParkingZoneTypeCovered    = "covered"
)

type ParkingZone struct {
	gorm.Model
	Name          string  `json:"name" gorm:"type:varchar(100);not null"`
	Type          string  `json:"type" gorm:"type:varchar(20);default:general"`
	TotalCapacity int     `json:"total_capacity" gorm:"type:int;default:0"`
	PricePerHour  float64 `json:"price_per_hour" gorm:"type:decimal(10,2);default:0.00"`
}

type SpotWithAvailableSpots struct {
	gorm.Model
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	TotalCapacity  int     `json:"total_capacity"`
	PricePerHour   float64 `json:"price_per_hour"`
	AvailableSpots int     `json:"available_spots"`
}
