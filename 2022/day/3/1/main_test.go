package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	s := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	for _, ss := range strings.Split(s, "\n") {
		err := process()(ss)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestName2(t *testing.T) {
	fmt.Print('A' == 65)
}
