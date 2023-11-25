package permutations

import (
	"errors"
	"fmt"
	"slices"
)

type Permutations interface {
	GetPermutations(input string) ([]string, error)
}

type permutations struct {
}

func NewPermutations() Permutations {
	return &permutations{}
}

func (p *permutations) GetPermutations(input string) (output []string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(fmt.Sprint(r))
		}
	}()

	chars := []string{}
	for _, v := range input {
		chars = append(chars, string(v))
	}
	output = shuffle(chars, "", []string{})
	output = deduplicate(output)
	slices.Sort(output)
	return output, nil
}

func shuffle(chars []string, s string, output []string) []string {
	if len(chars) == 1 {
		s = s + chars[0]
		return append(output, s)
	}

	l := len(chars)
	for i := 0; i < l; i++ {
		ss := s + chars[0]
		output = shuffle(chars[1:], ss, output)
		chars = shift(chars)
	}
	return output
}

func shift(chars []string) []string {
	if len(chars) == 1 {
		return chars
	}
	return append(chars[1:], chars[0])
}

func deduplicate(output []string) []string {
	temp := make([]string, 0, len(output))
	for i := 0; i < len(output); i++ {
		current := output[i]
		isContains := slices.Contains(temp, current)
		if !isContains {
			temp = append(temp, current)
		}
	}
	return temp
}
