package main

import (
	"jwt/controller"
	"jwt/database"
	"jwt/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect("root:root@123@tcp(localhost:3306)/jwt_demo?parseTime=true")
	database.Migrate()
	r := initRouter()
	r.Run(":8000")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controller.GenerateToken)
		api.POST("/user/register", controller.RegisterUser)
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/ping", controller.Ping)
		}
	}
	return router
}
