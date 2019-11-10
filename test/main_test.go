package main

import (
	"net/http"
	"testing"
)

const (
	baseURL = "http://localhost:8080/"
)

func TestListOfAllMocks(t *testing.T) {
	url := baseURL + "mockserver"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("Can't get list of all mocked paths")
	}
}

func TestShouldGet404OnNonMockedEndpoint(t *testing.T) {
	url := baseURL + ""

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if res.StatusCode != 404 {
		t.Fatalf("Get %d instead of 404 status code", res.StatusCode)
	}
}
