package reader

import (
	"bufio"
	"fmt"
	"os"
)

func Read(apply func(string) error) error {
	//f, err := os.Open("./input_small.txt")
	f, err := os.Open("./input.txt")
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		l := scanner.Text()
		err = apply(l)
		if err != nil {
			return err
		}
	}

	return nil
}

func ReadFile(path string, apply func(string) error) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		l := scanner.Text()
		err = apply(l)
		if err != nil {
			return err
		}
	}

	return nil
}

func ReadN(n int, apply func([]string) error) error {
	//f, err := os.Open("./input_small.txt")
	f, err := os.Open("./input.txt")
	if err != nil {
		return fmt.Errorf("failed to open input file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	s := make([]string, 0, n)
	count := 0
	for scanner.Scan() {
		l := scanner.Text()

		count++

		if count <= n {
			s = append(s, l)
		}

		if count == n {
			if err = apply(s); err != nil {
				return err
			}
			count = 0
			s = make([]string, 0, n)
		}
	}

	return nil
}
