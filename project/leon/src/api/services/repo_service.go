package services

import (
	"example/config"
	"example/domain/github"
	"example/domain/repo"
	"example/provider/github_provider"
	"example/utils/errors"
	"net/http"
)

type repoService struct{}
type repoServiceInterface interface {
	CreateRepo(req repo.ServiceRepoRequest) (*repo.ServiceRepoResponse, errors.ApiError)
	CreateRepos(req []repo.ServiceRepoRequest) (repo.ServiceReposResponse, errors.ApiError)
}

var (
	RepoService repoServiceInterface
)

func init() {
	RepoService = &repoService{}
}

func (s *repoService) CreateRepo(req repo.ServiceRepoRequest) (*repo.ServiceRepoResponse, errors.ApiError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	request := github.CreateRepoRequest{
		Name:        req.Name,
		Description: req.Description,
		Private:     false,
	}
	response, err := github_provider.CreateRepo(config.GetGitHubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repo.ServiceRepoResponse{
		Name:  response.Name,
		Owner: response.Owner.Login,
		Id:    response.Id,
	}

	return &result, nil
}

func (s *repoService) CreateRepos(req []repo.ServiceRepoRequest) (repo.ServiceReposResponse, errors.ApiError) {
	ip := make(chan repo.ReposResponse)
	op := make(chan repo.ServiceReposResponse)
	defer close(op)
	go s.handleResult(ip, op)
	for _, r := range req {
		go s.createRepoCon(r, ip)
	}
	result := <-op
	successCount := 0
	for _, current := range result.Result {
		if current.Response != nil {
			successCount++
		}
	}
	if successCount == 0 {
		result.StatusCode = result.Result[0].Error.Status()
	} else if successCount == len(req) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}
	return result, nil
}

func (s *repoService) handleResult(ip chan repo.ReposResponse, op chan repo.ServiceReposResponse) {
	var results repo.ServiceReposResponse
	defer close(op)
	for resp := range ip {
		results.Result = append(results.Result, resp)
	}
	op <- results
}

func (s *repoService) createRepoCon(r repo.ServiceRepoRequest, c chan repo.ReposResponse) {

	if err := r.Validate(); err != nil {
		c <- repo.ReposResponse{
			Error: err,
		}
		return
	}
	response, err := s.CreateRepo(r)
	if err != nil {
		c <- repo.ReposResponse{
			Error: err,
		}
		return
	}
	c <- repo.ReposResponse{
		Response: response,
	}
}
