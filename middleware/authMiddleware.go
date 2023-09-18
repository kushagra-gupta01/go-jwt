package middleware

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/kushagra-gupta01/go-jwt/helpers"
)

func Authenticate()gin.HandlerFunc{
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error":fmt.Sprintf("No authorization header provided")})
			c.Abort()
			return
		}

		claims,err := helpers.ValidateToken(clientToken)
		if err != ""{
			c.JSON(http.StatusInternalServerError,gin.H{"error":err})
			c.Abort()
			return
		}
		c.Set("email",claims.Email)
		c.Set("first_name",claims.First_Name)
		c.Set("last_name",claims.Last_Name)
		c.Set("uid",claims.Uid)
		c.Set("user_type",claims.User_type)
		c.Next()
	}
}

