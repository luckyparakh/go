package repositories

import (
	"example/domain/repo"
	"example/services"
	"example/utils/errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRepo(c *gin.Context) {
	var request repo.ServiceRepoRequest
	fmt.Println("Createrepo")
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr.Message())
		return
	}
	response, err := services.RepoService.CreateRepo(request)
	if err != nil {
		apiErr := errors.NewBadRequestError(err.Message())
		c.JSON(apiErr.Status(), apiErr.Message())
		return
	}
	c.JSON(http.StatusCreated, response)
}

func CreateRepos(c *gin.Context) {
	var request []repo.ServiceRepoRequest
	fmt.Println("Createrepo")
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr.Message())
		return
	}
	response, err := services.RepoService.CreateRepos(request)
	if err != nil {
		apiErr := errors.NewBadRequestError("invalid request")
		c.JSON(apiErr.Status(), apiErr.Message())
		return
	}
	c.JSON(http.StatusCreated, response)
}
