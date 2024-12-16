package mindnum

import ()

// MindnumRepository is an interface that provides the repository of the mindnum.
type MindnumRepository interface {
	FindByNumber(number int) (*Mindnum, error)
}
