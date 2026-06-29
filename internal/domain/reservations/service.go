package reservations

import "SpotSync/internal/domain/reservations/dto"

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
