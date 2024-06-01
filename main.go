package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/joaoscorissa/mac-addr-lookup/routes"
    "log"
)

func main() {
    app := fiber.New()

    routes.SetupRoutes(app)

    log.Fatal(app.Listen(":3000"))
}
