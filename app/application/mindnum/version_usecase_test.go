package mindnum

import (
	"reflect"
	"testing"
)

func TestNewVersionUseCase(t *testing.T) {
	tests := []struct {
		name string
		want *VersionUseCase
	}{
		{
			name: "positive case",
			want: &VersionUseCase{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVersionUseCase(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVersionUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersionUseCase_Run(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		uc   *VersionUseCase
		args args
		want *VersionUsecaseOutputDto
	}{
		{
			name: "positive case",
			uc:   &VersionUseCase{},
			args: args{
				version: "v0.0.0",
			},
			want: &VersionUsecaseOutputDto{
				Version: "v0.0.0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &VersionUseCase{}
			if got := uc.Run(tt.args.version); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VersionUseCase.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
