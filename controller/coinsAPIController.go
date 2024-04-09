package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Uttkarsh-raj/Crypto/helper"
	"github.com/Uttkarsh-raj/Crypto/model"
	"github.com/gin-gonic/gin"
)

func TestServer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "Hello there"})
	}
}

func ConvertPrice() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), time.Second*100)
		defer cancel()

		var body map[string]string
		err := json.NewDecoder(c.Request.Body).Decode(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			return
		}

		request := model.ConvertFromJson(body)

		fromCoin, err := helper.GetCoinByIdAndDate(request.FromCurrency, request.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			return
		}

		toCoin, err := helper.GetCoinByIdAndDate(request.ToCurrency, request.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			return
		}

		if toCoin != nil && fromCoin != nil {

			priceOfOneUsdToCoin := 1 / toCoin.MarketData.CurrentPrice["usd"].(float64)
			priceInTermsOfToCoin := priceOfOneUsdToCoin * fromCoin.MarketData.CurrentPrice["usd"].(float64)

			response := gin.H{
				"success": true,
				"data":    strconv.FormatFloat(priceInTermsOfToCoin, 'f', -1, 64) + " " + toCoin.Name,
			}

			c.JSON(http.StatusOK, response)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Some error occured."})
	}
}

func FetchAllCompanies() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, cancel := context.WithTimeout(context.Background(), time.Second*100)
		defer cancel()

		var body map[string]interface{}
		err := json.NewDecoder(c.Request.Body).Decode(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			return
		}

		coinId := body["currency"].(string)
		if coinId == "" {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "currency field missing from body."})
			return
		}
		listOfCompanies, err := helper.FetchComapaniesFromId(coinId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
			return
		}

		response := gin.H{
			"success": true,
			"data":    listOfCompanies,
		}

		c.JSON(http.StatusOK, response)
	}
}
