package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kushagra-gupta01/go-jwt/middleware"
	"github.com/kushagra-gupta01/go-jwt/controllers"
)

func UserRouter(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users",controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id",controllers.GetUser())
}