package services

import (
	"example/clients/restclient"
	"example/domain/repo"
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
