package mindnum

import (
	"reflect"
	"testing"
)

func TestNewMindnum(t *testing.T) {
	type args struct {
		mindNumber  int
		description string
	}
	tests := []struct {
		name    string
		args    args
		want    *Mindnum
		wantErr bool
	}{
		{
			name: "positive case",
			args: args{
				mindNumber:  1,
				description: "test",
			},
			want: &Mindnum{
				MindNumber:  1,
				Description: "test",
			},
			wantErr: false,
		},
		{
			name: "negative case (mind number less than 1)",
			args: args{
				mindNumber:  0,
				description: "test",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "negative case (mind number greater than 9)",
			args: args{
				mindNumber:  10,
				description: "test",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMindnum(tt.args.mindNumber, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMindnum() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMindnum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMindNumber(t *testing.T) {
	type args struct {
		birthday string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "positive case",
			args: args{
				birthday: "19900719",
			},
			want:    9,
			wantErr: false,
		},
		{
			name: "negative case (birthday is empty)",
			args: args{
				birthday: "",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "negative case (birth day is invalid)",
			args: args{
				birthday: "1990-07-19",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMindNumber(tt.args.birthday)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMindNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMindNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
