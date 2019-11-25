package mock

import (
	"encoding/json"
	"io/ioutil"
	"psmockserver/pkg/utils"

	"github.com/kataras/golog"
)

// LoadFromFile loads default mocks from file from given path
func LoadFromFile(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		golog.Debugf("Can't read file %+v", err)
		return
	}

	payloads, err := parseFileMocks(content)
	if err != nil {
		golog.Errorf("Can't parse file content %s with content %s", path, content)
		return
	}
	for _, p := range payloads {
		M.Add(GetMockHash(RequestId{
			Method:       p.HttpRequest.Method,
			Path:         p.HttpRequest.Path,
			QueryStrings: p.HttpRequest.QueryStrings,
		}), MockResponse{
			Headers:     utils.AddHeaders(p.HttpResponse.Headers),
			StatusCode:  p.HttpResponse.StatusCode,
			Body:        p.HttpResponse.Body,
			ContentType: p.HttpRequest.ContentType,
			Method:      p.HttpRequest.Method,
			RemainingTimes: Remaining{
				Times:     p.Times.RemainingTimes,
				Unlimited: p.Times.Unlimited,
			},
		})
		golog.Infof("Mock added for %s %s", p.HttpRequest.Method, p.HttpRequest.Path)
	}
}

func parseFileMocks(source []byte) ([]Payload, error) {
	var mocks []Payload

	err := json.Unmarshal(source, &mocks)
	return mocks, err
}
