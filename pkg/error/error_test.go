package error

import (
	"errors"
	"testing"
)

func TestWrap(t *testing.T) {
	type args struct {
		err error
		msg string
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		errMessage string
	}{
		{
			name: "should wrap error with message",
			args: args{
				err: errors.New("test"),
				msg: "added_message",
			},
			wantErr:    true,
			errMessage: "added_message test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Wrap(tt.args.err, tt.args.msg)
			if err != nil != tt.wantErr {
				t.Errorf("Wrap() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.errMessage != err.Error() {
				t.Errorf("Wrap() error = %v, errMessage %v", err.Error(), tt.errMessage)
			}
		})
	}
}
