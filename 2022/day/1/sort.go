package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/bigspawn/adventofcode-2022/day/reader"
)

func processSort() error {
	elfs := make([]int, 0)
	count := 1
	elfs = append(elfs, 0)

	reader.Read(func(l string) error {
		if l == "" {
			count++
			elfs = append(elfs, 0)
		} else {
			v, err := strconv.Atoi(l)
			if err != nil {
				return fmt.Errorf("failed to parse %q: %w", l, err)
			}
			elfs[count-1] += v
		}
		return nil
	})

	//fmt.Println("before", elfs)

	sort.Ints(elfs)

	fmt.Println(len(elfs))

	//fmt.Println("sorted", elfs)
	fmt.Println("total", elfs[len(elfs)-1])

	var total int
	for _, v := range elfs[len(elfs)-3:] {
		fmt.Print(v, " ")
		total += v
	}
	fmt.Println()
	fmt.Println("total", total)

	return nil
}
