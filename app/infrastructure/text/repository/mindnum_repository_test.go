package repository

import (
	"reflect"
	"testing"

	mindnumDomain "github.com/yanosea/mindnum/app/domain/mindnum"
)

func TestNewMindnumRepository(t *testing.T) {
	tests := []struct {
		name string
		want mindnumDomain.MindnumRepository
	}{
		{
			name: "positive case",
			want: &mindnumRepository{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMindnumRepository(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMindnumRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mindnumRepository_FindByNumber(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name    string
		args    args
		want    *mindnumDomain.Mindnum
		wantErr bool
	}{
		{
			name: "positive case (number is 1)",
			args: args{
				number: 1,
			},
			want: &mindnumDomain.Mindnum{
				MindNumber:  1,
				Description: DescriptionOne,
			},
			wantErr: false,
		},
		{
			name: "positive case (number is 2)",
			args: args{
				number: 2,
			},
			want: &mindnumDomain.Mindnum{
				MindNumber:  2,
				Description: DescriptionTwo,
			},
			wantErr: false,
		},
		{
			name: "positive case (number is 3)",
			args: args{
				number: 3,
			},
			want: &mindnumDomain.Mindnum{
				MindNumber:  3,
				Description: DescriptionThree,
			},
			wantErr: false,
		},
		{
			name: "positive case (number is 4)",
			args: args{
				number: 4,
			},
			want: &mindnumDomain.Mindnum{
				MindNumber:  4,
				Description: DescriptionFour,
			},
			wantErr: false,
		},
		{
			name: "positive case (number is 5)",
			args: args{
				number: 5,
			},
			want: &mindnumDomain.Mindnum{
				MindNumber:  5,
				Description: DescriptionFive,
			},
			wantErr: false,
		},
		{
			name: "positive case (number is 6)",
			args: args{
				number: 6,
			},
			want: &mindnumDomain.Mindnum{
				MindNumber:  6,
				Description: DescriptionSix,
			},
			wantErr: false,
		},
		{
			name: "positive case (number is 7)",
			args: args{
				number: 7,
			},
			want: &mindnumDomain.Mindnum{
				MindNumber:  7,
				Description: DescriptionSeven,
			},
			wantErr: false,
		},
		{
			name: "positive case (number is 8)",
			args: args{
				number: 8,
			},
			want: &mindnumDomain.Mindnum{
				MindNumber:  8,
				Description: DescriptionEight,
			},
			wantErr: false,
		},
		{
			name: "positive case (number is 9)",
			args: args{
				number: 9,
			},
			want: &mindnumDomain.Mindnum{
				MindNumber:  9,
				Description: DescriptionNine,
			},
			wantErr: false,
		},
		{
			name: "negative case (number is 0)",
			args: args{
				number: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "negative case (number is 10)",
			args: args{
				number: 10,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mindnumRepository{}
			got, err := r.FindByNumber(tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("mindnumRepository.FindByNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mindnumRepository.FindByNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
