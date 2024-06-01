package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joaoscorissa/mac-addr-lookup/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/lookup/:mac", controllers.LookupVendor)
}
