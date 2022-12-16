package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"
)

//region ShopifyIncomingRateRequest

type ShopifyGetRatesRequest struct {
	Rate shopifyRate `json:"rate"`
}

type shopifyRate struct {
	Origin      shopifyOrigin      `json:"origin"`
	Destination shopifyDestination `json:"destination"`
	Items       []shopifyItem      `json:"items"`
	Currency    string             `json:"currency"`
	Locale      string             `json:"locale"`
}

type shopifyOrigin struct {
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

type shopifyDestination struct {
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

type shopifyItem struct {
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

func (req *ShopifyGetRatesRequest) Validate() []*fiber.Error {
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

//endregion ShopifyIncomingRateRequest

//region SkypostalGetRatesRequest

type SpGetRatesRequest struct {
	UserInfo         SkypostalUserInfo
	Weight           float64
	WeightType       string
	MerchandiseValue float64
	CopaID           int
	CountryCode      int
	CityCode         int
	HeightDim        float64
	LengthDim        float64
	WidthDim         float64
	DimType          string
	CouponCode       string
	IataCodeOrigin   string
	ZipCode          string
	RateServiceCode  int
	InsuranceCode    int
}

func (skypostalObj *SpGetRatesRequest) PopulateFromShopify(shopifyObj *ShopifyGetRatesRequest, additionalObj *SpGetRatesAdditionalRequest) {

	skypostalObj.UserInfo = SkypostalUserInfo{
		UserCode: 0,
		AppKey:   "",
		UserKey:  "",
	}

	for _, item := range shopifyObj.Rate.Items {
		skypostalObj.Weight += float64(item.Grams) * 0.00220462 //Converting to LB
	}

	skypostalObj.WeightType = "LB"

	for _, item := range shopifyObj.Rate.Items {
		skypostalObj.MerchandiseValue += float64(item.Price)
	}

}

//endregion SkypostalGetRatesRequest

//region SkypostalGetRatesAdditionalRequest

type SpGetRatesAdditionalRequest struct {
	UserInfo     SkypostalUserInfo
	ShipmentInfo spShipmentInfo
}

func (skypostalObj *SpGetRatesAdditionalRequest) PopulateFromShopify(shopifyObj *ShopifyGetRatesRequest) {

	skypostalObj.UserInfo = SkypostalUserInfo{
		UserCode: 120,
		AppKey:   "zgo4oD0DiMOVN02172dhMXC4o739TwdH",
		UserKey:  "",
	}
	skypostalObj.ShipmentInfo.CopaID = 616 //??
	skypostalObj.ShipmentInfo.Consignee.Address.CountryIsoCode = shopifyObj.Rate.Destination.Country
	skypostalObj.ShipmentInfo.Consignee.Address.StateName = shopifyObj.Rate.Destination.Province
	skypostalObj.ShipmentInfo.Consignee.Address.CityName = shopifyObj.Rate.Destination.City
	skypostalObj.ShipmentInfo.Consignee.Address.ZipCode = shopifyObj.Rate.Destination.PostalCode
	skypostalObj.ShipmentInfo.Consignee.Address.Address01 = shopifyObj.Rate.Destination.Address1 + shopifyObj.Rate.Destination.Address2 + shopifyObj.Rate.Destination.Address3

}

type spShipmentInfo struct {
	CopaID    int
	Consignee spConsignee
	Options   spOptions
}

type spConsignee struct {
	FirstName      string
	LastName       string
	Email          string
	IdNumber       string
	IdSearchString string
	Address        spAddress
	Phone          []spPhone
}

type spAddress struct {
	CountryIsoCode string
	StateName      string
	CityName       string
	ZipCode        string
	Address01      string
	Address02      string
}

type spPhone struct {
	PhoneType      int
	PhoneNumber    int
	Error          SkypostalError
	AdditionalInfo SkypostalAdditionalInfo
}

type spOptions struct {
	GenerateLabelDefault bool
}

//endregion SkypostalGetRatesAdditionalRequest\

//region SkypostalGetRatesAdditionalResponse

type SpGetRatesAdditionalResponse struct {
	BoxCountryCopaInfo        []spBoxCountryCopaInfo        `json:"box_country_copa_info"`
	ConnGeoProviders          []spConnGeoProviders          `json:"conn_geo_providers"`
	ConnGeoProvidersFirstMile []spConnGeoProvidersFirstMile `json:"conn_geo_providers_first_mile"`
	CurrencyConversions       []spCurrencyConversions       `json:"currency_conversions"`
	DestinationCity           spDestinationCity             `json:"destination_city"`
	Error                     []SkypostalError              `json:"error"`
	HscodeFmprRelationsData   spHscodeFmprRelationsData     `json:"hscode_fmpr_relations_data"`
	PoeAddressData            []spPoeAddressData            `json:"poe_address_data"`
}

type spBoxCountryCopaInfo struct {
	Verify                 bool                    `json:"_verify"`
	AdditionalInfo         SkypostalAdditionalInfo `json:"additional_info"`
	BoxsID                 int                     `json:"boxs_id"`
	BoxsIDOrig             int                     `json:"boxs_id_orig"`
	CityCodeUnknown        int                     `json:"city_code_unknown"`
	CityIataCode           string                  `json:"city_iata_code"`
	CityName               string                  `json:"city_name"`
	CopaAllowUnknownCities int                     `json:"copa_allow_unknown_cities"`
	CopaBusinessType       int                     `json:"copa_business_type"`
	CopaDdupGetMcdService  int                     `json:"copa_ddup_get_mcd_service"`
	CountryCode            int                     `json:"country_code"`
	CountryName            string                  `json:"country_name"`
	Error                  SkypostalError          `json:"error"`
}
type spProviderServiceType struct {
	ServiceAddInfo         string `json:"service_add_info"`
	ServiceAddInfo2        string `json:"service_add_info2"`
	ServiceAddInfo3        string `json:"service_add_info3"`
	ServiceProviderCode    string `json:"service_provider_code"`
	ServiceTypeCode        int    `json:"service_type_code"`
	ServiceTypeDescription string `json:"service_type_description"`
	ServiceTypeID          string `json:"service_type_id"`
	ServiceTypeNumber      string `json:"service_type_number"`
}
type spProviderSettings struct {
	PostingCard                interface{} `json:"posting_card"`
	ServiceProviderCode        string      `json:"service_provider_code"`
	SettingsAdministrativeCode string      `json:"settings_administrative_code"`
	SettingsCardNumber         string      `json:"settings_card_number"`
	SettingsCepOrigin          string      `json:"settings_cep_origin"`
	SettingsCode               string      `json:"settings_code"`
	SettingsContractNumber     string      `json:"settings_contract_number"`
	SettingsEnvironment        string      `json:"settings_environment"`
	SettingsPostingCardDef     string      `json:"settings_posting_card_def"`
	SettingsRemAddress         string      `json:"settings_rem_address"`
	SettingsRemCity            string      `json:"settings_rem_city"`
	SettingsRemComplement      string      `json:"settings_rem_complement"`
	SettingsRemContact         string      `json:"settings_rem_contact"`
	SettingsRemCountryCode     string      `json:"settings_rem_country_code"`
	SettingsRemName            string      `json:"settings_rem_name"`
	SettingsRemNeighborhood    string      `json:"settings_rem_neighborhood"`
	SettingsRemNumber          string      `json:"settings_rem_number"`
	SettingsRemPhone           string      `json:"settings_rem_phone"`
	SettingsRemUf              string      `json:"settings_rem_uf"`
	SettingsRemZipcode         string      `json:"settings_rem_zipcode"`
	SettingsTrackingPassword   string      `json:"settings_tracking_password"`
	SettingsTrackingUser       string      `json:"settings_tracking_user"`
	SettingsUserName           string      `json:"settings_user_name"`
	SettingsUserPassword       string      `json:"settings_user_password"`
}
type spConnGeoProviders struct {
	Verify                    bool                    `json:"_verify"`
	AdditionalInfo            SkypostalAdditionalInfo `json:"additional_info"`
	ConnProviderEnvironment   string                  `json:"conn_provider_environment"`
	ConnProviderParam1        string                  `json:"conn_provider_param1"`
	ConnProviderParam2        string                  `json:"conn_provider_param2"`
	ConnProviderParam3        string                  `json:"conn_provider_param3"`
	ConnProviderPassword      string                  `json:"conn_provider_password"`
	ConnProviderToken         string                  `json:"conn_provider_token"`
	ConnProviderUser          string                  `json:"conn_provider_user"`
	CopaCtryServiceDefault    int                     `json:"copa_ctry_service_default"`
	CopaCtryServiceOrder      int                     `json:"copa_ctry_service_order"`
	DduEnabled                bool                    `json:"ddu_enabled"`
	Error                     SkypostalError          `json:"error"`
	ProviderServiceType       spProviderServiceType   `json:"provider_service_type"`
	ProviderSettings          spProviderSettings      `json:"provider_settings"`
	RateServiceCode           int                     `json:"rate_service_code"`
	ServiceProviderCode       int                     `json:"service_provider_code"`
	ServiceProviderCountryDef int                     `json:"service_provider_country_def"`
	ServiceTypeCode           int                     `json:"service_type_code"`
	ZipCovered                int                     `json:"zip_covered"`
}
type spConnGeoProvidersFirstMile struct {
	Verify                    bool                    `json:"_verify"`
	AdditionalInfo            SkypostalAdditionalInfo `json:"additional_info"`
	ConnProviderEnvironment   string                  `json:"conn_provider_environment"`
	ConnProviderParam1        string                  `json:"conn_provider_param1"`
	ConnProviderParam2        string                  `json:"conn_provider_param2"`
	ConnProviderParam3        string                  `json:"conn_provider_param3"`
	ConnProviderPassword      string                  `json:"conn_provider_password"`
	ConnProviderToken         string                  `json:"conn_provider_token"`
	ConnProviderUser          string                  `json:"conn_provider_user"`
	CopaCtryServiceDefault    int                     `json:"copa_ctry_service_default"`
	CopaCtryServiceOrder      int                     `json:"copa_ctry_service_order"`
	DduEnabled                bool                    `json:"ddu_enabled"`
	Error                     SkypostalError          `json:"error"`
	ProviderServiceType       spProviderServiceType   `json:"provider_service_type"`
	ProviderSettings          spProviderSettings      `json:"provider_settings"`
	RateServiceCode           int                     `json:"rate_service_code"`
	ServiceProviderCode       int                     `json:"service_provider_code"`
	ServiceProviderCountryDef int                     `json:"service_provider_country_def"`
	ServiceTypeCode           int                     `json:"service_type_code"`
	ZipCovered                int                     `json:"zip_covered"`
}
type spCurrencyConversions struct {
	Verify          bool                    `json:"_verify"`
	AdditionalInfo  SkypostalAdditionalInfo `json:"additional_info"`
	BuyRate         float64                 `json:"buy_rate"`
	CountryCode     int                     `json:"country_code"`
	CurrencyIsoCode string                  `json:"currency_iso_code"`
	Error           SkypostalError          `json:"error"`
	SellRate        float64                 `json:"sell_rate"`
}

type spDestinationCity struct {
	Verify                   bool                    `json:"_verify"`
	AdditionalInfo           SkypostalAdditionalInfo `json:"additional_info"`
	CityCode                 int                     `json:"city_code"`
	CityIataCode             string                  `json:"city_iata_code"`
	CityLocalID              string                  `json:"city_local_id"`
	CityName                 string                  `json:"city_name"`
	CityService              string                  `json:"city_service"`
	CityUniqueID             int                     `json:"city_unique_id"`
	CodeFound                int                     `json:"code_found"`
	CountryCode              int                     `json:"country_code"`
	CountryName              string                  `json:"country_name"`
	CountyLocalID            string                  `json:"county_local_id"`
	Error                    SkypostalError          `json:"error"`
	IsCityUnknown            bool                    `json:"is_city_unknown"`
	IsCityUnknownCopaAllowed bool                    `json:"is_city_unknown_copa_allowed"`
	StateAbreviation         string                  `json:"state_abreviation"`
	StateCode                int                     `json:"state_code"`
	StateLocalID             string                  `json:"state_local_id"`
	StateName                string                  `json:"state_name"`
}
type spHscodeFmprRelations struct {
	FmprCdg     string `json:"fmpr_cdg"`
	FmprDescEng string `json:"fmpr_desc_eng"`
	FmprDescEsp string `json:"fmpr_desc_esp"`
	Hscode      string `json:"hscode"`
}
type spHscodeFmprRelationsData struct {
	HscodeFmprRelations []spHscodeFmprRelations `json:"hscode_fmpr_relations"`
}
type spPoeAddress struct {
	AdditionalInfo    SkypostalAdditionalInfo `json:"additional_info"`
	Address01         string                  `json:"address_01"`
	Address02         string                  `json:"address_02"`
	Address03         string                  `json:"address_03"`
	CityCode          int                     `json:"city_code"`
	CityName          string                  `json:"city_name"`
	CountryCode       int                     `json:"country_code"`
	CountryIsoCode    string                  `json:"country_iso_code"`
	CountryName       string                  `json:"country_name"`
	CountyCode        int                     `json:"county_code"`
	CountyName        string                  `json:"county_name"`
	Error             SkypostalError          `json:"error"`
	Neighborhood      string                  `json:"neighborhood"`
	StateAbbreviation string                  `json:"state_abbreviation"`
	StateCode         int                     `json:"state_code"`
	StateName         string                  `json:"state_name"`
	TownCode          int                     `json:"town_code"`
	TownName          string                  `json:"town_name"`
	ZipCode           string                  `json:"zip_code"`
}
type spPoePhone struct {
	AdditionalInfo SkypostalAdditionalInfo `json:"additional_info"`
	Error          SkypostalError          `json:"error"`
	PhoneExtension interface{}             `json:"phone_extension"`
	PhoneNumber    string                  `json:"phone_number"`
	PhoneType      int                     `json:"phone_type"`
}
type spPoeAddressData struct {
	Verify         bool                    `json:"_verify"`
	AdditionalInfo SkypostalAdditionalInfo `json:"additional_info"`
	CopaID         int                     `json:"copa_id"`
	Error          SkypostalError          `json:"error"`
	Name           string                  `json:"name"`
	PoeAddress     spPoeAddress            `json:"poe_address"`
	PoeIata        string                  `json:"poe_iata"`
	PoePhone       spPoePhone              `json:"poe_phone"`
	PoeReturnType  string                  `json:"poe_return_type"`
}

//endregion SkypostalGetRatesAdditionalResponse

//region SkypostalApiGetRatesResponse

type SpGetRatesResponse struct {
	Data           []spGetRatesResponseBody `json:"data"`
	Error          SkypostalError           `json:"error"`
	AdditionalInfo SkypostalAdditionalInfo  `json:"additional_info"`
}

type spGetRatesResponseBody struct {
	Verify                        bool                    `json:"_verify"`
	AdditionalDiscount            float64                 `json:"additional_discount"`
	CustomValue                   float64                 `json:"custom_value"`
	CustomValueAdditional         float64                 `json:"custom_value_additional"`
	ExtraValue                    float64                 `json:"extra_value"`
	FmprCdg                       string                  `json:"fmpr_cdg"`
	FuelSurcharge                 float64                 `json:"fuel_surcharge"`
	HandlingFee                   float64                 `json:"handling_fee"`
	InsuranceValue                float64                 `json:"insurance_value"`
	Is100PreDiscount              float64                 `json:"is_100pre_discount"`
	PromotionCode                 string                  `json:"promotion_code"`
	RateDisComPart                int                     `json:"rate_dis_com_part"`
	RateDscBeyond                 int                     `json:"rate_dsc_beyond"`
	RateDscCommercialPartner      string                  `json:"rate_dsc_commercial_partner"`
	RateDscGateway                int                     `json:"rate_dsc_gateway"`
	RateDscValueCommercialPartner int                     `json:"rate_dsc_value_commercial_partner"`
	RateVlrGateway                float64                 `json:"rate_vlr_gatewaty"` //misspelled as "gatewaty" on API
	ShipDiscount                  float64                 `json:"ship_discount"`
	ShipSubtotalRate              float64                 `json:"ship_subtotal_rate"`
	ShipTotalRate                 float64                 `json:"ship_total_rate"`
	TotalCustoms                  float64                 `json:"total_customs"`
	TotalShipping                 float64                 `json:"total_shipping"`
	TotalValue                    float64                 `json:"total_value"`
	DefaultRateApplied            bool                    `json:"default_rate_applied"`
	FirstMileRate                 spFirstMileRate         `json:"first_mile_rate"`
	Error                         SkypostalError          `json:"error"`
	AdditionalInfo                SkypostalAdditionalInfo `json:"additional_info"`
}

type spFirstMileRate struct {
	Weight                float64 `json:"weight"`
	WeightUnit            string  `json:"weight_unit"`
	PricingWeight         float64 `json:"pricing_weight"`
	PricingCubicFt        float64 `json:"pricing_cubic_ft"`
	Value                 float64 `json:"value"`
	PostageAmount         float64 `json:"postage_amount"`
	FeesAmount            float64 `json:"fees_amount"`
	DiscountedAmount      float64 `json:"discounted_amount"`
	TotalAmount           float64 `json:"total_amount"`
	EstimatedDeliveryDays int     `json:"estimated_delivery_days"`
}

//endregion SkypostalApiGetRatesResponse
