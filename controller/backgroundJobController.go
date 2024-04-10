package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"

	"github.com/Uttkarsh-raj/Crypto/database"
	"github.com/Uttkarsh-raj/Crypto/helper"
	"github.com/Uttkarsh-raj/Crypto/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartBackgroundJob(client *mongo.Client) {
	c := cron.New()

	_, err := c.AddFunc("0 * * * *", func() {
		fmt.Printf("%s:Background process started.\n", time.Now())
		GetCoinsAndStore(client)
	})
	if err != nil {
		log.Fatal("Error adding cron job:", err)
	}

	c.Start()

	select {}
}

func GetCoinsAndStore(client *mongo.Client) {
	coins := helper.GetAllCoins()
	StoreCoins(client, coins)
}

func StoreCoins(client *mongo.Client, coins []*model.CoinModel) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	collection := database.OpenCollection(client, "CryptoCoins")
	var storeCoins []interface{}
	for i := 0; i < len(coins) && i < 100; i++ {
		var coin *model.CoinModel
		err := collection.FindOne(ctx, bson.M{"name": coins[i].Name, "id": coins[i].Id, "symbol": coins[i].Symbol}).Decode(&coin)
		if err != nil {
			coins[i].ObjectId = primitive.NewObjectID()
			storeCoins = append(storeCoins, coins[i])
		}
	}
	if len(storeCoins) > 0 {
		_, err := collection.InsertMany(ctx, storeCoins)
		if err != nil {
			log.Fatalf("Error pushing data to the server: %s", err)
			return
		}
	}

	fmt.Printf("%s: Successfully updated data to DB.\n", time.Now())
}
