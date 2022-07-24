package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := UserDao.GetUser(0)
	assert.Nil(t, user, "not expecting a user")
	// if user != nil {
	// 	t.Error("not expecting a user")
	// }
	assert.NotNil(t, err, "expecting error")
	// if err == nil {
	// 	t.Error("expecting error")
	// }
	assert.Equal(t, err.StatusCode, http.StatusNotFound)
	// if err.StatusCode != http.StatusNotFound {
	// 	t.Error("expecting 404 error code")
	// }
	assert.Equal(t, err.Message, "user 0 not found")
	assert.Equal(t, err.Code, "not found")
}
func TestGetUserFound(t *testing.T) {
	//Initialization
	
	//Execution
	user, err := UserDao.GetUser(123)

	//Validation
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t,user.Id,123)
}
