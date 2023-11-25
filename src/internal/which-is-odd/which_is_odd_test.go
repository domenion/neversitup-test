package whichisodd

import "testing"

func TestCase1(t *testing.T) {
	expected := 7
	input := []int{7}
	wo := NewWhichIsOdd()
	output, err := wo.FindTheOddInt(input)
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
	expected := 0
	input := []int{0}
	wo := NewWhichIsOdd()
	output, err := wo.FindTheOddInt(input)
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
	expected := 2
	input := []int{1, 1, 2}
	wo := NewWhichIsOdd()
	output, err := wo.FindTheOddInt(input)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if output != expected {
		t.Error("failed", output)
		return
	}
}

func TestCase4(t *testing.T) {
	expected := 0
	input := []int{0, 1, 0, 1, 0}
	wo := NewWhichIsOdd()
	output, err := wo.FindTheOddInt(input)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if output != expected {
		t.Error("failed", output)
		return
	}
}

func TestCase5(t *testing.T) {
	expected := 4
	input := []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}
	wo := NewWhichIsOdd()
	output, err := wo.FindTheOddInt(input)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if output != expected {
		t.Error("failed", output)
		return
	}
}
