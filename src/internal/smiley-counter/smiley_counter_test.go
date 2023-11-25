package smileycounter

import "testing"

func TestCase1(t *testing.T) {
	expected := 2
	input := []string{":)", ";(", ";}", ":-D"}
	sc := NewSmileyCounter()
	output, err := sc.CountSmileys(input)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if output != expected {
		t.Error("failed", output)
		return
	}
}

func TestCase2(t *testing.T) {
	expected := 3
	input := []string{";D", ":-(", ":-)", ";~)"}
	sc := NewSmileyCounter()
	output, err := sc.CountSmileys(input)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if output != expected {
		t.Error("failed", output)
		return
	}
}

func TestCase3(t *testing.T) {
	expected := 1
	input := []string{";]", ":[", ";*", ":$", ";-D"}
	sc := NewSmileyCounter()
	output, err := sc.CountSmileys(input)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if output != expected {
		t.Error("failed", output)
		return
	}
}

func TestCaseEx(t *testing.T) {
	expected := 0
	input := []string{}
	sc := NewSmileyCounter()
	output, err := sc.CountSmileys(input)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if output != expected {
		t.Error("failed", output)
		return
	}
}
