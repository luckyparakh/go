package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstant(t *testing.T) {
	assert.EqualValues(t, apiGithubAccessToken, "SECRET_GITHUB_ACCESS_TOKEN")
}

func TestGetGitHubAccessToken(t *testing.T) {
	assert.EqualValues(t, GetGitHubAccessToken(), "")
}
