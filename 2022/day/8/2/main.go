package main

import (
	"fmt"
	"os"

	"github.com/bigspawn/adventofcode-2022/day/conv"
	"github.com/bigspawn/adventofcode-2022/day/reader"
)

func main() {
	args := os.Args[1:]
	filePath := args[0]

	var trees [][]int
	var w, h int
	err := reader.ReadFile(filePath, func(l string) error {
		w = len(l)
		arr := make([]int, 0, w)
		for _, r := range l {
			n := conv.RuneToInt(r)
			arr = append(arr, n)
		}
		trees = append(trees, arr)
		h++
		return nil
	})
	if err != nil {
		panic(err)
	}

	max := 0

	for j := 1; j < h-1; j++ {
		for i := 1; i < w-1; i++ {
			current := trees[j][i]
			c := checkW(trees, current, i, j, w)
			c *= checkH(trees, current, i, j, h)
			if max < c {
				max = c
			}
			//fmt.Println("=", c)
		}
	}

	fmt.Println(max)
}

func checkW(trees [][]int, current, i, j, w int) int {
	count := 0
	res := 0
	for x := i - 1; x >= 0; x-- {
		t := trees[j][x]
		count++
		//fmt.Printf("%d -> %d: %d\n", t, current, count)
		if t >= current {
			break
		}
	}
	res = count
	count = 0
	for x := i + 1; x < w; x++ {
		count++
		t := trees[j][x]
		//fmt.Printf("%d -> %d: %d\n", t, current, count)
		if t >= current {
			break
		}
	}

	return res * count
}

func checkH(trees [][]int, current, i, j, h int) int {
	count := 0
	res := 0
	for x := j - 1; x >= 0; x-- {
		t := trees[x][i]
		count++
		//fmt.Printf("%d -> %d: %d\n", t, current, count)
		if t >= current {
			break
		}
	}
	res = count
	count = 0
	for x := j + 1; x < h; x++ {
		t := trees[x][i]
		count++
		//fmt.Printf("%d -> %d: %d\n", t, current, count)
		if t >= current {
			break
		}
	}

	return res * count
}
