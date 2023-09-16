package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kushagra-gupta01/go-jwt/controllers"
)

func AuthRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup",controllers.SignUp())
	incomingRoutes.POST("/users/login",controllers.Login())
}