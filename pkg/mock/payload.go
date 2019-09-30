package mock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/kataras/golog"
)

type Payload struct {
	HttpRequest  `json:"httpRequest"`
	HttpResponse `json:"httpResponse"`
	Times        `json:"times"`
}

type HttpRequest struct {
	Method      string `json:"method"`
	Path        string `json:"path"`
	ContentType string `json:"content-type"`
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

const MockDefaultBody = `{
					"defaultBody": true
				}`

func Parse(body []byte) (Payload, error) {
	payload := Payload{
		HttpRequest: HttpRequest{
			Method:      http.MethodGet,
			Path:        "/",
			ContentType: "application/json",
		},
		HttpResponse: HttpResponse{
			StatusCode: http.StatusOK,
			Headers:    map[string]interface{}{},
			Body:       MockDefaultBody,
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
	golog.Infof("Got schema %+v", payload)

	return payload, nil
}

func GetMockHash(method, path string) string {
	return fmt.Sprintf("%s|%s", strings.ToUpper(method), path)
}
