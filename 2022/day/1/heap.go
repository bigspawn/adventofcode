package main

import (
	"container/heap"
	"fmt"
	"strconv"

	"github.com/bigspawn/adventofcode-2022/day/reader"
)

type IntHeap []int

func (h *IntHeap) Len() int           { return len(*h) }
func (h *IntHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *IntHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap struct {
	IntHeap
}

func (h MaxHeap) Less(i, j int) bool { return h.IntHeap[i] > h.IntHeap[j] }
func processHeap() error {
	v := 0
	h := &MaxHeap{}
	err := reader.Read(func(l string) error {
		if l != "" {
			vv, err := strconv.Atoi(l)
			if err != nil {
				return fmt.Errorf("failed to parse %q: %w", l, err)
			}
			v += vv
		} else {
			heap.Push(h, v)
			v = 0
		}
		return nil
	})
	if err != nil {
		return err
	}

	max := heap.Pop(h).(int)
	fmt.Println("max", max)
	a := max + heap.Pop(h).(int) + heap.Pop(h).(int)
	fmt.Println("total", a)

	fmt.Println()

	return nil
}
