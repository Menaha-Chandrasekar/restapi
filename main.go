package main

import (
	"RESTAPI/constants"
	"RESTAPI/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default()
	routes.AppRoutes(router)
	log.Fatal(router.Run(constants.Port))
}