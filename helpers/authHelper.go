package helpers

import(
	"errors"
	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context,role string)(err error){
	userType := c.GetString("user_type")
	err=nil

	if userType !=role{
		err = errors.New("Unauthorized to access this resource")
		return err
	}
	return err
}

func MatchUserTypeToUid(ctx gin.Context, userId string)(err error){
	userType := ctx.GetString("user_type")
	uid := ctx.GetString("uid")
	err = nil

	if userType == "USER" && uid !=userId{
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	err = CheckUserType(ctx,userType)
	return err
}