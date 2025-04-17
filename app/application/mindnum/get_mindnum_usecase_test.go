package mindnum

import (
	"errors"
	"reflect"
	"testing"

	mindnumDomain "github.com/yanosea/mindnum/v2/app/domain/mindnum"
	mindnumRepo "github.com/yanosea/mindnum/v2/app/infrastructure/text/repository"

	"go.uber.org/mock/gomock"
)

func TestNewGetMindnumUseCase(t *testing.T) {
	mindnumRepo := mindnumRepo.NewMindnumRepository()
	type args struct {
		mindnumRepo mindnumDomain.MindnumRepository
	}
	tests := []struct {
		name string
		args args
		want *GetMindnumUseCase
	}{
		{
			name: "positive case",
			args: args{
				mindnumRepo: mindnumRepo,
			},
			want: &GetMindnumUseCase{
				mindnumRepo: mindnumRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGetMindnumUseCase(tt.args.mindnumRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGetMindnumUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMindnumUseCase_Run(t *testing.T) {
	type fields struct {
		mindnumRepo mindnumDomain.MindnumRepository
	}
	type args struct {
		birthday string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *GetMindnumUsecaseOutputDto
		wantErr bool
		setup   func(mockCtrl *gomock.Controller, tt *fields)
	}{
		{
			name: "positive case",
			fields: fields{
				mindnumRepo: mindnumRepo.NewMindnumRepository(),
			},
			args: args{
				birthday: "19900719",
			},
			want: &GetMindnumUsecaseOutputDto{
				MindNumber:  9,
				Description: mindnumRepo.DescriptionNine,
			},
			wantErr: false,
			setup:   nil,
		},
		{
			name: "negative case (mindnumDomain.GetMindNumber() returns error)",
			fields: fields{
				mindnumRepo: mindnumRepo.NewMindnumRepository(),
			},
			args: args{
				birthday: "",
			},
			want:    nil,
			wantErr: true,
			setup:   nil,
		},
		{
			name: "negative case (mindnumRepo.FindByNumber() returns error)",
			fields: fields{
				mindnumRepo: nil,
			},
			args: args{
				birthday: "19900719",
			},
			want:    nil,
			wantErr: true,
			setup: func(mockCtrl *gomock.Controller, tt *fields) {
				mockMindnumRepo := mindnumDomain.NewMockMindnumRepository(mockCtrl)
				mockMindnumRepo.EXPECT().FindByNumber(9).Return(nil, errors.New("mindnumRepo.FindByNumber() failed"))
				tt.mindnumRepo = mockMindnumRepo
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			if tt.setup != nil {
				tt.setup(mockCtrl, &tt.fields)
			}
			uc := &GetMindnumUseCase{
				mindnumRepo: tt.fields.mindnumRepo,
			}
			got, err := uc.Run(tt.args.birthday)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMindnumUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMindnumUseCase.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
