package controller

import (
	"mvc/service"
	"mvc/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	userIdParam, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if err != nil {
		userErr := &utils.AppError{
			Message:    "user should be number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondErr(ctx, userErr)
		return
	}
	user, appErr := service.UserService.GetUser(userIdParam)
	if appErr != nil {
		utils.RespondErr(ctx, appErr)
		return
	}
	utils.Respond(ctx, http.StatusOK, user)
}
