package main

import (
	"fmt"

	"github.com/bigspawn/adventofcode-2022/day/reader"
)

var total = 0

func main() {
	err := reader.ReadN(3, process)
	if err != nil {
		panic(err)
	}
	fmt.Println(total)
}

type set map[rune]struct{}

func process(arr []string) error {
	set1, set2, set3 := make(set), make(set), make(set)
	for i, s := range arr {
		for _, ss := range s {
			switch i {
			case 0:
				set1[ss] = struct{}{}
			case 1:
				set2[ss] = struct{}{}
			case 2:
				set3[ss] = struct{}{}
			}
		}
	}
	for s := range set1 {
		if _, ok := set2[s]; ok {
			if _, ok2 := set3[s]; ok2 {
				total += mapRune(s)
			}
		}
	}
	return nil
}

func mapRune(r rune) int {
	if r >= 97 && r <= 122 {
		return int(r) - 96
	}
	if r >= 65 && r <= 90 {
		return int(r) - 38
	}
	return 0
}
