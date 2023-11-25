package whichisodd

import (
	"errors"
	"fmt"
	"slices"
)

type WhichIsOdd interface {
	FindTheOddInt(input []int) (int, error)
}

type whichIsOdd struct {
}

func NewWhichIsOdd() WhichIsOdd {
	return &whichIsOdd{}
}

func (wo *whichIsOdd) FindTheOddInt(numberList []int) (r int, err error) {
	r = -1

	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
		}
	}()

	compactedList := slices.Clone(numberList)
	compactedList = slices.Compact(compactedList)
	for _, n := range compactedList {
		c := countNumber(numberList, n)
		o := c % 2
		if o != 0 {
			return n, nil
		}
	}

	return r, nil
}

func countNumber(numberList []int, exp int) int {
	c := 0
	for _, n := range numberList {
		if n == exp {
			c++
		}
	}
	return c
}
