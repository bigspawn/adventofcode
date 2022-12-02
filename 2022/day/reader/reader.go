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
