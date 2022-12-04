package main

import (
	"fmt"
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

func main() {
	count := 0
	err := reader.Read(func(l string) error {
		s := strings.Split(l, ",")

		fmt.Println(s)

		i1, i2, err := parseIntervals(s)
		if err != nil {
			return err
		}
		if i1.Include(i2) {
			count++
		} else if i2.Include(i1) {
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

		fmt.Println(v)

		var ii interval
		ii.left, err = strconv.Atoi(v[0])
		if err != nil {
			return interval{}, interval{}, err
		}
		ii.right, err = strconv.Atoi(v[1])
		if err != nil {
			return interval{}, interval{}, err
		}

		fmt.Println(ii)

		if i == 0 {
			i1 = ii
		} else {
			i2 = ii
		}
	}
	return i1, i2, nil
}
