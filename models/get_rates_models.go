package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type GetRatesRequest struct {
	Rate rate `json:"rate"`
}

type rate struct {
	Origin      origin      `json:"origin"`
	Destination destination `json:"destination"`
	Items       []item      `json:"items"`
	Currency    string      `json:"currency"`
	Locale      string      `json:"locale"`
}

type origin struct {
	Country     string `json:"country"`
	PostalCode  string `json:"postal_code"`
	Province    string `json:"province"`
	City        string `json:"city"`
	Name        string `json:"name"`
	Address1    string `json:"address1" validate:"required"`
	Address2    string `json:"address2" validate:"required"`
	Address3    string `json:"address3"`
	Phone       string `json:"phone"`
	Fax         string `json:"fax"`
	Email       string `json:"email"`
	AddressType string `json:"address_type"`
	CompanyName string `json:"company_name"`
}

type destination struct {
	Country     string `json:"country"`
	PostalCode  string `json:"postal_code"`
	Province    string `json:"province"`
	City        string `json:"city"`
	Name        string `json:"name"`
	Address1    string `json:"address1" validate:"required"`
	Address2    string `json:"address2" validate:"required"`
	Address3    string `json:"address3"`
	Phone       string `json:"phone"`
	Fax         string `json:"fax"`
	Email       string `json:"email"`
	AddressType string `json:"address_type"`
	CompanyName string `json:"company_name"`
}

type item struct {
	Name               string `json:"name"`
	Sku                string `json:"sku"`
	Quantity           string `json:"quantity"`
	Grams              int    `json:"grams"`
	Price              int    `json:"price"`
	Vendor             string `json:"vendor"`
	RequiresShipping   bool   `json:"requires_shipping"`
	Taxable            bool   `json:"taxable"`
	FulfillmentService string `json:"fulfillment_service"`
	Properties         string `json:"properties"`
	ProductId          int    `json:"product_id"`
	VariantId          int    `json:"variant_id"`
}

func (req *GetRatesRequest) Validate() []*fiber.Error {
	var errors []*fiber.Error
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element fiber.Error
			element.Code = fiber.StatusBadRequest
			element.Message = `Error (` + err.Tag() + `) in JSON field`
			field := strings.SplitAfter(err.StructNamespace(), ".")
			element.Message += ` (` + field[len(field)-1] + `)`
			errors = append(errors, &element)
		}
	}
	return errors
}
