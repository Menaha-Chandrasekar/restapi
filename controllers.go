package controllers//are independent
// controllers are consumed by routes and routes by main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func Login(c *gin.Context){
	c.String(http.StatusOK, "Hello")
}

func Register(c *gin.Context){
     
	c.String(http.StatusOK, "Helloworld")

 }

// func GetProfile(c *gin.Context){

// }