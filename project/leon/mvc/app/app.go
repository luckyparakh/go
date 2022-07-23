package app

import (
	"mvc/controller"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controller.GetUser)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		panic(err)
	}
}
