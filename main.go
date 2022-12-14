package main

import (
	"SkypostalBridgeApi/setup"
	"github.com/gofiber/fiber/v2"
)

func main() {
	api := fiber.New()
	setup.LaunchRoutes(api)
	api.Listen(":6969")
}
