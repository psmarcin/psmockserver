package main

import (
	"net/http"
	"testing"
)

func TestReadinesEndpoint(t *testing.T) {
	BeforeEach()

	url := baseURL + "health/r"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("Expect http status to be %d but got %d", http.StatusOK, res.StatusCode)
	}

	AfterEach()
}

func TestLivenessEndpoint(t *testing.T) {
	BeforeEach()

	url := baseURL + "health/l"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("Expect http status to be %d but got %d", http.StatusOK, res.StatusCode)
	}

	AfterEach()
}
