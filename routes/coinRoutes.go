package routes

import (
	"github.com/Uttkarsh-raj/Crypto/controller"
	"github.com/gin-gonic/gin"
)

func IntegrateRoutes(router *gin.Engine) {
	router.GET("/convert", controller.ConvertPrice())
	router.GET("/companies", controller.FetchAllCompanies())
}
