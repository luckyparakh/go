package app

import "example/controller/repositories"

func routes() {
	router.POST("/repo", repositories.CreateRepo)
}
