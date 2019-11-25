package mock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	DefaultContentType = "application/json"
	DefaultBody        = "{\"defaultBody\": true}"
	DefaultPath        = "/"
)

type RequestId struct {
	Method       string
	Path         string
	QueryStrings url.Values
}

type Payload struct {
	HttpRequest  `json:"httpRequest"`
	HttpResponse `json:"httpResponse"`
	Times        `json:"times"`
}

type HttpRequest struct {
	Method       string     `json:"method"`
	Path         string     `json:"path"`
	ContentType  string     `json:"content-type"`
	QueryStrings url.Values `json:"queryStringParameters"`
}

type HttpResponse struct {
	StatusCode int                    `json:"statusCode"`
	Headers    map[string]interface{} `json:"headers"`
	Body       string                 `json:"body"`
}
type Times struct {
	RemainingTimes int  `json:"remainingTimes"`
	Unlimited      bool `json:"unlimited"`
}

func Parse(body []byte) (Payload, error) {
	payload := Payload{
		HttpRequest: HttpRequest{
			Method:      http.MethodGet,
			Path:        DefaultPath,
			ContentType: DefaultContentType,
		},
		HttpResponse: HttpResponse{
			StatusCode: http.StatusOK,
			Headers:    map[string]interface{}{},
			Body:       DefaultBody,
		},
		Times: Times{
			RemainingTimes: 1,
			Unlimited:      true,
		},
	}
	err := json.Unmarshal(body, &payload)
	if err != nil {
		return payload, fmt.Errorf("can't parse mock payload %w", err)
	}

	return payload, nil
}

// GetMockHash returns id for mock
func GetMockHash(params RequestId) *http.Request {
	url, _ := url.Parse(params.Path)
	// add query to parsed url
	query := url.Query()
	for key, values := range params.QueryStrings {
		for _, value := range values {
			query.Add(key, value)
		}
	}
	url.RawQuery = query.Encode()
	// create request
	id, _ := http.NewRequest(params.Method, params.Path, nil)
	// set url with query to request
	id.URL = url
	return id
}
