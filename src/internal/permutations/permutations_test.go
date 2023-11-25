package permutations

import (
	"slices"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []string{"a"}
	input := "a"
	p := NewPermutations()
	output, err := p.GetPermutations(input)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if slices.Compare(expected, output) != 0 {
		t.Error("failed", output)
		return
	}
}

func TestCase2(t *testing.T) {
	expected := []string{"ab", "ba"}
	input := "ab"
	p := NewPermutations()
	output, err := p.GetPermutations(input)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if slices.Compare(expected, output) != 0 {
		t.Error("failed", output)
		return
	}
}

func TestCase3(t *testing.T) {
	expected := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
	input := "abc"
	p := NewPermutations()
	output, err := p.GetPermutations(input)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if slices.Compare(expected, output) != 0 {
		t.Error("failed", output)
		return
	}
}

func TestCase4(t *testing.T) {
	expected := []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}
	input := "aabb"
	p := NewPermutations()
	output, err := p.GetPermutations(input)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if slices.Compare(expected, output) != 0 {
		t.Error("failed", output)
		return
	}
}
