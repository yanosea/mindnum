package formatter

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewFormatter(t *testing.T) {
	type args struct {
		format string
	}
	tests := []struct {
		name    string
		args    args
		want    Formatter
		wantErr bool
	}{
		{
			name: "positive case (text)",
			args: args{
				format: "text",
			},
			want:    &TextFormatter{},
			wantErr: false,
		},
		{
			name: "negative case (invalid format)",
			args: args{
				format: "invalid",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFormatter(tt.args.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFormatter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFormatter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendErrorToOutput(t *testing.T) {
	type args struct {
		err    error
		output string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "positive case (error and output are empty)",
			args: args{
				err:    nil,
				output: "",
			},
			want: "",
		},
		{
			name: "positive case (error is not empty and output is empty)",
			args: args{
				err:    errors.New("test"),
				output: "",
			},
			want: "Error : test",
		},
		{
			name: "positive case (error is not empty and output is not empty)",
			args: args{
				err:    errors.New("test"),
				output: "test",
			},
			want: "Error : test\ntest",
		},
		{
			name: "positive case (error is empty and output is not empty)",
			args: args{
				err:    nil,
				output: "test",
			},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AppendErrorToOutput(tt.args.err, tt.args.output); got != tt.want {
				t.Errorf("AppendErrorToOutput() = %v, want %v", got, tt.want)
			}
		})
	}
}
