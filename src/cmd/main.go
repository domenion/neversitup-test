package main

import (
	"fmt"
	"nerversitup-test/src/internal/permutations"
	smileycounter "nerversitup-test/src/internal/smiley-counter"
	whichisodd "nerversitup-test/src/internal/which-is-odd"
	"strconv"
	"strings"
)

func getNumberList(input string) []int {
	ns := []int{}
	ss := strings.Split(input, ",")
	for _, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			continue
		}
		ns = append(ns, n)
	}
	return ns
}

func getFaceList(input string) []string {
	fs := []string{}
	ss := strings.Split(input, ",")
	for _, s := range ss {
		fs = append(fs, strings.TrimSpace(s))
	}
	return fs
}

func main() {
	var input string

	fmt.Println("Permutations")
	fmt.Print("Input: ")
	fmt.Scanln(&input)
	p := permutations.NewPermutations()
	permutations, err := p.GetPermutations(input)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Permutations:", permutations)
	}
	fmt.Println("================================================================")

	fmt.Println("Which is the Odd?")
	fmt.Print("Input: ")
	fmt.Scanln(&input)
	ns := getNumberList(input)
	wo := whichisodd.NewWhichIsOdd()
	odd, err := wo.FindTheOddInt(ns)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("The Odd is", odd)
	}
	fmt.Println("================================================================")

	fmt.Println("Smiley Faces Counter")
	fmt.Print("Input: ")
	fmt.Scanln(&input)
	fs := getFaceList(input)
	sc := smileycounter.NewSmileyCounter()
	sm, err := sc.CountSmileys(fs)
	if err != nil {
	} else {
		fmt.Println("Found smiley faces:", sm)
	}
	fmt.Println("================================================================")
}
