package github_provider

import (
	"errors"
	"example/clients/restclient"
	"example/domain/github"
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
func TestGetAuth(t *testing.T) {
	auth := getAuth("abc123")
	assert.EqualValues(t, "token abc123", auth)

}

func TestCreateRepo(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMock(fmt.Sprintf("%s_%s", "https://api.github.com/user/repos", http.MethodPost),
		restclient.Mock{
			Err:      errors.New("error while creating new github repo"),
			Response: nil,
		})
	resp, err := CreateRepo("abc123", github.CreateRepoRequest{})
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, "error while creating new github repo", err.Message)
}

func TestCreateRepoBodyErr(t *testing.T) {
	restclient.FlushMocks()
	invalid, _ := os.Open("asfe")
	restclient.AddMock(fmt.Sprintf("%s_%s", "https://api.github.com/user/repos", http.MethodPost),
		restclient.Mock{
			Err: nil,
			Response: &http.Response{
				Body: invalid,
			},
		})
	resp, err := CreateRepo("abc123", github.CreateRepoRequest{})
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, "error while reading body", err.Message)
}

func TestCreateRepoStatusCodeErr(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMock(fmt.Sprintf("%s_%s", "https://api.github.com/user/repos", http.MethodPost),
		restclient.Mock{
			Err: nil,
			Response: &http.Response{
				Body:       ioutil.NopCloser(strings.NewReader("{\"message\":1}")),
				StatusCode: http.StatusInternalServerError,
			},
		})
	resp, err := CreateRepo("abc123", github.CreateRepoRequest{})
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, "can't unmarshal body", err.Message)
}

func TestCreateRepoStatusCodeNok(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMock(fmt.Sprintf("%s_%s", "https://api.github.com/user/repos", http.MethodPost),
		restclient.Mock{
			Err: nil,
			Response: &http.Response{
				Body:       ioutil.NopCloser(strings.NewReader(`{"id":"123"}`)), // id as string
				StatusCode: http.StatusOK,
			},
		})
	resp, err := CreateRepo("abc123", github.CreateRepoRequest{})
	assert.Nil(t, resp)
	assert.NotNil(t, err)
	assert.EqualValues(t, "error while parsing result", err.Message)
}

func TestCreateRepoStatusCodeOk(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMock(fmt.Sprintf("%s_%s", "https://api.github.com/user/repos", http.MethodPost),
		restclient.Mock{
			Err: nil,
			Response: &http.Response{
				Body:       ioutil.NopCloser(strings.NewReader(`{"name":"rishi","id":123,"full_name":"Rishi Parakh"}`)),
				StatusCode: http.StatusOK,
			},
		})
	resp, err := CreateRepo("abc123", github.CreateRepoRequest{})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.EqualValues(t, resp.Id, 123)
	assert.EqualValues(t, resp.Name, "rishi")
}
