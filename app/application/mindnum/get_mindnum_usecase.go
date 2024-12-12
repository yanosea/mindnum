package mindnum

import (
	mindnumDomain "github.com/yanosea/mindnum/domain/mindnum"
)

type GetMindnumUseCase struct {
	mindnumRepo mindnumDomain.MindnumRepository
}

func NewGetMindnumUseCase(
	mindnumRepo mindnumDomain.MindnumRepository,
) *GetMindnumUseCase {
	return &GetMindnumUseCase{
		mindnumRepo: mindnumRepo,
	}
}

type GetMindnumUsecaseOutputDto struct {
	MindNumber  int
	Description string
}

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
