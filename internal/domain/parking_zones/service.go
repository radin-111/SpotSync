package parking_zones

import (
	"SpotSync/internal/domain/parking_zones/dto"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) CreateParkingZone(parkingZone *dto.CreateParkingZoneRequest) (*dto.CreateParkingZoneResponse, error) {
	zone := ParkingZone{
		Name: parkingZone.Name,

		Type:          parkingZone.Type,
		TotalCapacity: parkingZone.TotalCapacity,
		PricePerHour:  parkingZone.PricePerHour,
	}

	err := s.repo.CreateParkingZone(&zone)
	if err != nil {
		return nil, err
	}
	response := &dto.CreateParkingZoneResponse{
		Success: true,
		Message: "Parking zone created successfully",
		Data: dto.ParkingZone{
			ID:            zone.ID,
			Name:          zone.Name,
			Type:          zone.Type,
			TotalCapacity: zone.TotalCapacity,
			PricePerHour:  zone.PricePerHour,
			CreatedAt:     zone.CreatedAt,
			UpdatedAt:     zone.UpdatedAt,
		},
	}
	return response, nil
}

func (s *service) GetAllParkingZones() (*dto.GetAllParkingZonesResponse, error) {
	zones, err := s.repo.GetAllParkingZones()
	if err != nil {
		return nil, err
	}
	response := &dto.GetAllParkingZonesResponse{
		Success: true,
		Message: "Parking zones retrieved successfully",
	}

	for _, zone := range zones {
		response.Data = append(response.Data, dto.SingleParkingZone{
			ID:             zone.ID,
			Name:           zone.Name,
			Type:           zone.Type,
			TotalCapacity:  zone.TotalCapacity,
			AvailableSpots: zone.AvailableSpots,
			PricePerHour:   zone.PricePerHour,
			CreatedAt:      zone.CreatedAt,
		})
	}
	return response, nil
}

func (s *service) GetParkingZoneByID(id uint) (*dto.GetParkingZoneByIDResponse, error) {
	zone, err := s.repo.GetParkingZoneByID(id)
	if err != nil {
		return nil, err
	}
	response := &dto.GetParkingZoneByIDResponse{
		Success: true,
		Message: "Parking zone retrieved successfully",
		Data: dto.SingleParkingZone{
			ID:             zone.ID,
			Name:           zone.Name,
			Type:           zone.Type,
			TotalCapacity:  zone.TotalCapacity,
			AvailableSpots: zone.AvailableSpots,
			PricePerHour:   zone.PricePerHour,
			CreatedAt:      zone.CreatedAt,
		},
	}
	return response, nil
}
