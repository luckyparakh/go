package app

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func StartApp() {
	router = gin.Default()
	routes()
	if err := router.Run(":8090"); err != nil {
		panic(err)
	}
}
