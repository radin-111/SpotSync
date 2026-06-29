package reservations

import "gorm.io/gorm"

type Repository interface {
	CreateReservation(reservation *Reservation) error
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
