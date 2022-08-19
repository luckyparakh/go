package repositories

import (
	"encoding/json"
	"example/domain/repo"
	"example/services"
	"example/utils/errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type repoServiceMock struct{}

var (
	funcCreateRepo  func(req repo.ServiceRepoRequest) (*repo.ServiceRepoResponse, errors.ApiError)
	funcCreateRepos func(req []repo.ServiceRepoRequest) (repo.ServiceReposResponse, errors.ApiError)
)

func (rs *repoServiceMock) CreateRepo(req repo.ServiceRepoRequest) (*repo.ServiceRepoResponse, errors.ApiError) {
	return funcCreateRepo(req)
}

func (rs *repoServiceMock) CreateRepos(req []repo.ServiceRepoRequest) (repo.ServiceReposResponse, errors.ApiError) {
	return funcCreateRepos(req)
}

func TestCreateRepoErrorMockEntireService(t *testing.T) {
	services.RepoService = &repoServiceMock{}
	funcCreateRepo = func(req repo.ServiceRepoRequest) (*repo.ServiceRepoResponse, errors.ApiError) {
		return nil, errors.NewBadRequestError("error")
	}
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, _ := http.NewRequest(http.MethodPost, "/repo", strings.NewReader(`{"name":"rp"}`))
	c.Request = req
	CreateRepo(c)
	
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	assert.EqualValues(t, response.Body.String(), "\"error\"")
}

func TestCreateRepoNoErrorMockEntireService(t *testing.T) {
	services.RepoService = &repoServiceMock{}
	funcCreateRepo = func(req repo.ServiceRepoRequest) (*repo.ServiceRepoResponse, errors.ApiError) {
		return &repo.ServiceRepoResponse{
			Name:  "mocked service",
			Id:    123,
			Owner: "Go",
		}, nil
	}
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, _ := http.NewRequest(http.MethodPost, "/repo", strings.NewReader(`{"name":"rp"}`))
	c.Request = req
	CreateRepo(c)

	assert.EqualValues(t, http.StatusCreated, response.Code)
	var resp repo.ServiceRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.EqualValues(t, resp.Id, 123)
}
