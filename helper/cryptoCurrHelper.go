package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Uttkarsh-raj/Crypto/model"
)

var (
	httpClient = &http.Client{}
)

func GetAllCoins() []*model.CoinModel {
	_, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	var cryptoCoins []*model.CoinModel

	req, err := http.NewRequest("GET", "https://api.coingecko.com/api/v3/coins/list", nil)
	if err != nil {
		log.Fatalf("Unable to create request: %s", err)
		return cryptoCoins
	}
	token := os.Getenv("APIKEY")
	req.Header.Set("x_cg_demo_api_key", token)

	response, err := httpClient.Do(req)
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

	req, err := http.NewRequest("GET", "https://api.coingecko.com/api/v3/coins/"+id+"/history?date="+date, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err.Error())
	}

	token := os.Getenv("APIKEY")
	req.Header.Set("x_cg_demo_api_key", token)

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %s", err.Error())
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
			return nil, fmt.Errorf("error decoding error response: %s", err.Error())
		}
		return nil, fmt.Errorf("error response: %s", errorResponse["error"].(string))
	}

	var jsonBody map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&jsonBody); err != nil {
		return nil, fmt.Errorf("error decoding response body: %s", err.Error())
	}

	coinWithPrice := model.ConvertJsonToCoinWithPriceModel(jsonBody)

	return coinWithPrice, nil
}

func FetchCompaniesFromId(id string) ([]*model.Company, error) {
	req, err := http.NewRequest("GET", "https://api.coingecko.com/api/v3/companies/public_treasury/"+id, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err.Error())
	}

	token := os.Getenv("APIKEY")
	req.Header.Set("x_cg_demo_api_key", token)

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %s", err.Error())
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
			return nil, fmt.Errorf("error decoding error response: %s", err.Error())
		}
		return nil, fmt.Errorf("error response: %s", errorResponse["error"].(string))
	}

	var jsonResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&jsonResponse); err != nil {
		return nil, fmt.Errorf("error decoding response body: %s", err.Error())
	}

	companiesData, err := json.Marshal(jsonResponse["companies"])
	if err != nil {
		return nil, fmt.Errorf("error marshaling companies data: %s", err.Error())
	}

	var companies []*model.Company
	if err := json.Unmarshal(companiesData, &companies); err != nil {
		return nil, fmt.Errorf("error unmarshaling companies data: %s", err.Error())
	}

	return companies, nil
}
