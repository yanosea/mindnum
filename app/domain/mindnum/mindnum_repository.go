package mindnum

import ()

var (
	DescriptionPath = "descriptions"
)

type MindnumRepository interface {
	FindByNumber(number int) (*Mindnum, error)
}
