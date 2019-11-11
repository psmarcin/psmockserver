package mock

import (
	"encoding/json"
	"errors"
	"net/http"
	"psmockserver/pkg/config"

	"github.com/kataras/golog"
)

type Mock struct {
	Body           string
	Headers        http.Header
	Method         string
	StatusCode     int
	ContentType    string
	RemainingTimes Remaining
}

type Remaining struct {
	Times     int
	Unlimited bool
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
	mock, ok := Mocks[id]
	if ok != true {
		return mock, errors.New("Not found")
	}

	if mock.RemainingTimes.Unlimited {
		return mock, nil
	}

	if mock.RemainingTimes.Times > 0 {
		mock.RemainingTimes.Times = mock.RemainingTimes.Times - 1
		Mocks[id] = mock
		return mock, nil
	}

	delete(Mocks, id)

	return Mock{}, errors.New("Not found")
}

// List logs all mocked endpoints
func List() {
	golog.Infof("Mocks list: ")
	for x := range Mocks {
		golog.Infof("\t - %s %s", x.Method, x.RequestURI)
	}
}

// Serialize marshals json
func Serialize() ([]byte, error) {
	return json.Marshal(Mocks)
}

// Reset cleanup mock collection
func Reset() {
	Mocks = make(map[*http.Request]Mock)
}

// init loads mocks from file
func init() {
	LoadFromFile(config.Cfg.MocksFilePath)
}
