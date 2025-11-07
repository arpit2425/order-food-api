package routes

import (
	"github.com/gofiber/fiber/v2"
	"oilio.com/internal/store"
)

func SetupRoutes(app *fiber.App, store store.Store) {
	api := app.Group("/api")
	RegisterProductRoutes(api, store)
	RegisterOrderRoutes(api, store)

}
