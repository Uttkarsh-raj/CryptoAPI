package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Uttkarsh-raj/Crypto/helper"
	"github.com/Uttkarsh-raj/Crypto/model"
	"github.com/gin-gonic/gin"
)

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
		println(request.FromCurrency)

		fromCoin, err := helper.GetCoinByIdAndDate(request.FromCurrency, request.Date)
		if err != nil {
			// fmt.Println(err.Error())
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
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Some error occured."})
	}
}

func FetchAllCompanies() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
		defer cancel()

		var body map[string]interface{}
		err := json.NewDecoder(c.Request.Body).Decode(&body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
			return
		}

		listOfCompanies := helper.FetchComapaniesFromId(body["currency"])
		fmt.Println(ctx)
	}
}
