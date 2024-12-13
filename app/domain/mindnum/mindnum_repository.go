package mindnum

import ()

type MindnumRepository interface {
	FindByNumber(number int) (*Mindnum, error)
}
