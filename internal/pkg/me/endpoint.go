// Package me endpoint
package me

import (
	"ecommerce-user/internal/core/config"
	"ecommerce-user/internal/handlers"
	"ecommerce-user/internal/models"

	"github.com/gofiber/fiber/v2"
)

// Endpoint endpoint interface
type Endpoint interface {
	GetProfile(c *fiber.Ctx) error
	FindAllOrder(c *fiber.Ctx) error
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

// GetProfile get profile
// @Tags Me
// @Summary GetProfile
// @Description GetProfile
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Success 200 {object} models.User
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Security ApiKeyAuth
// @Router /me [get]
func (ep *endpoint) GetProfile(c *fiber.Ctx) error {
	return handlers.ResponseObjectWithoutRequest(c, ep.service.GetProfile)
}

// FindAllOrder find all order
// @Tags Me
// @Summary FindAllOrder
// @Description FindAllOrder
// @Accept json
// @Produce json
// @Param Accept-Language header string false "(en, th)" default(th)
// @Param request query models.PageForm true "request"
// @Success 200 {object} models.Page
// @Failure 400 {object} models.Message
// @Failure 401 {object} models.Message
// @Failure 404 {object} models.Message
// @Failure 410 {object} models.Message
// @Security ApiKeyAuth
// @Router /me/orders [get]
func (ep *endpoint) FindAllOrder(c *fiber.Ctx) error {
	return handlers.ResponseObject(c, ep.service.FindOrders, &models.PageForm{})
}
