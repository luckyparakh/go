package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "golang intro",
		Description: "a golang repo",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}
	var target CreateRepoRequest
	byte, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, byte)
	err = json.Unmarshal(byte, &target)
	assert.Nil(t, err)
	assert.Equal(t, target.Name, request.Name)
}
