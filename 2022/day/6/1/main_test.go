package main

import "testing"

func Test_findMarker(t *testing.T) {
	tests := map[string]struct {
		r string
		i int
	}{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb": {
			r: "jpqm",
			i: 7,
		},
		"bvwbjplbgvbhsrlpgdmjqwftvncz": {
			r: "vwbj",
			i: 5,
		},
		"nppdvjthqldpwncqszvftbrmjlhg": {
			r: "pdvj",
			i: 6,
		},
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": {
			r: "rfnt",
			i: 10,
		},
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw": {
			r: "zqfr",
			i: 11,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			r, i := findMarker(name)
			if r != tt.r {
				t.Fatal(r, tt.r)
			} else if i != tt.i {
				t.Fatal(i, tt.i)
			}
		})
	}
}
