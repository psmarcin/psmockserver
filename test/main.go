package main

import "net/http"

// Reset cleanup all mocked endpoints
func Reset() {
	url := baseURL + "mockserver/reset"
	req, _ := http.NewRequest("PUT", url, nil)
	http.DefaultClient.Do(req)
}

// BeforeEach runs before every single test case
func BeforeEach() {
	Reset()
}

// AfterEach runs after every single test case
func AfterEach() {

}
