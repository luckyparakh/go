package restclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var enableMocking bool
var mocks = make(map[string]Mock)

func StartMocking() {
	enableMocking = true
}
func StopMocking() {
	enableMocking = false
}

type Mock struct {
	Response *http.Response
	Err      error
	Token    string
	Method   string
}

func AddMock(key string, m Mock) {
	mocks[key] = m
}
func FlushMocks() {
	mocks = make(map[string]Mock)
}

func getKey(url string, method string) string {
	return fmt.Sprintf("%s_%s", url, method)
}

func Post(url string, body any, header http.Header) (*http.Response, error) {
	if enableMocking {
		m := mocks[getKey(url, http.MethodPost)]
		return m.Response, m.Err
	}
	jsonByte, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonByte))
	if err != nil {
		return nil, err
	}
	request.Header = header
	client := http.Client{}
	return client.Do(request)
}
