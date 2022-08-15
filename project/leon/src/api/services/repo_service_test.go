package services

import (
	"example/clients/restclient"
	"example/domain/repo"
	"example/utils/errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMocking()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidName(t *testing.T) {
	req := repo.ServiceRepoRequest{
		Name: "",
	}

	resp, err := RepoService.CreateRepo(req)
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Status(), http.StatusBadRequest)
	assert.EqualValues(t, err.Message(), "invalid repo name")
}

func TestCreateRepoErr(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMock(fmt.Sprintf("%s_%s", "https://api.github.com/user/repos", http.MethodPost),
		restclient.Mock{
			Err: nil,
			Response: &http.Response{
				Body:       ioutil.NopCloser(strings.NewReader("{\"message\":1}")),
				StatusCode: http.StatusInternalServerError,
			},
		})
	req := repo.ServiceRepoRequest{
		Name: "abc",
	}
	resp, err := RepoService.CreateRepo(req)
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, err.Status(), http.StatusInternalServerError)
	assert.EqualValues(t, err.Message(), "can't unmarshal body")
}

func TestCreateRepoNoErr(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMock(fmt.Sprintf("%s_%s", "https://api.github.com/user/repos", http.MethodPost),
		restclient.Mock{
			Err: nil,
			Response: &http.Response{
				Body:       ioutil.NopCloser(strings.NewReader(`{"name":"rishi","id":123,"full_name":"Rishi Parakh"}`)),
				StatusCode: http.StatusOK,
			},
		})
	req := repo.ServiceRepoRequest{
		Name: "abc",
	}
	resp, err := RepoService.CreateRepo(req)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.EqualValues(t, resp.Id, 123)
	assert.EqualValues(t, resp.Name, "rishi")
}

func TestCreateRepoConInvalidName(t *testing.T) {
	var req repo.ServiceRepoRequest
	op := make(chan repo.ReposResponse)
	s := repoService{}
	go s.createRepoCon(req, op)
	result := <-op
	assert.NotNil(t, result)
	assert.NotNil(t, result.Error)
	assert.EqualValues(t, http.StatusBadRequest, result.Error.Status())
	assert.EqualValues(t, "invalid repo name", result.Error.Message())
	assert.Nil(t, result.Response)
}

func TestCreateRepoConErr(t *testing.T) {
	var req repo.ServiceRepoRequest
	req.Name = "abc"
	op := make(chan repo.ReposResponse)
	s := repoService{}
	restclient.FlushMocks()
	restclient.AddMock(fmt.Sprintf("%s_%s", "https://api.github.com/user/repos", http.MethodPost),
		restclient.Mock{
			Err: nil,
			Response: &http.Response{
				Body:       ioutil.NopCloser(strings.NewReader(`{"message":1}`)),
				StatusCode: http.StatusInternalServerError,
			},
		})
	go s.createRepoCon(req, op)
	result := <-op
	assert.NotNil(t, result)
	assert.NotNil(t, result.Error)
	assert.Nil(t, result.Response)
	assert.EqualValues(t, http.StatusInternalServerError, result.Error.Status())
	assert.EqualValues(t, "can't unmarshal body", result.Error.Message())
}

func TestCreateRepoConNoErr(t *testing.T) {
	var req repo.ServiceRepoRequest
	req.Name = "abc"
	op := make(chan repo.ReposResponse)
	s := repoService{}
	restclient.FlushMocks()
	restclient.AddMock(fmt.Sprintf("%s_%s", "https://api.github.com/user/repos", http.MethodPost),
		restclient.Mock{
			Err: nil,
			Response: &http.Response{
				Body:       ioutil.NopCloser(strings.NewReader(`{"name":"rishi","id":123}`)),
				StatusCode: http.StatusOK,
			},
		})
	go s.createRepoCon(req, op)
	result := <-op
	assert.NotNil(t, result)
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Response)
	assert.EqualValues(t, 123, result.Response.Id)
	assert.EqualValues(t, "rishi", result.Response.Name)
}

func TestHandleResult(t *testing.T) {
	ip := make(chan repo.ReposResponse)

	op := make(chan repo.ServiceReposResponse)
	s := repoService{}

	go s.handleResult(ip, op)
	go func() {
		defer close(ip)
		ip <- repo.ReposResponse{
			Error: errors.NewBadRequestError("abc"),
			Response: &repo.ServiceRepoResponse{
				Id: 1,
			},
		}
	}()
	result := <-op
	assert.NotNil(t, result)
	assert.EqualValues(t, result.Result[0].Response.Id, 1)
	assert.EqualValues(t, result.Result[0].Error.Message(), "abc")
}

func TestCreateReposInvalidReq(t *testing.T) {
	req := []repo.ServiceRepoRequest{
		{},
		{Name: ""},
	}

	resp, err := RepoService.CreateRepos(req)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}
