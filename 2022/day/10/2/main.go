package main

import (
	"fmt"
	"os"

	"github.com/bigspawn/adventofcode-2022/day/reader"
)

func main() {
	args := os.Args[1:]
	filePath := args[0]

	count := 0
	err := reader.ReadFile(filePath, func(l string) error {
		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(count)
}
