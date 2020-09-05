package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHanlder())
	defer ts.Close()
	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHanlder())
	defer ts.Close()
	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No Users")
}

func TestUsers_WithUsersData(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHanlder())
	defer ts.Close()

	mockUsers := make([]*User, 0, 10)
	for i := 0; i < 10; i++ {
		user := new(User)
		user.FirstName = "Lee" + strconv.Itoa(i+1)
		user.LastName = "sangming" + strconv.Itoa(i+1)
		user.Email = "mytkdals93@gmail.com" + strconv.Itoa(i+1)
		mockUsers = append(mockUsers, user)
	}
	for _, u := range mockUsers {
		data, _ := json.Marshal(u)
		userStr := fmt.Sprintf(`%v`, string(data))
		resp, err := http.Post(ts.URL+"/users", "application/json",
			strings.NewReader(userStr))
		assert.NoError(err)
		assert.Equal(http.StatusCreated, resp.StatusCode)
	}

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	users := []*User{}
	err = json.NewDecoder(resp.Body).Decode(&users)
	assert.NoError(err)
	assert.Equal(10, len(users))

}
func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHanlder())
	defer ts.Close()
	resp, err := http.Get(ts.URL + "/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User Id: 89")
	resp, err = http.Get(ts.URL + "/users/55")
	data, _ = ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User Id: 55")
}
func TestCreateUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHanlder())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"sangmin","last_name":"Lee","email":"mytkdals93@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	id := user.ID
	resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	user2 := new(User)
	err = json.NewDecoder(resp.Body).Decode(user2)
	assert.NoError(err)
	assert.Equal(user.ID, user2.ID)
}
func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHanlder())
	defer ts.Close()

	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err := http.DefaultClient.Do(req)

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID: 1")
	// log.Print(string(data))

	resp, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"sangmin","last_name":"Lee","email":"mytkdals93@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ = ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "Deleted User ID: 1")
}
func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHanlder())
	defer ts.Close()

	req, _ := http.NewRequest("PUT", ts.URL+"/users",
		strings.NewReader(`{"id":1, "first_name":"updated", "last_name":"Lee", "email":"updated@gmail.com"}`))

	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User ID: 1")

	resp, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"sangmin","last_name":"Lee","email":"mytkdals93@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)
	updateStr := fmt.Sprintf(`
	{
		"id":%d, 
		"first_name":"updated",
		"update_first_name":true,
		"last_name":"",
		"update_last_name":true,
		"email":""
	}
	`,
		user.ID)
	req, _ = http.NewRequest("PUT", ts.URL+"/users",
		strings.NewReader(updateStr))
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	updateUser := new(User)
	err = json.NewDecoder(resp.Body).Decode(updateUser)
	assert.NoError(err)
	assert.Equal(updateUser.ID, user.ID)
	assert.Equal("updated", updateUser.FirstName)
	assert.Equal("", updateUser.LastName)
	assert.Equal(user.Email, updateUser.Email)
}
