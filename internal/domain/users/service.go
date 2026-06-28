package users

import (
	"SpotSync/internal/domain/users/dto"
	"time"
)

type service struct {
	repo Repository
}

func newService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) CreateUser(req *dto.RegistrationRequest) (*dto.RegistrationResponse, error) {
	user := User{
		Name:  req.Name,
		Email: req.Email,
	}

	err := user.hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	response := &dto.RegistrationResponse{
		Success: true,
		Message: "User registered successfully",
		Data: dto.RegisterUserResponse{
			ID:        int(user.ID),
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.UTC().Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.UTC().Format(time.RFC3339),
		},
	}

	return response, nil

}
