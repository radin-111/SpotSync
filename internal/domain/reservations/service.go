package reservations

import (
	"SpotSync/internal/domain/reservations/dto"
	"errors"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) CreateReservation(reservation *dto.CreateReservationRequest, userId uint) (*dto.CreateReservationResponse, error) {
	reservationPayload := Reservation{
		UserId:       userId,
		ZoneId:       reservation.ZoneID,
		LicensePlate: reservation.LicensePlate,
		Status:       "active",
	}
	err := s.repo.CreateReservation(&reservationPayload)

	if err != nil {
		return nil, err
	}

	response := &dto.CreateReservationResponse{
		Success: true,
		Message: "Reservation created successfully",
		Data: dto.CreateReservation{
			ID:           reservationPayload.ID,
			UserID:       reservationPayload.UserId,
			ZoneID:       reservationPayload.ZoneId,
			LicensePlate: reservationPayload.LicensePlate,
			Status:       reservationPayload.Status,
			CreatedAt:    reservationPayload.CreatedAt,
			UpdatedAt:    reservationPayload.UpdatedAt,
		},
	}
	return response, nil

}

func (s *service) DeleteReservation(reservationId uint, userId uint) error {
	reservation, err := s.repo.GetReservationById(reservationId)
	if err != nil {
		return err
	}
	if reservation.UserId != userId {
		return errors.New("unauthorized")
	}
	err = s.repo.DeleteReservation(reservationId)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetAllReservationsByUserId(userId uint) (*dto.GetAllReservationsResponse, error) {
	var response dto.GetAllReservationsResponse
	reservations, err := s.repo.GetAllReservationsByUserId(userId)
	if err != nil {
		return nil, err
	}
	for _, reservation := range reservations {
		response.Data = append(response.Data, dto.GetAllReservation{
			ID: reservation.ID,

			LicensePlate: reservation.LicensePlate,
			Status:       reservation.Status,
			CreatedAt:    reservation.CreatedAt,
			Zone: dto.ParkingZone{
				ID:   reservation.Zone.ID,
				Name: reservation.Zone.Name,
				Type: reservation.Zone.Type,
			},
		})
	}
	return &response, nil
}
