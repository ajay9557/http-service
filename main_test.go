package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {

	req, _ := http.NewRequest(method, path, nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}

func TestShowAllUsers(t *testing.T) {
	body := map[string][]map[string]interface{}{
		"data": {
			{"id": float64(1), "name": "ajay", "email": "ajay@gmail.com"},
			{"id": float64(2), "name": "abhi", "email": "abhi@gmail.com"},
		},
	}

	server := Setup()

	w := performRequest(server, "GET", "/users")

	assert.Equal(t, http.StatusOK, w.Code)

	//Convert the JSON response to a map
	var response map[string][]map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["data"]

	fmt.Println("value--", value)
	fmt.Println("exits--", exists)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["data"], response["data"], "not equal")

}

func TestGetSingleUser(t *testing.T) {

	body := map[string]map[string]interface{}{
		"data": {"id": float64(2), "name": "abhi", "email": "abhi@gmail.com"},
	}

	server := Setup()

	w := performRequest(server, "GET", "/users/2")

	assert.Equal(t, http.StatusOK, w.Code)

	//Convert the JSON response to a map
	var response map[string]map[string]interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	value, exists := response["data"]

	fmt.Println(value)
	fmt.Println(exists)

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["data"], response["data"], "not equal")

}
