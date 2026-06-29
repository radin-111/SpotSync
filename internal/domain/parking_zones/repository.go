package parking_zones

import (
	"errors"

	"gorm.io/gorm"
)

const ReservationStatusActive = "active"

type Repository interface {
	CreateParkingZone(parkingZone *ParkingZone) error
	GetAllParkingZones() ([]SpotWithAvailableSpots, error)
	GetParkingZoneByID(id uint) (*SpotWithAvailableSpots, error)
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

func (r *repository) GetAllParkingZones() ([]SpotWithAvailableSpots, error) {
	var parkingZones []SpotWithAvailableSpots

	result := r.db.
		Model(&ParkingZone{}).
		Select(`
			parking_zones.*,
			total_capacity - (
				SELECT COUNT(*)
				FROM reservations
				WHERE reservations.zone_id = parking_zones.id
				AND reservations.status = ?
			) AS available_spots
		`, ReservationStatusActive).
		Scan(&parkingZones)

	if result.Error != nil {
		return nil, result.Error
	}

	return parkingZones, nil
}

func (r *repository) GetParkingZoneByID(id uint) (*SpotWithAvailableSpots, error) {

	var parkingZone SpotWithAvailableSpots
	result := r.db.
		Model(&ParkingZone{}).
		Select(`
			parking_zones.*,
			total_capacity - (
				SELECT COUNT(*)
				FROM reservations
				WHERE reservations.zone_id = parking_zones.id
				AND reservations.status = ?
			) AS available_spots
		`, ReservationStatusActive).
		Where("parking_zones.id = ?", id).
		First(&parkingZone)
	if result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("parking zone not found")
		}
		return nil, result.Error
	}
	return &parkingZone, nil
}
