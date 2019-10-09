package utils

import (
	"fmt"
	"net/http"
)

func AddHeaders(source map[string]interface{}) http.Header {
	h := http.Header{}
	// todo: support values different then strings

	for k, v := range source {
		h.Add(k, fmt.Sprintf("%s", v))
	}

	return h
}
