package user

import (
	"ecommerce-user/internal/core/config"
	"ecommerce-user/internal/core/context"
	"ecommerce-user/internal/repositories"
	"ecommerce-user/internal/request"
	"sync"

	"ecommerce-user/internal/models"
)

// Service service interface
type Service interface {
	CreateUser(c *context.Context, request *request.CreateProfile) error
}

type service struct {
	config         *config.Configs
	result         *config.ReturnResult
	userRepository repositories.UserRepository
	mutex          sync.Mutex
}

// NewService new service
func NewService() Service {
	return &service{
		config:         config.CF,
		result:         config.RR,
		userRepository: repositories.UserNewRepository(),
	}
}

// CreateUser create user
func (s *service) CreateUser(c *context.Context, request *request.CreateProfile) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	user := &models.User{
		Model: models.Model{
			ID: request.ID,
		},
		FirstName: request.FirstName,
		LastName:  request.LastName,
		ImageURL:  request.ImageURL,
	}

	err := s.userRepository.Create(c.GetDatabase(), user)
	if err != nil {
		return err
	}

	return nil
}
