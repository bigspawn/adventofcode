package main

import (
	"fmt"
	"strings"

	"github.com/bigspawn/adventofcode-2022/day/reader"
)

var points = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var pointsWhenItAction = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

func main() {
	total := 0
	err := reader.Read(secondStrategy(&total))
	if err != nil {
		panic(err)
	}

	fmt.Println(total)
}

func secondStrategy(total *int) func(l string) error {
	return func(l string) error {
		if l == "" {
			return nil
		}

		arr := strings.Split(l, " ")
		if len(arr) != 2 {
			return fmt.Errorf("wrong line")
		}

		player1 := arr[0]
		whatIMustToDo := arr[1]

		c := getCodeForSecondPlayer(player1, whatIMustToDo)

		a1 := pointsWhenItAction[whatIMustToDo]
		a2 := points[c]

		*total += a1 + a2

		return nil
	}
}

func getCodeForSecondPlayer(a, b string) string {
	switch a {
	case "A":
		switch b {
		case "X":
			return "Z"
		case "Y":
			return "X"
		case "Z":
			return "Y"
		}
	case "B":
		switch b {
		case "X":
			return "X"
		case "Y":
			return "Y"
		case "Z":
			return "Z"
		}
	case "C":
		switch b {
		case "X":
			return "Y"
		case "Y":
			return "Z"
		case "Z":
			return "X"
		}
	}
	return ""
}
