package app

import (
	"mvc/controller"
)

func routes() {
	// localhost:8090/users/123
	router.GET("/users/:user_id", controller.GetUser)
}
