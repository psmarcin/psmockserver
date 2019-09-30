package server

import (
	"net/http"
	"reflect"
	"testing"
)

func Test_addHeaders(t *testing.T) {
	type args struct {
		source map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want http.Header
	}{
		{
			name: "should set headers from map",
			args: args{
				source: map[string]interface{}{
					"Test": "123",
				},
			},
			want: http.Header{
				"Test": []string{"123"},
			},
		},
		{
			name: "should set headers from map, first character uppercase",
			args: args{
				source: map[string]interface{}{
					"test": "123",
				},
			},
			want: http.Header{
				"Test": []string{"123"},
			},
		},
		{
			name: "should set headers from map, multiple keys",
			args: args{
				source: map[string]interface{}{
					"test":         "123",
					"content-type": "text/html",
					"ContentType":  "application/json",
				},
			},
			want: http.Header{
				"Test":         []string{"123"},
				"Content-Type": []string{"text/html"},
				"Contenttype":  []string{"application/json"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addHeaders(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}
