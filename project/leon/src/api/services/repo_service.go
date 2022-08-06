package services

import (
	"example/config"
	"example/domain/github"
	"example/domain/repo"
	"example/provider/github_provider"
	"example/utils/errors"
	"strings"
)

type repoService struct{}
type repoServiceInterface interface {
	CreateRepo(req repo.ServiceRepoRequest) (*repo.ServiceRepoResponse, errors.ApiError)
}

var (
	RepoService repoServiceInterface
)

func init() {
	RepoService = &repoService{}
}

func (s *repoService) CreateRepo(req repo.ServiceRepoRequest) (*repo.ServiceRepoResponse, errors.ApiError) {
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return nil, errors.NewBadRequestError("invalid repo name")
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
