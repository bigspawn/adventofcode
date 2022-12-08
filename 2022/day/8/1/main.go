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

	count := 0
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

	//fmt.Println(trees)

	for j := 1; j < h-1; j++ {
		for i := 1; i < w-1; i++ {
			current := trees[j][i]

			if checkW(trees, current, i, j, w) || checkH(trees, current, i, j, h) {
				//fmt.Println("\n", current, j, i)
				count++
				continue
			}

		}
	}

	count = 2*(w+h-2) + count

	fmt.Println(count)
}

func checkW(trees [][]int, current, i, j, w int) bool {
	visibleLeft := true
	for x := 0; x < i; x++ {
		t := trees[j][x]
		//fmt.Printf("[%d,%d]=%d", x, j, t)
		if t >= current {
			visibleLeft = false
			break
		}
	}

	visibleRight := true
	for x := i + 1; x < w; x++ {
		t := trees[j][x]
		if t >= current {
			visibleRight = false
			break
		}
	}

	//fmt.Print(" visibleLeft=", visibleLeft, " visibleRight=", visibleRight)

	return visibleLeft || visibleRight
}

func checkH(trees [][]int, current, i, j, h int) bool {
	visibleUp := true
	for x := 0; x < j; x++ {
		t := trees[x][i]
		if t >= current {
			visibleUp = false
			break
		}
	}

	visibleDown := true
	for x := j + 1; x < h; x++ {
		t := trees[x][i]
		if t >= current {
			visibleDown = false
			break
		}
	}

	//fmt.Print(" visibleUp=", visibleUp, " visibleDown=", visibleDown)

	return visibleUp || visibleDown
}
