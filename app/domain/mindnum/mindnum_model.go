package mindnum

import (
	"strconv"
	"strings"
	"time"

	"github.com/yanosea/mindnum/pkg/errors"
)

// Mindnum is a struct that represents a mind number and its description.
type Mindnum struct {
	MindNumber  int
	Description string
}

// NewMindnum returns a new instance of the Mindnum struct.
func NewMindnum(mindNumber int, description string) (*Mindnum, error) {
	if mindNumber < 1 || 9 < mindNumber {
		return nil, errors.New("mind number out of range")
	}
	return &Mindnum{
		MindNumber:  mindNumber,
		Description: description,
	}, nil
}

// GetMindNumber returns a mind number from a birthday.
func GetMindNumber(birthday string) (int, error) {
	if birthday == "" {
		return 0, errors.New("birthday is required")
	}

	if _, err := time.Parse("20060102", birthday); err != nil {
		return 0, errors.New("invalid date format")
	}

	var mindNum int
	tmpNum := birthday
	for {
		tmpSlice := strings.Split(tmpNum, "")
		for _, v := range tmpSlice {
			n, _ := strconv.Atoi(v)
			mindNum += int(n)
		}
		if mindNum < 10 {
			break
		} else {
			tmpNum = strconv.Itoa(mindNum)
			mindNum = 0
		}
	}

	return mindNum, nil
}
