package me

import (
	"ecommerce-user/internal/core/config"
	"ecommerce-user/internal/core/context"
	"ecommerce-user/internal/models"
	"ecommerce-user/internal/pkg/client"
	"ecommerce-user/internal/repositories"
	"fmt"

	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
)

// Service service interface
type Service interface {
	GetProfile(c *context.Context) (*models.User, error)
	FindOrders(c *context.Context, request *models.PageForm) (*models.Page, error)
}

type service struct {
	config         *config.Configs
	result         *config.ReturnResult
	userRepository repositories.UserRepository
	clientService  client.Service
}

// NewService new service
func NewService() Service {
	return &service{
		config:         config.CF,
		result:         config.RR,
		userRepository: repositories.UserNewRepository(),
		clientService:  client.NewService(),
	}
}

// GetProfile get profile
func (s *service) GetProfile(c *context.Context) (*models.User, error) {
	user, err := s.userRepository.FindOneByIDWithPreload(c.GetDatabase(), c.GetUserID())
	if err != nil {
		logrus.Errorf("get profile error: %s", err)
		return nil, s.result.Internal.DatabaseNotFound
	}

	return user, nil
}

// FindOrders find all orders
func (s *service) FindOrders(c *context.Context, form *models.PageForm) (*models.Page, error) {
	header := req.Header{
		"Content-Type":  "application/json",
		"Authorization": c.Get("Authorization"),
	}

	request := req.Param{
		"page": form.Page,
		"size": form.Size,
	}

	response := &models.Page{}
	url := fmt.Sprintf("%s%s", s.config.Order.URL, s.config.Order.Path.Order)
	err := s.clientService.GetRequest(url, header, request, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
