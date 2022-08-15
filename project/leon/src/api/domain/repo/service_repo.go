package repo

import (
	"example/utils/errors"
	"strings"
)

type ServiceRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *ServiceRepoRequest) Validate() errors.ApiError {
	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		return errors.NewBadRequestError("invalid repo name")
	}
	return nil
}

type ServiceRepoResponse struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type ServiceReposResponse struct {
	StatusCode int             `json:"status"`
	Result     []ReposResponse `json:"results"`
}

type ReposResponse struct {
	Response *ServiceRepoResponse `json:"response"`
	Error    errors.ApiError      `json:"error"`
}
