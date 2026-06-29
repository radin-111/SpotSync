package parking_zones

import "gorm.io/gorm"

type Repository interface {
	CreateParkingZone(parkingZone *ParkingZone) error
	// GetAllParkingZones() ([]ParkingZone, error)
	// GetParkingZoneByID(id uint) (*ParkingZone, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateParkingZone(parkingZone *ParkingZone) error {
	result := r.db.Create(parkingZone)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
