package reservations

import (
	"SpotSync/internal/domain/parking_zones"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	CreateReservation(reservation *Reservation) error
	GetAllReservationsByUserId(userId uint) ([]Reservation, error)
	DeleteReservation(reservationId uint) error
	GetAllReservations() ([]Reservation, error)
	GetReservationById(reservationId uint) (*Reservation, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateReservation(reservation *Reservation) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var zone parking_zones.SpotWithAvailableSpots

		result := r.db.Model(&parking_zones.ParkingZone{}).
			Where("id = ?", reservation.ZoneId).
			Select(`parking_zones.*,
			total_capacity - (
				SELECT COUNT(*)
				FROM reservations
				WHERE reservations.zone_id = parking_zones.id
				AND reservations.status = ?
			) AS available_spots
		`, ReservationStatusActive).
			Scan(&zone)

		if result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				return errors.New("zone not found")
			}
			return result.Error
		}

		if zone.AvailableSpots <= 0 {
			return errors.New("no available spots")
		}
		if err := tx.Create(reservation).Error; err != nil {
			return err
		}
		return nil
	})

}
func (r *repository) GetAllReservationsByUserId(userId uint) ([]Reservation, error) {
	var reservations []Reservation

	result := r.db.
		Preload("Zone").
		Where("user_id = ?", userId).
		Find(&reservations)

	if result.Error != nil {
		return nil, result.Error
	}

	return reservations, nil
}
func (r *repository) GetAllReservations() ([]Reservation, error) {
	var reservations []Reservation

	result := r.db.
		Preload("Zone").
		Find(&reservations)

	if result.Error != nil {
		return nil, result.Error
	}

	return reservations, nil
}
func (r *repository) DeleteReservation(reservationId uint) error {
	result := r.db.Delete(&Reservation{}, reservationId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return errors.New("reservation not found")
		}
		return result.Error
	}
	return nil
}
func (r *repository) GetReservationById(reservationId uint) (*Reservation, error) {
	var reservation Reservation
	result := r.db.Where("id = ?", reservationId).First(&reservation)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("reservation not found")
		}
		return nil, result.Error
	}
	return &reservation, nil
}
