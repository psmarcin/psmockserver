package mock

import (
	"net/http"
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	type args struct {
		id *http.Request
	}
	id, _ := http.NewRequest(http.MethodGet, "/not-foud", nil)
	tests := []struct {
		name    string
		args    args
		want    Mock
		wantErr bool
	}{
		{
			name: "should not find mock",
			args: args{
				id,
			},
			want:    Mock{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Find(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindTimes0(t *testing.T) {
	id, _ := http.NewRequest(http.MethodGet, "/test-mock-times-0", nil)
	m := Mock{
		Body: "test-mock-times-0",
		RemainingTimes: Remaining{
			Times:     0,
			Unlimited: false,
		},
	}
	Add(id, m)
	_, err := Find(id)
	if err != nil {
		return
	}

	t.Error("Find() found but expected not to found")
}

func TestFindQueryString(t *testing.T) {
	id, _ := http.NewRequest(http.MethodGet, "/test-mock-times-0?queryString=test", nil)
	m := Mock{
		Body: "test-mock-times-0",
		RemainingTimes: Remaining{
			Times:     1,
			Unlimited: false,
		},
	}
	Add(id, m)
	_, err := Find(id)
	if err == nil {
		return
	}

	t.Error("Find() found but expected not to found")
}

func TestFindShouldDecreaseTimes(t *testing.T) {
	id, _ := http.NewRequest(http.MethodGet, "/test-mock-times-10", nil)
	m := Mock{
		Body: "test-mock-times-10",
		RemainingTimes: Remaining{
			Times:     10,
			Unlimited: false,
		},
	}
	Add(id, m)
	Find(id)
	Find(id)
	found, err := Find(id)
	if err != nil {
		t.Errorf("Find() expect to find mock but go err %v", err)
	}
	if found.RemainingTimes.Times != 7 {
		t.Errorf("Find().RemainingTimes.Times = %d, want 8", found.RemainingTimes.Times)
	}
	return
}

func TestAdd(t *testing.T) {
	type args struct {
		id   *http.Request
		mock Mock
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
				mock: Mock{
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
			if err := Add(tt.args.id, tt.args.mock); (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReset(t *testing.T) {
	id, _ := http.NewRequest(http.MethodGet, "/", nil)
	// mock
	Add(id, Mock{
		Body: "123",
	})

	if len(Mocks) == 0 {
		t.Fatalf("Didn't set any mock before Reset()")
	}

	// act
	Reset()

	// expect
	if len(Mocks) != 0 {
		t.Fatalf("Reset() should clean all mocks but got %d mocks", len(Mocks))
	}
}
