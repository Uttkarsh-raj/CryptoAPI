package controller

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Uttkarsh-raj/Crypto/database"
	"github.com/Uttkarsh-raj/Crypto/helper"
	"github.com/Uttkarsh-raj/Crypto/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartBackgroundJob(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*60)
	defer cancel()
	fmt.Printf("%s: Background process started.\n", time.Now())
	Schedule(ctx, time.Second*60, time.Second*60, func(t time.Time) { // change the period and duration for running the data
		GetCoinsAndStore(client)
	})
}

func Schedule(ctx context.Context, p time.Duration, o time.Duration, execFunc func(time.Time)) {
	first := time.Now().Truncate(p).Add(o)
	if first.Before(time.Now()) {
		first = first.Add(p)
	}
	firstTime := time.After(time.Until(first))

	ticker := &time.Ticker{C: nil}

	for {
		select {
		case v := <-firstTime:
			ticker = time.NewTicker(p)
			execFunc(v)
		case v := <-ticker.C:
			execFunc(v)
		case <-ctx.Done():
			ticker.Stop()
			return
		}
	}
}

func GetCoinsAndStore(client *mongo.Client) {
	coins := helper.GetAllCoins()
	StoreCoins(client, coins)
}

func StoreCoins(client *mongo.Client, coins []*model.CoinModel) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	collection := database.OpenCollection(client, "CryptoCoins")
	for i := 0; i < len(coins) && i < 2000; i++ {
		var coin *model.CoinModel
		err := collection.FindOne(ctx, bson.M{"name": coins[i].Name, "id": coins[i].Id, "symbol": coins[i].Symbol}).Decode(&coin)
		if err != nil {
			coins[i].ObjectId = primitive.NewObjectID()
			_, err := collection.InsertOne(ctx, coins[i])
			if err != nil {
				log.Fatalf("Error pushing data to the server: %s", err)
				return
			}
		}
	}

	fmt.Printf("%s: Successfully updated data to DB.\n", time.Now())
}
