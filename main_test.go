package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

func init() {
	gin.SetMode(gin.TestMode)
	r.GET("/ping", ping)
}

func TestPing(t *testing.T) {
	type testCase struct {
		method               string
		path                 string
		expectedResponseCode int
		expectedBody         string
	}

	cases := []testCase{
		{"GET", "/ping", http.StatusOK, "{\"message\":\"Welcome to the Coffeeshop!\"}"},
	}

	for _, tc := range cases {
		req := getRequest(tc.method, tc.path)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		body, _ := ioutil.ReadAll(w.Body)
		if tc.expectedResponseCode != w.Code {
			t.Errorf("Expected '%v', but got '%v'", tc.expectedResponseCode, w.Code)
		}
		if tc.expectedBody != string(body) {
			t.Errorf("Expected '%v', but got '%v'", tc.expectedBody, string(body))
		}
	}
}

func getRequest(method, path string) *http.Request {
	req, _ := http.NewRequest(method, path, nil)
	return req
}

// package main
//
// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
//
// 	"github.com/gin-gonic/gin"
// )
//
// func TestAddTwoNumbers(t *testing.T) {
// 	type TestCase struct {
// 		Input1   int
// 		Input2   int
// 		Expected int
// 	}
//
// 	cases := []TestCase{
// 		{1, 2, 3},
// 		{2, 3, 5},
// 		{3, 4, 7},
// 		{4, 5, 9},
// 	}
//
// 	for _, tc := range cases {
// 		fmt.Printf("%#v", tc)
// 		got := AddTwoNumbers(tc.Input1, tc.Input2)
// 		if tc.Expected != got {
// 			t.Errorf("Expected %d, but got %d", tc.Expected, got)
// 		}
// 	}
// }
//
// var r = gin.Default()
//
// func init() {
// 	gin.SetMode(gin.TestMode)
// 	r.GET("/ping", ping)
// }
//
// func TestPing(t *testing.T) {
// 	type testCase struct {
// 		method               string
// 		path                 string
// 		expectedResponseCode int
// 		expectedBody         string
// 	}
//
// 	cases := []testCase{
// 		{"GET", "/ping", http.StatusOK, "{\"message\":\"pong\"}"},
// 	}
//
// 	for _, tc := range cases {
// 		req := getRequest(tc.method, tc.path)
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)
//
// 		body, _ := ioutil.ReadAll(w.Body)
// 		if tc.expectedResponseCode != w.Code {
// 			t.Errorf("Expected '%v', but got '%v'", tc.expectedResponseCode, w.Code)
// 		}
// 		if tc.expectedBody != string(body) {
// 			t.Errorf("Expected '%v', but got '%v'", tc.expectedBody, string(body))
// 		}
// 	}
// }
//
// func getRequest(method, path string) *http.Request {
// 	req, _ := http.NewRequest("GET", "/ping", nil)
// 	return req
// }
