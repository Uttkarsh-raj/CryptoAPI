package model

type CoinWithPriceModel struct {
	Id         string          `json:"id"`
	Name       string          `json:"name"`
	Symbol     string          `json:"symbol"`
	MarketData MarketDataModel `json:"market_data"`
}

type MarketDataModel struct {
	CurrentPrice map[string]interface{} `json:"current_price"`
}

func ConvertJsonToCoinWithPriceModel(json map[string]interface{}) *CoinWithPriceModel {
	marketData := json["market_data"].(map[string]interface{})
	currentPrice := marketData["current_price"].(map[string]interface{})
	var marketDataModel MarketDataModel
	marketDataModel.CurrentPrice = currentPrice
	return &CoinWithPriceModel{
		Id:         json["id"].(string),
		Name:       json["name"].(string),
		Symbol:     json["symbol"].(string),
		MarketData: marketDataModel,
	}
}
