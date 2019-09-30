package mock

import (
	"net/http"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Payload
		wantErr bool
	}{
		{
			name: "default values",
			args: args{
				body: []byte("{}"),
			},
			want: Payload{
				HttpRequest: HttpRequest{
					Method:      "GET",
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
			},
			wantErr: false,
		},
		{
			name: "get request method DELETE",
			args: args{
				body: []byte(`{
					"httpRequest":{
						"Method":"DELETE"
					}
				}`),
			},
			want: Payload{
				HttpRequest: HttpRequest{
					Method:      "DELETE",
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
			},
			wantErr: false,
		},
		{
			name: "get request path /some/path/nested/many/times",
			args: args{
				body: []byte(`{
					"httpRequest":{
						"Path":"/some/path/nested/many/times"
					}
				}`),
			},
			want: Payload{
				HttpRequest: HttpRequest{
					Method:      "GET",
					Path:        "/some/path/nested/many/times",
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
			},
			wantErr: false,
		},
		{
			name: "get request content-type text/html",
			args: args{
				body: []byte(`{
					"httpRequest":{
						"content-type":"text/html"
					}
				}`),
			},
			want: Payload{
				HttpRequest: HttpRequest{
					Method:      "GET",
					Path:        "/",
					ContentType: "text/html",
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
			},
			wantErr: false,
		},
		{
			name: "get response statusCode 301",
			args: args{
				body: []byte(`{
					"httpResponse":{
						"statusCode":301
					}
				}`),
			},
			want: Payload{
				HttpRequest: HttpRequest{
					Method:      "GET",
					Path:        "/",
					ContentType: "application/json",
				},
				HttpResponse: HttpResponse{
					StatusCode: 301,
					Headers:    map[string]interface{}{},
					Body:       MockDefaultBody,
				},
				Times: Times{
					RemainingTimes: 1,
					Unlimited:      true,
				},
			},
			wantErr: false,
		},
		{
			name: "get response body lipsum",
			args: args{
				body: []byte(`{
					"httpResponse":{
						"body":"lipsum"
					}
				}`),
			},
			want: Payload{
				HttpRequest: HttpRequest{
					Method:      "GET",
					Path:        "/",
					ContentType: "application/json",
				},
				HttpResponse: HttpResponse{
					StatusCode: http.StatusOK,
					Headers:    map[string]interface{}{},
					Body:       "lipsum",
				},
				Times: Times{
					RemainingTimes: 1,
					Unlimited:      true,
				},
			},
			wantErr: false,
		},
		{
			name: "get response body lipsum",
			args: args{
				body: []byte(`{
					"httpResponse":{
						"body":"lipsum"
					}
				}`),
			},
			want: Payload{
				HttpRequest: HttpRequest{
					Method:      "GET",
					Path:        "/",
					ContentType: "application/json",
				},
				HttpResponse: HttpResponse{
					StatusCode: http.StatusOK,
					Headers:    map[string]interface{}{},
					Body:       "lipsum",
				},
				Times: Times{
					RemainingTimes: 1,
					Unlimited:      true,
				},
			},
			wantErr: false,
		},
		{
			name: "get times",
			args: args{
				body: []byte(`{
					"times":{
						"remainingTimes":99,
						"unlimited": true
					}
				}`),
			},
			want: Payload{HttpRequest: HttpRequest{
				Method:      "GET",
				Path:        "/",
				ContentType: "application/json",
			},
				HttpResponse: HttpResponse{
					StatusCode: http.StatusOK,
					Headers:    map[string]interface{}{},
					Body:       MockDefaultBody,
				},
				Times: Times{
					RemainingTimes: 99,
					Unlimited:      true,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMockHash(t *testing.T) {
	type args struct {
		method string
		path   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "METHOD=get, PATH=/",
			args: args{
				method: "get",
				path:   "/",
			},
			want: "GET|/",
		},
		{
			name: "METHOD=GET, PATH=/",
			args: args{
				method: "GET",
				path:   "/",
			},
			want: "GET|/",
		},
		{
			name: "METHOD=poST, PATH=/",
			args: args{
				method: "poST",
				path:   "/",
			},
			want: "POST|/",
		},
		{
			name: "METHOD=GET, PATH=/nested/values",
			args: args{
				method: "GET",
				path:   "/nested/values",
			},
			want: "GET|/nested/values",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMockHash(tt.args.method, tt.args.path); got != tt.want {
				t.Errorf("GetMockHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
