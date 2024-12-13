package mindnum

import (
	"errors"
	"reflect"
	"testing"

	mindnumDomain "github.com/yanosea/mindnum/app/domain/mindnum"
	mindnumRepo "github.com/yanosea/mindnum/app/infrastructure/text/repository"

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
				MindNumber: 9,
				Description: `You are a BALANCER.

BALANCER is like...
The key word in your life is "2".
You are a person who is destined to have two of everything.
You may travel back and forth between the country and abroad, have two homes, or be attracted to two people at the same time.
You have an inquisitive mind and take pleasure in expanding your knowledge in areas you are not familiar with.
They have a keen intuition and sense of perception,
so they may feel pain at events they don't feel comfortable with or at relationships that are only superficial.
Creative workplaces that allow you to express your own opinions are
more suitable for you than simple work or general office work that does not allow you to feel change.
You may get double blessings such as pregnancy and childbirth at the same time of marriage.
`,
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
