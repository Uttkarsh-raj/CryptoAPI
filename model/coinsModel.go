package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CoinModel struct {
	ObjectId primitive.ObjectID `bson:"_id"`
	Id       string             `json:"id"`
	Name     string             `json:"name"`
	Symbol   string             `json:"symbol"`
}

func GetCoinFromJson(json map[string]string) *CoinModel {
	return &CoinModel{
		Id:     json["id"],
		Name:   json["name"],
		Symbol: json["symbol"],
	}
}

func GetCoinsList(json []map[string]string) []*CoinModel {
	var coins []*CoinModel
	for i := 0; i < len(json); i++ {
		coin := GetCoinFromJson(json[i])
		coins = append(coins, coin)
	}
	return coins
}
