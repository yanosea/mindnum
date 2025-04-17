package mindnum

import (
	mindnumDomain "github.com/yanosea/mindnum/v2/app/domain/mindnum"
)

// GetMindnumUseCase is a struct that provides the use case of getting a mind number.
type GetMindnumUseCase struct {
	mindnumRepo mindnumDomain.MindnumRepository
}

// NewGetMindnumUseCase returns a new instance of the GetMindnumUseCase struct.
func NewGetMindnumUseCase(
	mindnumRepo mindnumDomain.MindnumRepository,
) *GetMindnumUseCase {
	return &GetMindnumUseCase{
		mindnumRepo: mindnumRepo,
	}
}

// GetMindnumUsecaseOutputDto is a struct that provides the output DTO of the GetMindnumUseCase struct.
type GetMindnumUsecaseOutputDto struct {
	MindNumber  int
	Description string
}

// Run returns the output of the GetMindnumUseCase struct.
func (uc *GetMindnumUseCase) Run(birthday string) (*GetMindnumUsecaseOutputDto, error) {
	mindNumber, err := mindnumDomain.GetMindNumber(birthday)
	if err != nil {
		return nil, err
	}

	mindnum, err := uc.mindnumRepo.FindByNumber(mindNumber)
	if err != nil {
		return nil, err
	}

	return &GetMindnumUsecaseOutputDto{
		MindNumber:  mindnum.MindNumber,
		Description: mindnum.Description,
	}, nil
}
