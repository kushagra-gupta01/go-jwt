package main

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kushagra-gupta01/go-jwt/middleware"
	"github.com/kushagra-gupta01/go-jwt/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err !=nil{
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port ==""{
		port="8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRouter(router)
	routes.UserRouter(router)
	router.Use(middleware.Authenticate())
	router.GET("/api-1",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{"success":"Access granted for api-1"})
	})

	router.GET("/api-2",func(ctx *gin.Context){
		ctx.JSON(200,gin.H{"success":"Access granted for api-2"})
	})

	router.Run(":"+port)
}