package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestName(t *testing.T) {
	s := []string{
		"/",
		"/a",
		"/a/b",
		"/b",
		"/b/c",
		"/b/c/c",
		"/b/c/d",
		"/w",
	}
	sort.Strings(s)
	fmt.Println(s)
}
