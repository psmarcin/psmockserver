package mock

import (
	"net/http"
	"testing"
)

func TestFindNoMock(t *testing.T) {
	m := Mock{}
	id, _ := http.NewRequest(http.MethodGet, "/test-url/nested", nil)
	_, err := m.Find(id)
	if err == nil {
		t.Fatalf("Find() didn't return error but expected to")
	}
}

func TestFindByPathAndMethod(t *testing.T) {
	m := Mock{
		Items: make(map[*http.Request]MockResponse),
	}
	path := "/test-endpoint"
	method := http.MethodGet
	mockID, _ := http.NewRequest(method, path, nil)
	responseBody := "response"
	m.Add(mockID, MockResponse{
		Body: responseBody,
		RemainingTimes: Remaining{
			Unlimited: true,
		},
	})
	id, _ := http.NewRequest(method, path, nil)

	found, err := m.Find(id)

	if err != nil {
		t.Fatalf("Find() returns error but expected not to %+v", err)
	}
	if found.Body != responseBody {
		t.Fatalf("Find() returns body %s but expected %s", found.Body, responseBody)
	}
}

func TestFindTimes0(t *testing.T) {
	m := Mock{
		Items: make(map[*http.Request]MockResponse),
	}
	id, _ := http.NewRequest(http.MethodGet, "/test-mock-times-0", nil)
	mock := MockResponse{
		Body: "test-mock-times-0",
		RemainingTimes: Remaining{
			Times:     0,
			Unlimited: false,
		},
	}
	m.Add(id, mock)
	_, err := m.Find(id)
	if err != nil {
		return
	}

	t.Error("Find() found but expected not to found")
}

func TestFindByQueryString(t *testing.T) {
	id, _ := http.NewRequest(http.MethodGet, "/test-mock-times-0?queryString=test", nil)
	m := MockResponse{
		Body: "test-mock-times-0",
		RemainingTimes: Remaining{
			Times:     1,
			Unlimited: false,
		},
	}
	M.Add(id, m)
	_, err := M.Find(id)
	if err == nil {
		return
	}

	t.Error("Find() found but expected not to found")
}

func TestFindShouldDecreaseTimes(t *testing.T) {
	id, _ := http.NewRequest(http.MethodGet, "/test-mock-times-10", nil)
	m := MockResponse{
		Body: "test-mock-times-10",
		RemainingTimes: Remaining{
			Times:     10,
			Unlimited: false,
		},
	}
	M.Add(id, m)
	M.Find(id)
	M.Find(id)
	found, err := M.Find(id)
	if err != nil {
		t.Errorf("Find() expect to find mock but go err %v", err)
	}
	if found.RemainingTimes.Times != 7 {
		t.Errorf("Find().RemainingTimes.Times = %d, want 8", found.RemainingTimes.Times)
	}
	return
}

func BenchmarkFind(b *testing.B) {
	m := Mock{
		Items: make(map[*http.Request]MockResponse),
	}
	path := "/test-endpoint"
	method := http.MethodGet
	mockID, _ := http.NewRequest(method, path, nil)
	responseBody := "response"
	m.Add(mockID, MockResponse{
		Body: responseBody,
		RemainingTimes: Remaining{
			Unlimited: true,
		},
	})
	id, _ := http.NewRequest(method, path, nil)

	// run the Add method b.N times
	for n := 0; n < b.N; n++ {
		m.Find(id)
	}
}

func TestAdd(t *testing.T) {
	m := Mock{
		Items: make(map[*http.Request]MockResponse),
	}
	type args struct {
		id   *http.Request
		mock MockResponse
	}
	id, _ := http.NewRequest(http.MethodGet, "/", nil)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should add mock to root handler",
			args: args{
				id: id,
				mock: MockResponse{
					Body:        "test",
					Headers:     http.Header{},
					Method:      http.MethodGet,
					StatusCode:  http.StatusOK,
					ContentType: "application/json",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := m.Add(tt.args.id, tt.args.mock); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func BenchmarkAdd(b *testing.B) {
	m := Mock{
		Items: make(map[*http.Request]MockResponse),
	}
	id, _ := http.NewRequest(http.MethodGet, "/test-mock-times-0", nil)
	mock := MockResponse{
		Body: "test-mock-times-0",
		RemainingTimes: Remaining{
			Times:     0,
			Unlimited: false,
		},
	}
	// run the Add method b.N times
	for n := 0; n < b.N; n++ {
		m.Add(id, mock)
	}
}

func TestReset(t *testing.T) {
	m := Mock{
		Items: make(map[*http.Request]MockResponse),
	}
	id, _ := http.NewRequest(http.MethodGet, "/", nil)
	// mock
	m.Add(id, MockResponse{
		Body: "123",
	})

	if len(m.Items) == 0 {
		t.Fatalf("Didn't set any mock before Reset()")
	}

	// act
	m.Reset()

	// expect
	if len(m.Items) != 0 {
		t.Fatalf("Reset() should clean all mocks but got %d mocks", len(m.Items))
	}
}
