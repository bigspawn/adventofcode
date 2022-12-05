package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/bigspawn/adventofcode-2022/day/reader"
)

func main() {
	args := os.Args[1:]
	filePath := args[0]

	actions := make(actions, 0, 3)
	stacks := make(rows, 0, 3)
	startsMove := false
	err := reader.ReadFile(filePath, func(l string) error {
		if strings.HasPrefix(l, " 1") {
			return nil
		}
		if l == "" {
			startsMove = true
			return nil
		}
		if !startsMove {
			stack := parser(l)
			stacks = append(stacks, stack)
			return nil
		}

		action, err := parseAction(l)
		if err != nil {
			return err
		}
		actions = append(actions, action)

		return nil
	})
	if err != nil {
		panic(err)
	}

	//fmt.Println(stacks)

	stacks2 := rowsToCells(stacks)
	//fmt.Println(stacks2)

	//fmt.Println(actions)

	for _, a := range actions {
		stacks2.act(a)
		//fmt.Println(stacks2)
	}

	stacks2.top()
}

func parseAction(l string) (action, error) {
	reg := regexp.MustCompile("^move\\s(\\d+)\\sfrom\\s(\\d+)\\sto\\s(\\d+)$")
	arrs := reg.FindAllStringSubmatch(l, 1)
	return action{
		numb: mustAtoi(arrs[0][1]),
		from: mustAtoi(arrs[0][2]) - 1,
		to:   mustAtoi(arrs[0][3]) - 1,
	}, nil
}

func mustAtoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

type action struct {
	numb int
	from int
	to   int
}

type actions []action

// [] [] []
type row []string

// [] [] []
// [] [] []
type rows []row

type cell []string

func pop(c cell) (string, cell) {
	var x string
	for {
		if len(c) == 0 {
			return "", c
		}
		x, c = c[len(c)-1], c[:len(c)-1]
		if x == EMPTY {
			continue
		} else {
			return x, c
		}
	}
}

func appendd(c cell, v string) cell {
	if len(c) == 0 {
		c = append(c, v)
		return c
	}
	count := 0
	i := len(c) - 1
	for ; i > 0; i-- {
		if c[i] == EMPTY {
			count++
			continue
		}
		break
	}
	if count != 0 {
		c = c[:len(c)-count]
		//fmt.Println(c)
	}
	c = append(c, v)
	return c
}

func (c cell) revers() cell {
	cc := make(cell, 0)
	for i := len(c) - 1; i >= 0; i-- {
		cc = append(cc, c[i])
	}
	return cc
}

type cells []cell

func (c cells) act(a action) {
	//fmt.Println("---<", a)
	for i := a.numb; i > 0; i-- {
		var v string
		v, c[a.from] = pop(c[a.from])
		c[a.to] = appendd(c[a.to], v)
	}
}

func (c cells) top() {
	for _, cc := range c {
		v, _ := pop(cc)
		if v == EMPTY || v == "" {
			continue
		}
		fmt.Print(v[1:2])
	}
	fmt.Println()
}

func rowsToCells(r rows) cells {
	i2 := 2 * len(r)
	ccs := make(cells, i2)
	for i := 0; i < i2; i++ {
		ccs[i] = make(cell, 0, len(r))
	}
	for _, rr := range r {
		for j, rrr := range rr {
			ccs[j] = append(ccs[j], rrr)
		}
	}
	for i := range ccs {
		ccs[i] = ccs[i].revers()
	}
	return ccs
}

const EMPTY = "   "

func parser(l string) row {
	count := int(math.Round(float64(len(l) / 3)))
	row := make(row, 0, count)
	for i := 0; i < len(l); {
		v := l[i : i+3]
		row = append(row, v)
		i += 4
	}
	return row
}
