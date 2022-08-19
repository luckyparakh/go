package app

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func StartApp() {
	router = gin.Default()
	mapping()
	router.Run(":8880")
}
