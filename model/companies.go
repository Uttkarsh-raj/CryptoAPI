package model

type Company struct {
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	Country       string  `json:"country"`
	TotalHoldings float64 `json:"total_holdings"`
}

func GetCompany(json map[string]interface{}) *Company {
	return &Company{
		Name:          json["name"].(string),
		Symbol:        json["symbol"].(string),
		Country:       json["country"].(string),
		TotalHoldings: json["total_holdings"].(float64),
	}
}

func GetListOfCompanies(json []map[string]interface{}) []*Company {
	var listOfCompanies []*Company
	for i := 0; i < len(json); i++ {
		company := GetCompany(json[i])
		listOfCompanies = append(listOfCompanies, company)
	}
	return listOfCompanies
}
