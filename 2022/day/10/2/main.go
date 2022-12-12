package main

import (
	"fmt"
	"os"

	"github.com/bigspawn/adventofcode-2022/day/reader"
)

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

	x := 1
	s := ""
	for j := 0; ; j++ {
		for i := 0; i < 40; i++ {
			if len(commands) == 0 {
				return
			}

			ls := len(s)

			if ls == x || ls == (x-1) || ls == (x+1) {
				s += "#"
			} else {
				s += "."
			}

			c := commands[0]

			if (i+1)/40 == 1 {
				fmt.Println(s)
				s = ""
			}

			x += c
			commands = commands[1:]
		}
	}
}
