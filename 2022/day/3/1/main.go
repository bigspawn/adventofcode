package main

import (
	"fmt"

	"github.com/bigspawn/adventofcode-2022/day/reader"
)

var total = 0

func main() {
	err := reader.Read(process())
	if err != nil {
		panic(err)
	}
	fmt.Println(total)
}

type set map[string]struct{}

func process() func(l string) error {
	return func(l string) error {
		set1, set2 := make(set), make(set)
		i1, i2 := 0, len(l)/2

		for i1 < len(l)/2 {
			s1 := l[i1 : i1+1]
			set1[s1] = struct{}{}
			//fmt.Println("1: ", s1)

			s2 := l[i2 : i2+1]
			set2[s2] = struct{}{}
			//fmt.Println("2: ", s2)

			i1++
			i2++
		}

		//fmt.Println(set1, set2)

		for k := range set1 {
			if _, ok := set2[k]; ok {
				r := []rune(k)
				i := mapRune(r[0])
				//fmt.Println("->", k, "=>", i)

				total += i
			}
		}

		return nil
	}
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
