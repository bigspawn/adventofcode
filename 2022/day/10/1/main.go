package main

import (
	"fmt"
	"os"

	"github.com/bigspawn/adventofcode-2022/day/reader"
)

var cycleTimes = map[int]struct{}{20: {}, 60: {}, 100: {}, 140: {}, 180: {}, 220: {}}

func main() {
	args := os.Args[1:]
	filePath := args[0]

	var commands []int
	err := reader.ReadFile(filePath, func(l string) error {
		var cmd string
		var n int
		fmt.Sscanf(l, "%s %d", &cmd, &n)

		switch cmd {
		case "addx":
			commands = append(commands, 0, n)
		case "noop":
			commands = append(commands, 0)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	res := 0
	X := 1
	for i := 0; i < 220; i++ {
		cycle := i + 1
		//fmt.Println(cycle, X)
		if _, ok := cycleTimes[cycle]; ok {
			res += cycle * X
			fmt.Println("->", cycle, X, res)
		}
		X += commands[i]
	}

}
