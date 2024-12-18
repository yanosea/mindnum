package repository

import (
	mindnumDomain "github.com/yanosea/mindnum/app/domain/mindnum"

	"github.com/yanosea/mindnum/pkg/errors"
)

// mindnumRepository is a struct that implements the MindnumRepository interface.
type mindnumRepository struct{}

// NewMindnumRepository returns a new instance of the mindnumRepository struct.
func NewMindnumRepository() mindnumDomain.MindnumRepository {
	return &mindnumRepository{}
}

// FindByNumber returns a mindnum by number.
func (r *mindnumRepository) FindByNumber(number int) (*mindnumDomain.Mindnum, error) {
	var description string
	switch number {
	case 1:
		description = DescriptionOne
	case 2:
		description = DescriptionTwo
	case 3:
		description = DescriptionThree
	case 4:
		description = DescriptionFour
	case 5:
		description = DescriptionFive
	case 6:
		description = DescriptionSix
	case 7:
		description = DescriptionSeven
	case 8:
		description = DescriptionEight
	case 9:
		description = DescriptionNine
	default:
		return nil, errors.New("number out of range")
	}

	return mindnumDomain.NewMindnum(number, description)
}
