package mindnum

import (
	"strconv"
	"strings"
	"time"

	"github.com/yanosea/mindnum/pkg/errors"
)

type Mindnum struct {
	MindNumber  int
	Description string
}

func NewMindnum(mindNumber int, description string) (*Mindnum, error) {
	if mindNumber < 1 || 9 < mindNumber {
		return nil, errors.New("mind number out of range")
	}
	return &Mindnum{
		MindNumber:  mindNumber,
		Description: description,
	}, nil
}

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
