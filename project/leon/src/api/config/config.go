package config

import "os"

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
	LogLevel             = "info"
)

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

func GetGitHubAccessToken() string {
	return githubAccessToken
}
