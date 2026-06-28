package users

import (
	"SpotSync/internal/domain/users/dto"
	"SpotSync/internal/utils/auth"
	"fmt"
	"time"
)

type service struct {
	repo       Repository
	jwtService auth.JWTService
}

func newService(repo Repository, jwtService auth.JWTService) *service {
	return &service{repo: repo, jwtService: jwtService}
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

func (s *service) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if err := user.checkPassword(req.Password); err != nil {
		return nil, err
	}

	token, err := s.jwtService.GenerateToken(user.ID, user.Name, user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	response := &dto.LoginResponse{
		Success: true,
		Message: "Login successful",
		Data: dto.LoginUserData{
			Token: token,
			User: dto.LoginUserResponse{
				ID:    int(user.ID),
				Name:  user.Name,
				Email: user.Email,
				Role:  user.Role,
			},
		},
	}
	return response, nil
}
