package helper

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

func GetCoinByIdAndDate(id, date string) (*model.CoinWithPriceModel, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			// http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()
	queryUrl := "https://api.coingecko.com/api/v3/coins/" + string(id) + "/history?date=" + string(date)
	res, err := httpClient.Get(queryUrl)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}
	defer res.Body.Close()
	if res.StatusCode == 401 {
		return nil, fmt.Errorf("error: Public API users are limited to querying historical data within the past 365 days")
	}
	if res.StatusCode == 429 {
		return nil, fmt.Errorf("error: You've exceeded the Rate Limit")
	}
	if res.StatusCode != http.StatusOK {
		var errorResponse map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}
		return nil, errors.New(errorResponse["error"].(string))
	}

	var jsonBody map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&jsonBody)
	if err != nil {
		return nil, err
	}

	coinWithPrice := model.ConvertJsonToCoinWithPriceModel(jsonBody)

	return coinWithPrice, nil
}

func FetchComapaniesFromId(id string) ([]*model.Company, error) {
	url := "https://api.coingecko.com/api/v3/companies/public_treasury/" + id
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 401 {
		return nil, fmt.Errorf("error: Public API users are limited to querying historical data within the past 365 days")
	}
	if resp.StatusCode == 429 {
		return nil, fmt.Errorf("error: You've exceeded the Rate Limit")
	}
	if resp.StatusCode != http.StatusOK {
		var errorResponse map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}
		return nil, errors.New(errorResponse["error"].(string))
	}
	var jsonResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		return nil, err
	}

	var companies []*model.Company
	companiesData, err := json.Marshal(jsonResponse["companies"])
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(companiesData, &companies)
	if err != nil {
		return nil, err
	}

	return companies, nil
}
