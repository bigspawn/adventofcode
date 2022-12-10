package main

import (
	"fmt"
	"os"

	"github.com/bigspawn/adventofcode-2022/day/reader"
)

func main() {
	args := os.Args[1:]
	filePath := args[0]

	var r rope
	r.n = 2
	r.points = make([]position, r.n)
	r.path = map[position]struct{}{position{}: {}}
	err := reader.ReadFile(filePath, func(l string) error {
		var dir byte
		var n int
		fmt.Sscanf(l, "%c %d", &dir, &n)
		for i := 0; i < n; i++ {
			r.moveHead(string(dir))
			r.moveTail()
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	r.count()
}

type position struct {
	x, y int
}

type rope struct {
	n      int
	points []position
	path   map[position]struct{}
}

func (r *rope) count() {
	fmt.Println(len(r.path))
}

func (r *rope) moveHead(s string) {
	switch s {
	case "U":
		r.points[0].y++
	case "D":
		r.points[0].y--
	case "L":
		r.points[0].x--
	case "R":
		r.points[0].x++
	}
}

func (r *rope) moveTail() {
	for i := 1; i < r.n; i++ {
		delta := position{r.points[i-1].x - r.points[i].x, r.points[i-1].y - r.points[i].y}
		if abs(delta.x) <= 1 && abs(delta.y) <= 1 {
			return
		}
		if delta.y > 0 {
			r.points[i].y++
		} else if delta.y < 0 {
			r.points[i].y--
		}
		if delta.x > 0 {
			r.points[i].x++
		} else if delta.x < 0 {
			r.points[i].x--
		}
	}
	r.path[r.points[r.n-1]] = struct{}{}
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
