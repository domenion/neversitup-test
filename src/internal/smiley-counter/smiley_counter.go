package smileycounter

import (
	"errors"
	"fmt"
	"slices"
)

type SmileyCounter interface {
	CountSmileys(input []string) (int, error)
}

type smileycounter struct{}

func NewSmileyCounter() SmileyCounter {
	return &smileycounter{}
}

var eyesList = []string{":", ";"}
var noseList = []string{"-", "~"}
var mouthList = []string{")", "D"}

func (sc *smileycounter) CountSmileys(input []string) (c int, err error) {
	c = 0

	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
		}
	}()

	for _, s := range input {
		if validateFace(s) {
			c++
		}
	}

	return c, nil
}

func validateFace(s string) bool {
	if len(s) < 2 {
		return false
	}
	if !slices.Contains(eyesList, string(s[0])) {
		return false
	}
	if len(s) > 2 {
		if !slices.Contains(noseList, string(s[1])) {
			return false
		}
		if !slices.Contains(mouthList, string(s[2])) {
			return false
		}
	} else {
		if !slices.Contains(mouthList, string(s[1])) {
			return false
		}
	}
	return true
}
