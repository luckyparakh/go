package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func MatchUserTypeToUid(c *gin.Context, userId string) error {
	userType := c.GetString("user_type")
	uid := c.GetString("user_id")
	if userType == "USER" && uid != userId {
		err := errors.New("Unauthorized access")
		return err
	}
	err := checkUserType(c, userType)
	return err
}

func CheckUserType(c *gin.Context, role string) error {
	userType := c.GetString("user_type")
	if userType != role {
		return errors.New("Unauthorized access")
	}
	return nil
}
