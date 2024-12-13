package errors

import (
	"errors"
	"testing"
)

func Test_wrappedError_Error(t *testing.T) {
	type fields struct {
		err error
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "positive case (e.err is nil and e.msg is empty)",
			fields: fields{
				err: nil,
				msg: "",
			},
			want: "",
		},
		{
			name: "positive case (e.err is nil and e.msg is not empty)",
			fields: fields{
				err: nil,
				msg: "test",
			},
			want: "test",
		},
		{
			name: "positive case (e.err is not nil and e.msg is empty)",
			fields: fields{
				err: errors.New("test"),
				msg: "",
			},
			want: "test",
		},
		{
			name: "positive case (e.err is not nil and e.msg is not empty)",
			fields: fields{
				err: errors.New("test"),
				msg: "test",
			},
			want: "test : test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &wrappedError{
				err: tt.fields.err,
				msg: tt.fields.msg,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("wrappedError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "positive case",
			args: args{
				msg: "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := New(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	type args struct {
		err error
		msg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "positive case",
			args: args{
				err: errors.New("test"),
				msg: "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Wrap(tt.args.err, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Wrap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
