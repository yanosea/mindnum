package utility

import (
	"reflect"
	"runtime/debug"
	"testing"

	"github.com/yanosea/mindnum/pkg/proxy"

	"go.uber.org/mock/gomock"
)

func TestNewVersionUtil(t *testing.T) {
	debug := proxy.NewDebug()

	type args struct {
		debug proxy.Debug
	}
	tests := []struct {
		name string
		args args
		want VersionUtil
	}{
		{
			name: "positive case",
			args: args{
				debug: debug,
			},
			want: &versionUtil{
				Debug: debug,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVersionUtil(tt.args.debug); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVersionUtil() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_versionUtil_GetVersion(t *testing.T) {
	type fields struct {
		debug proxy.Debug
	}
	type args struct {
		version string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		setup  func(mockCtrl *gomock.Controller, tt *fields)
	}{
		{
			name: "positive case (version is not empty)",
			fields: fields{
				debug: proxy.NewDebug(),
			},
			args: args{
				version: "v0.0.0",
			},
			want:  "v0.0.0",
			setup: nil,
		},
		{
			name: "positive case (version is empty, debug.ReadBuildInfo() returns false)",
			fields: fields{
				debug: nil,
			},
			args: args{
				version: "",
			},
			want: "unknown",
			setup: func(mockCtrl *gomock.Controller, tt *fields) {
				mockDebug := proxy.NewMockDebug(mockCtrl)
				mockDebug.EXPECT().ReadBuildInfo().Return(
					&debug.BuildInfo{},
					false,
				)
				tt.debug = mockDebug
			},
		},
		{
			name: "positive case (version is empty, debug.ReadBuildInfo() returns true and i.Main.Version is not empty)",
			fields: fields{
				debug: nil,
			},
			args: args{
				version: "",
			},
			want: "v0.0.0",
			setup: func(mockCtrl *gomock.Controller, tt *fields) {
				mockDebug := proxy.NewMockDebug(mockCtrl)
				mockDebug.EXPECT().ReadBuildInfo().Return(
					&debug.BuildInfo{
						Main: debug.Module{
							Version: "v0.0.0",
						},
					},
					true,
				)
				tt.debug = mockDebug
			},
		},

		{
			name: "positive case (version is empty, debug.ReadBuildInfo() returns true, i.Main.Version is empty)",
			fields: fields{
				debug: nil,
			},
			args: args{
				version: "",
			},
			want: "dev",
			setup: func(mockCtrl *gomock.Controller, tt *fields) {
				mockDebug := proxy.NewMockDebug(mockCtrl)
				mockDebug.EXPECT().ReadBuildInfo().Return(
					&debug.BuildInfo{
						Main: debug.Module{
							Version: "",
						},
					},
					true,
				)
				tt.debug = mockDebug
			},
		},
	}
	for _, tt := range tests {
		if tt.setup != nil {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			tt.setup(mockCtrl, &tt.fields)
		}
		t.Run(tt.name, func(t *testing.T) {
			v := &versionUtil{
				Debug: tt.fields.debug,
			}
			if got := v.GetVersion(tt.args.version); got != tt.want {
				t.Errorf("versionUtil.GetVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
