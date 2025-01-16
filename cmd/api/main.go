package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/theeeep/ticket-booking-be/handlers"
	"github.com/theeeep/ticket-booking-be/repositories"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "Ticket Booking",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(nil)

	// Routing
	server := app.Group("/api")

	// Handlers
	handlers.NewEventhandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}
