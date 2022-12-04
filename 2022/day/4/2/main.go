package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bigspawn/adventofcode-2022/day/reader"
)

type interval struct {
	left  int
	right int
}

func (i interval) Include(i2 interval) bool {
	return i.left <= i2.left && i.right >= i2.right
}

func (i interval) Overlap(i2 interval) bool {
	return (i.left <= i2.left) && (i.right >= i2.left)
}

func main() {
	args := os.Args[1:]
	filePath := args[0]

	count := 0
	err := reader.ReadFile(filePath, func(l string) error {
		s := strings.Split(l, ",")

		//fmt.Println(s)

		i1, i2, err := parseIntervals(s)
		if err != nil {
			return err
		}
		if i1.Overlap(i2) {
			fmt.Println(i1, i2)
			count++
		} else if i2.Overlap(i1) {
			fmt.Println(i2, i1)
			count++
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}

func parseIntervals(s []string) (interval, interval, error) {
	var i1, i2 interval
	var err error
	for i, ss := range s {
		v := strings.Split(ss, "-")

		//fmt.Println(v)

		var ii interval
		ii.left, err = strconv.Atoi(v[0])
		if err != nil {
			return interval{}, interval{}, err
		}
		ii.right, err = strconv.Atoi(v[1])
		if err != nil {
			return interval{}, interval{}, err
		}

		//fmt.Println(ii)

		if i == 0 {
			i1 = ii
		} else {
			i2 = ii
		}
	}
	return i1, i2, nil
}
