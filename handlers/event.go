package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theeeep/ticket-booking-be/models"
)

type Eventhandler struct {
	repository models.EventRepository
}

func NewEventhandler(router fiber.Router, repository models.EventRepository) {
	handler := &Eventhandler{
		repository: repository,
	}
	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:id", handler.GetOne)
}
