package main

import (
	"reflect"
	"testing"
)

func Test_parser(t *testing.T) {
	tests := []struct {
		name string
		l    string
		want row
	}{
		{
			name: "    [D]    ",
			l:    "    [D]    ",
			want: row{"   ", "[D]", "   "},
		},
		{
			name: "[N] [C]    ",
			l:    "[N] [C]    ",
			want: row{"[N]", "[C]", "   "},
		},
		{
			name: "    [N] [C]",
			l:    "    [N] [C]",
			want: row{"   ", "[N]", "[C]"},
		},
		{
			name: "[A] [N] [C]",
			l:    "[A] [N] [C]",
			want: row{"[A]", "[N]", "[C]"},
		},
		{
			name: "            ",
			l:    "            ",
			want: row{"   ", "   ", "   "},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parser(tt.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rowsToCells(t *testing.T) {
	tests := []struct {
		name string
		r    rows
		want cells
	}{
		{
			name: "",
			r: rows{
				row{"1", "2"},
				row{"1", "2"},
				row{"1", "2"},
			},
			want: cells{
				cell{"1", "1", "1"},
				cell{"2", "2", "2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rowsToCells(tt.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rowsToCells() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestName(t *testing.T) {
	parseAction("move 10 from 4 to 7")
}
