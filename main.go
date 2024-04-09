package main

import (
	"log"

	"github.com/Uttkarsh-raj/Crypto/controller"
	"github.com/Uttkarsh-raj/Crypto/database"
	"github.com/Uttkarsh-raj/Crypto/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	routes.IntegrateRoutes(router)
	mongoClient := database.DBInstance()
	go controller.StartBackgroundJob(mongoClient)
	log.Fatal(router.Run(":3000"))
}
