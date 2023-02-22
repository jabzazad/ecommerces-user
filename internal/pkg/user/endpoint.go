// Package guest package
package user

import (
	"ecommerce-user/internal/core/config"
	"ecommerce-user/internal/handlers"
	"ecommerce-user/internal/request"

	"github.com/gofiber/fiber/v2"
)

// Endpoint endpoint interface
type Endpoint interface {
	CreateUser(c *fiber.Ctx) error
}

type endpoint struct {
	config  *config.Configs
	result  *config.ReturnResult
	service Service
}

// NewEndpoint new endpoint
func NewEndpoint() Endpoint {
	return &endpoint{
		config:  config.CF,
		result:  config.RR,
		service: NewService(),
	}
}

func (ep *endpoint) CreateUser(c *fiber.Ctx) error {
	return handlers.ResponseSuccess(c, ep.service.CreateUser, &request.CreateProfile{})
}
