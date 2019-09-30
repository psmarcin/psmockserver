package mock

import (
	"encoding/json"
	"errors"
	"net/http"

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

var Mocks = make(map[string]Mock)

func Add(id string, mock Mock) error {
	Mocks[id] = mock
	return nil
}

func Find(id string) (Mock, error) {
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

func List() {
	golog.Infof("Mocks list: ")
	for x := range Mocks {
		golog.Infof("\t - %s", x)
	}
}

func Serialize() ([]byte, error) {
	return json.Marshal(Mocks)
}

func DefaultMocks() {
	Add(GetMockHash(http.MethodGet, "/"), Mock{
		Headers:     http.Header{},
		StatusCode:  http.StatusOK,
		Body:        "test body",
		Method:      http.MethodGet,
		ContentType: "text/plain",
		RemainingTimes: Remaining{
			Times:     10,
			Unlimited: true,
		},
	})

	Add(GetMockHash(http.MethodPost, "/x"), Mock{
		Headers:     http.Header{},
		StatusCode:  http.StatusOK,
		Body:        "test body1",
		Method:      http.MethodPost,
		ContentType: "text/asdlasdaslda",
		RemainingTimes: Remaining{
			Times:     10,
			Unlimited: true,
		},
	})
}

func init() {
	DefaultMocks()
}
