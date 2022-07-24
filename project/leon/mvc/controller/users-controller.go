package controller

import (
	"encoding/json"
	"mvc/service"
	"mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userIdParam, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		userErr:=&utils.AppError{
			Message: "user should be number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		jsonVal,_:=json.Marshal(userErr)
		w.WriteHeader(userErr.StatusCode)
		w.Write(jsonVal)
		return
	}
	user, appErr := service.UserService.GetUser(userIdParam)
	if appErr != nil {
		w.WriteHeader(appErr.StatusCode)
		w.Write([]byte(appErr.Message))
		return
	}

	jsonValue, _ := json.Marshal(user)
	w.Write(jsonValue)
}
