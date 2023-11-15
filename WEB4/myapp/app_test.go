package myapp

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)
	assert.Equal("hello world", string(data))
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())

	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)
	assert.Contains(string(data), "get userinfo")

}

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())

	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)
	assert.Contains(string(data), "User Id:89")
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())

	defer ts.Close()

	resp, err := http.Post(us.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"tucker", "last_name":"kim", "email":"tucker@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
}
