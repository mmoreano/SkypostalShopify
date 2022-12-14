package routes

import (
	"SkypostalBridgeApi/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

var getRatesRequest models.GetRatesRequest

func GetRates(ctx *fiber.Ctx) error {

	json.Unmarshal(ctx.Body(), &getRatesRequest)

	err := getRatesRequest.Validate()

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	return ctx.JSON(getRatesRequest)

	//urlExtension := "/wcf-services/service-user.svc/user/user-login"
	//url := urlExtension + setup.TestEnv
	//
	//method := "POST"
	//
	//client := &http.Client{}
	//req, err := http.NewRequest(method, url, nil)
	//
	//if err != nil {
	//	fmt.Println("handle ERRROR")
	//}
	//
	//res, err := client.Do(req)

}
