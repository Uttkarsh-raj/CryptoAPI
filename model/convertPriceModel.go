package model

type ConvertPriceModel struct {
	FromCurrency string `json:"fromCurrency"`
	ToCurrency   string `json:"toCurrency"`
	Date         string `json:"date"`
}

func ConvertFromJson(json map[string]string) *ConvertPriceModel {
	return &ConvertPriceModel{
		FromCurrency: json["fromCurrency"],
		ToCurrency:   json["toCurrency"],
		Date:         json["date"],
	}
}
