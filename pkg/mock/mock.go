package mock

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httputil"
	"psmockserver/pkg/config"
	"reflect"

	"github.com/kataras/golog"
)

type Mock struct {
	Body           string      `json:"body"`
	Headers        http.Header `json:"headers"`
	Method         string      `json:"method"`
	StatusCode     int         `json:"statusCode"`
	ContentType    string      `json:"contentType"`
	RemainingTimes Remaining   `json:"remainingTimes"`
}

type MockWithRequest struct {
	Mock    `json:"mock"`
	Request string `json:"request"`
}

type Remaining struct {
	Times     int  `json:"times"`
	Unlimited bool `json:"unlimited"`
}

// Mocks holds all mocks
var Mocks = make(map[*http.Request]Mock)

// Add should add new mock to collection
func Add(id *http.Request, mock Mock) error {
	Mocks[id] = mock
	return nil
}

// Find looks for mock by id in mock collection
func Find(id *http.Request) (Mock, error) {
	var foundKey *http.Request
	found := false
	mock := Mock{}
	for key, value := range Mocks {
		if reflect.DeepEqual(&id, &key) {
			golog.Infof("Found %s %s", key.RequestURI, key.Method)
			found = true
			mock = value
			foundKey = key
		}
	}
	if found != true {
		return mock, errors.New("Not found")
	}

	if mock.RemainingTimes.Unlimited {
		return mock, nil
	}

	if mock.RemainingTimes.Times > 0 {
		mock.RemainingTimes.Times = mock.RemainingTimes.Times - 1
		Mocks[foundKey] = mock
		return mock, nil
	}

	delete(Mocks, foundKey)

	return Mock{}, errors.New("Not found")
}

// List logs all mocked endpoints
func List() {
	golog.Infof("Mocks list: ")
	for x := range Mocks {
		golog.Infof("\t - %s %s %s", x.Method, x.URL.String(), x.URL.Query().Encode())
	}
}

// Serialize marshals json
func Serialize() ([]byte, error) {
	var collection = make([]MockWithRequest, len(Mocks))
	for request, mock := range Mocks {
		url, _ := httputil.DumpRequest(request, true)
		collection = append(collection, MockWithRequest{
			Mock:    mock,
			Request: string(url),
		})
	}

	return json.Marshal(collection)
}

// Reset cleanup mock collection
func Reset() {
	Mocks = make(map[*http.Request]Mock)
}

// init loads mocks from file
func init() {
	LoadFromFile(config.Cfg.MocksFilePath)
}
