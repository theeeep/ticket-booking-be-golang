package handlers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/theeeep/ticket-booking-be/models"
)

type Eventhandler struct {
	repository models.EventRepository
}

func (h *Eventhandler) GetMany(ctx *fiber.Ctx) error {
	context, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	events, err := h.repository.GetMany(context)
	if err != nil {
		return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
			"data":    events,
		})
	}

	return ctx.Status(fiber.StatusBadGateway).JSON(&fiber.Map{
		"status":  "success",
		"message": "events fetched successfully",
		"data":    events,
	})
}

func (h *Eventhandler) CreateOne(ctx *fiber.Ctx) error {
	return nil
}

func (h *Eventhandler) GetOne(ctx *fiber.Ctx) error {
	return nil
}

func NewEventhandler(router fiber.Router, repository models.EventRepository) {
	handler := &Eventhandler{
		repository: repository,
	}
	router.Get("/", handler.GetMany)
	router.Post("/", handler.CreateOne)
	router.Get("/:id", handler.GetOne)
}
