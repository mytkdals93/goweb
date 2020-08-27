package myapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mux = NewHttpHanlder()

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data))
}
func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello Bar", string(data))
}

func TestBarPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=SANGMIN", nil)
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello SANGMIN", string(data))
}
func TestFooHandler_withoutJson(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusBadRequest, res.Code)
}
func TestFooHandler_Json(t *testing.T) {
	assert := assert.New(t)
	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo",
		strings.NewReader(`{"first_name":"sangmin","last_name":"Lee","email":"mytkdals93@gmail.com"}`))
	mux.ServeHTTP(res, req)
	assert.Equal(http.StatusCreated, res.Code)
	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("sangmin", user.FirstName)
	assert.Equal("Lee", user.LastName)
	assert.Equal("mytkdals93@gmail.com", user.Email)
}
