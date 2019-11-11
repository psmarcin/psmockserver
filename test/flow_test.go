package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestFlowMockResetMock(t *testing.T) {
	BeforeEach()

	// mock
	url := baseURL + "mockserver"
	mockBody := "{'test': true}"

	payload := strings.NewReader("{\n    \"httpRequest\": {\n        \"path\": \"/\",\n        \"method\": \"GET\",\n        \"content-type\": \"application/json\"\n    },\n    \"httpResponse\": {\n        \"body\": \"" + mockBody + "\",\n        \"statusCode\": 200,\n        \"headers\": {\n            \"test\": \"true\"\n        }\n    },\n    \"times\": {\n        \"remainingTimes\": 1,\n        \"unlimited\": false\n    }\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if res.StatusCode != 200 {
		t.Fatalf("Can't get list of all mocked paths")
	}

	// test mock
	url = baseURL
	req, _ = http.NewRequest("GET", url, nil)
	res, _ = http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if string(body) != mockBody {
		t.Fatalf("expect %s to equal %s", body, mockBody)
	}

	// reset all mocks
	url = baseURL + "mockserver/reset"
	req, _ = http.NewRequest("PUT", url, nil)
	res, _ = http.DefaultClient.Do(req)

	if res.StatusCode != http.StatusAccepted {
		t.Fatalf("expect reset status code %d to equal %d", res.StatusCode, http.StatusAccepted)
	}

	url = baseURL
	req, _ = http.NewRequest("GET", url, nil)
	res, _ = http.DefaultClient.Do(req)

	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("expect %d for not mocked endpoint but got %d", http.StatusNotFound, res.StatusCode)
	}

	AfterEach()
}
