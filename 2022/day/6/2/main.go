package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/bigspawn/adventofcode-2022/day/reader"
)

func main() {
	args := os.Args[1:]
	filePath := args[0]

	err := reader.ReadFile(filePath, func(l string) error {
		_, i := findMarker(l, 14)
		fmt.Println(i)
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func findMarker(s string, length int) (string, int) {
	res := process(0, s, "", length)
	i := strings.Index(s, res)
	return res, i + len(res)
}

func process(start int, s, res string, length int) string {
	//fmt.Println("start=", start, ", s=", s, ", res=", res)
	if len(res) == length {
		return res
	}
	if start == len(s) {
		return res
	}
	index := strings.IndexByte(res, s[start])
	if index < 0 {
		return process(start+1, s, res+string(s[start]), length)
	} else {
		return process(0, s[index+1:], string(s[index]), length)
	}
}
