package main

import (
	"fmt"
	"strings"

	"github.com/bigspawn/adventofcode-2022/day/reader"
)

/*
	A for Rock, B for Paper, and C for Scissors
	X for Rock, Y for Paper, and Z for Scissors

	shape you selected
	(1 for Rock, 2 for Paper, and 3 for Scissors)
	plus the score for the outcome of the round
	(0 if you lost, 3 if the round was a draw, and 6 if you won)

	A Y
	B X
	C Z
*/

var points = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

type action uint

const (
	lost action = iota
	draw
	won
)

func (a action) String() string {
	switch a {
	case lost:
		return "lost"
	case won:
		return "won"
	case draw:
		return "draw"
	}
	return ""
}

func (a action) Revert() action {
	switch a {
	case lost:
		return won
	case won:
		return lost
	}
	return draw
}

var actionPoints = map[action]int{
	lost: 0,
	draw: 3,
	won:  6,
}

func main() {
	total := 0
	err := reader.Read(firstStrategy(&total))
	if err != nil {
		panic(err)
	}

	fmt.Println(total)
}

func firstStrategy(total *int) func(l string) error {
	return func(l string) error {
		if l == "" {
			return nil
		}

		arr := strings.Split(l, " ")
		if len(arr) != 2 {
			return fmt.Errorf("wrong line")
		}

		p1 := arr[0]
		p2 := arr[1]
		a := getAction(p1, p2) // for first player, but I am second
		a = a.Revert()         // for me it would be revert result

		i := points[p2]
		i2 := actionPoints[a]
		*total += i + i2

		//printName(p1)
		//fmt.Print("(", p1, ")")
		//fmt.Print(" <-> ")
		//printName(p2)
		//fmt.Print("(", p2, ")")
		//fmt.Print(" = ")
		//fmt.Print(a.String())
		//fmt.Print(" -> ")
		//fmt.Print(i)
		//fmt.Print(" + ")
		//fmt.Print(i2)
		//fmt.Print(" = ")
		//fmt.Print(i2 + i)
		//fmt.Print(" => ")
		//fmt.Print(*total)
		//fmt.Println()

		return nil
	}
}

func getAction(a, b string) action {
	switch a {
	case "A":
		switch b {
		case "X":
			return draw
		case "Y":
			return lost
		case "Z":
			return won
		}
	case "B":
		switch b {
		case "X":
			return won
		case "Y":
			return draw
		case "Z":
			return lost
		}
	case "C":
		switch b {
		case "X":
			return lost
		case "Y":
			return won
		case "Z":
			return draw
		}
	}
	return lost
}

func printName(s string) {
	if s == "A" {
		fmt.Print("Rock")
	}
	if s == "B" {
		fmt.Print("Paper")
	}
	if s == "C" {
		fmt.Print("Scissors")
	}
	if s == "X" {
		fmt.Print("Rock")
	}
	if s == "Y" {
		fmt.Print("Paper")
	}
	if s == "Z" {
		fmt.Print("Scissors")
	}
}
