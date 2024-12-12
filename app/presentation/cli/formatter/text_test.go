package formatter

import (
	"reflect"
	"testing"

	mindnumApp "github.com/yanosea/mindnum/application/mindnum"
)

func TestNewTextFormatter(t *testing.T) {
	tests := []struct {
		name string
		want *TextFormatter
	}{
		{
			name: "positive case",
			want: &TextFormatter{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTextFormatter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTextFormatter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTextFormatter_Format(t *testing.T) {
	type args struct {
		result interface{}
	}
	tests := []struct {
		name string
		f    *TextFormatter
		args args
		want string
	}{
		{
			name: "positive case (VersionUsecaseOutputDto)",
			f:    &TextFormatter{},
			args: args{
				result: &mindnumApp.VersionUsecaseOutputDto{
					Version: "v0.0.0",
				},
			},
			want: "mindnum version v0.0.0",
		},
		{
			name: "positive case (GetMindnumUsecaseOutputDto)",
			f:    &TextFormatter{},
			args: args{
				result: &mindnumApp.GetMindnumUsecaseOutputDto{
					MindNumber:  1,
					Description: "description\n",
				},
			},
			want: "Your mind number is 1!\n\ndescription",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &TextFormatter{}
			got := f.Format(tt.args.result)
			if got != tt.want {
				t.Errorf("TextFormatter.Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
