package routes

import (
	"RESTAPI/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine){
	router.GET("/api/login",controllers.Login)

	router.GET("/api/register",controllers.Register)
	//router.GET("/api/profile:id")
	
}