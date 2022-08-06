package github_provider

import (
	"encoding/json"
	"example/clients/restclient"
	"example/domain/github"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization = "Authorization"
	Url                 = "https://api.github.com/user/repos"
)

func getAuth(accessToken string) string {
	return fmt.Sprintf("token %s", accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GitHubErrorResponse) {
	//header:=getAuth(accessToken)
	header := http.Header{}
	header.Set(headerAuthorization, getAuth(accessToken))
	response, err := restclient.Post(Url, request, header)
	// fmt.Println(response)
	// fmt.Println(err)
	if err != nil {
		log.Printf("error while creating new github repo %s\n", err.Error())
		return nil, &github.GitHubErrorResponse{
			Message:    "error while creating new github repo",
			StatusCode: http.StatusInternalServerError,
		}
	}
	bodyJson, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		log.Printf("error while reading body %s\n", err.Error())
		return nil, &github.GitHubErrorResponse{
			Message:    "error while reading body",
			StatusCode: http.StatusInternalServerError,
		}
	}
	if response.StatusCode > 299 {
		var errResponse github.GitHubErrorResponse
		if err := json.Unmarshal(bodyJson, &errResponse); err != nil {
			log.Printf("can't unmarshal body%s\n", err.Error())
			return nil, &github.GitHubErrorResponse{
				Message:    "can't unmarshal body",
				StatusCode: http.StatusInternalServerError,
			}
		}
		return nil, &errResponse
	}
	var result github.CreateRepoResponse
	if err := json.Unmarshal(bodyJson, &result); err != nil {
		log.Printf("error while parsing result %s\n", err.Error())
		return nil, &github.GitHubErrorResponse{
			Message:    "error while parsing result",
			StatusCode: http.StatusInternalServerError,
		}
	}
	return &result, nil
}
