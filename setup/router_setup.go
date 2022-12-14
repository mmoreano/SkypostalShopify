package setup

import (
	"SkypostalBridgeApi/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func LaunchRoutes(api *fiber.App) {

	//Handles panics
	api.Use(recover.New())
	api.Post("/getRates", routes.GetRates)

}
