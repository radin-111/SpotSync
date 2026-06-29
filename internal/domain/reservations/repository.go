package reservations

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	CreateReservation(reservation *Reservation) error
	GetAllReservationsByUserId(userId uint) ([]Reservation, error)
	DeleteReservation(reservationId uint) error
	GetReservationById(reservationId uint) (*Reservation, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateReservation(reservation *Reservation) error {
	result := r.db.Create(reservation)
	return result.Error

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
