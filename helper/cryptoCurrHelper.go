package helper

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Uttkarsh-raj/Crypto/model"
)

var (
	httpClient = &http.Client{}
)

func GetAllCoins() []*model.CoinModel {
	_, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	var cryptoCoins []*model.CoinModel

	response, err := httpClient.Get("https://api.coingecko.com/api/v3/coins/list")
	if err != nil {
		log.Fatalf("Unable to fetch the coins data: %s", err)
		return cryptoCoins
	}
	defer response.Body.Close()

	var coinResponses []map[string]string

	if err := json.NewDecoder(response.Body).Decode(&coinResponses); err != nil {
		log.Fatalf("Unable to fetch the coins data: %s", err)
		return cryptoCoins
	}

	cryptoCoins = model.GetCoinsList(coinResponses)
	return cryptoCoins
}
