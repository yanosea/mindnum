package repository

import (
	"fmt"
	"path/filepath"

	mindnumDomain "github.com/yanosea/mindnum/app/domain/mindnum"

	"github.com/yanosea/mindnum/pkg/errors"
)

type mindnumRepository struct {
	DescriptionPath string
}

func NewMindnumRepository(descriptionPath string) mindnumDomain.MindnumRepository {
	if descriptionPath == "" {
		descriptionPath = mindnumDomain.DescriptionPath
	}
	return &mindnumRepository{
		DescriptionPath: descriptionPath,
	}
}

func (r *mindnumRepository) FindByNumber(number int) (*mindnumDomain.Mindnum, error) {
	if number < 1 || 9 < number {
		return nil, errors.New("number out of range")
	}
	fileName := fmt.Sprintf("%d.txt", number)

	description, err := mindnumDomain.Embedded.ReadFile(filepath.Join(r.DescriptionPath, fileName))
	if err != nil {
		return nil, err
	}

	return mindnumDomain.NewMindnum(number, string(description))
}
