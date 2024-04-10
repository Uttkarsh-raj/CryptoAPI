package main

import (
	"log"

	// "github.com/Uttkarsh-raj/Crypto/controller" //imports to run the background job
	// "github.com/Uttkarsh-raj/Crypto/database"
	"github.com/Uttkarsh-raj/Crypto/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	routes.IntegrateRoutes(router)
	// mongoClient := database.DBInstance()          // uncomment this to run connecct to the mongo server
	// go controller.StartBackgroundJob(mongoClient) // To run the background task uncomment this
	log.Fatal(router.Run(":3000"))
}
