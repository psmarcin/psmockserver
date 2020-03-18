package mock

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"sync"

	"psmockserver/pkg/config"

	"github.com/kataras/golog"
)

type MockResponse struct {
	Body           string      `json:"body"`
	Headers        http.Header `json:"headers"`
	Method         string      `json:"method"`
	StatusCode     int         `json:"statusCode"`
	ContentType    string      `json:"contentType"`
	RemainingTimes Remaining   `json:"remainingTimes"`
}

type Remaining struct {
	Times     int  `json:"times"`
	Unlimited bool `json:"unlimited"`
}

type Mock struct {
	sync.RWMutex
	Items map[*http.Request]MockResponse
}

// M holds all mocks
var M = Mock{
	Items: make(map[*http.Request]MockResponse),
}

// Add should add new mock to collection
func (m *Mock) Add(id *http.Request, mock MockResponse) error {
	m.Lock()
	defer m.Unlock()

	m.Items[id] = mock
	return nil
}

// Find looks for mock by id in mock collection
func (m *Mock) Find(id *http.Request) (MockResponse, error) {
	m.Lock()
	defer m.Unlock()

	var foundKey *http.Request
	found := false
	mock := MockResponse{}
	for key, value := range m.Items {
		if !reflect.DeepEqual(&key, &id) {
			continue
		}
		found = true
		mock = value
		foundKey = key
		break
	}

	if found != true {
		return mock, errors.New("Not found")
	}

	if mock.RemainingTimes.Unlimited {
		return mock, nil
	}

	if mock.RemainingTimes.Times > 0 {
		mock.RemainingTimes.Times = mock.RemainingTimes.Times - 1
		m.Items[foundKey] = mock
		return mock, nil
	}

	delete(m.Items, foundKey)

	return MockResponse{}, errors.New("Not found")
}

// List logs all mocked endpoints
func (m *Mock) List() {
	m.RLock()
	defer m.RUnlock()

	golog.Infof("Mocks list: ")
	for x := range m.Items {
		golog.Infof("\t - %s %s %s", x.Method, x.URL.String(), x.URL.Query().Encode())
	}
}

// Serialize marshals json
func (m *Mock) Serialize() ([]byte, error) {
	m.RLock()
	defer m.RUnlock()

	var collection = make([]MockResponse, len(m.Items))
	for _, mock := range m.Items {
		collection = append(collection, mock)
	}

	return json.Marshal(collection)
}

// Reset cleanup mock collection
func (m *Mock) Reset() {
	m.Lock()
	defer m.Unlock()

	m.Items = make(map[*http.Request]MockResponse)
}

// init loads mocks from file
func init() {
	LoadFromFile(config.Cfg.MocksFilePath)
}
