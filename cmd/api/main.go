package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/theeeep/ticket-booking-be/config"
	"github.com/theeeep/ticket-booking-be/db"
	"github.com/theeeep/ticket-booking-be/handlers"
	"github.com/theeeep/ticket-booking-be/repositories"
)

func main() {

	// Config
	envConfig := config.NewEnvConfig()

	// DB
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "Ticket Booking",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(db)

	// Routing
	server := app.Group("/api")

	// Handlers
	handlers.NewEventhandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprintf("%s", ":"+envConfig.ServerPort))
}
