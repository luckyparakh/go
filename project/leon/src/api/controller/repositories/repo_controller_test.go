package repositories

import (
	"encoding/json"
	"example/clients/restclient"
	"example/domain/repo"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMocking()
	os.Exit(m.Run())
}
func TestCreateInvalidJsonRequest(t *testing.T) {
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, _ := http.NewRequest(http.MethodPost, "/repo", strings.NewReader(""))
	c.Request = req
	CreateRepo(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	assert.EqualValues(t, response.Body.String(), "\"invalid json body\"")
}

func TestCreateInvalidGitResponse(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMock(fmt.Sprintf("%s_%s", "https://api.github.com/user/repos", http.MethodPost),
		restclient.Mock{
			Err: nil,
			Response: &http.Response{
				Body:       ioutil.NopCloser(strings.NewReader(`{"id":"123"}`)), // id as string
				StatusCode: http.StatusOK,
			},
		})
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, _ := http.NewRequest(http.MethodPost, "/repo", strings.NewReader(`{"name":"rp"}`))
	c.Request = req
	CreateRepo(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	assert.EqualValues(t, response.Body.String(), "\"error while parsing result\"")
}

func TestCreateValidGitResponse(t *testing.T) {
	restclient.FlushMocks()
	restclient.AddMock(fmt.Sprintf("%s_%s", "https://api.github.com/user/repos", http.MethodPost),
		restclient.Mock{
			Err: nil,
			Response: &http.Response{
				Body:       ioutil.NopCloser(strings.NewReader(`{"id":123}`)),
				StatusCode: http.StatusOK,
			},
		})
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
