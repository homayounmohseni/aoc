package main

import (
	"fmt"
	"testing"
)

func TestParseInput(t *testing.T) {
	var test = struct{
		in []string
		expected []Pair[Position]
	}{
		[]string{
			"Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
			"Sensor at x=9, y=16: closest beacon is at x=10, y=16",
			"Sensor at x=13, y=2: closest beacon is at x=15, y=3",
		},
		[]Pair[Position]{
			{{2, 18}, {-2, 15}},
			{{9, 16}, {10, 16}},
			{{13, 2}, {15, 3}},
		},
	}
	out := ParseInput(test.in)
	if len(out) != len(test.expected) {
		t.Errorf(
			"ParseInput: expected %v pairs got %v pairs.\nin %v, got %v, expected %v\n",
			len(test.expected), len(out), test.in, out, test.expected,
		)
	}
	for i, pair := range out {
		for j, pos := range pair {
			if pos != test.expected[i][j] {
				t.Errorf(
					"ParseInput: single line in %v, got %v, expected %v\n",
					test.in[i], out[i], test.expected[i],
				)
			}
		}
	}
}

func TestAddRange(t *testing.T) {
	var tests = []struct{
		fieldRanges []Pair[int]
		newRange Pair[int]
		expectedRanges []Pair[int]
	}{
		{
			[]Pair[int]{},
			Pair[int]{2, 7},
			[]Pair[int]{
				{2, 7},
			},
		},
		{
			[]Pair[int]{},
			Pair[int]{2, 2},
			[]Pair[int]{{2, 2}},
		},
		{
			[]Pair[int]{{1, 2}, {5, 10}, {25, 25}},
			Pair[int]{3, 4},
			[]Pair[int]{{1, 2}, {3, 4}, {5, 10}, {25, 25}},
		},
		{
			[]Pair[int]{{1, 2}, {5, 10}, {25, 25}},
			Pair[int]{3, 5},
			[]Pair[int]{{1, 2}, {3, 10}, {25, 25}},
		},
		{
			[]Pair[int]{{1, 2}, {5, 10}, {25, 25}},
			Pair[int]{2, 4},
			[]Pair[int]{{1, 4}, {5, 10}, {25, 25}},
		},
		{
			[]Pair[int]{{1, 2}, {5, 10}, {25, 25}},
			Pair[int]{4, 9},
			[]Pair[int]{{1, 2}, {4, 10}, {25, 25}},
		},
		{
			[]Pair[int]{{1, 2}, {5, 10}, {25, 25}},
			Pair[int]{6, 11},
			[]Pair[int]{{1, 2}, {5, 11}, {25, 25}},
		},
		{
			[]Pair[int]{{1, 2}, {5, 10}, {25, 25}, {35, 40}},
			Pair[int]{4, 26},
			[]Pair[int]{{1, 2}, {4, 26}, {35, 40}},
		},
	}
	for _, test := range tests {
		f := Field{
			ranges: test.fieldRanges,
		}
		f.AddRange(test.newRange)
		err := fmt.Errorf(
			"AddRange: fieldRanges: %v, newRange: %v, got: %v, expected: %v\n",
			test.fieldRanges,
			test.newRange,
			f.ranges,
			test.expectedRanges,
		)

		if len(f.ranges) != len(test.expectedRanges) {
			t.Error(err)
		} else {
			for i := 0; i < len(f.ranges); i++ {
				if f.ranges[i] != test.expectedRanges[i]{
					t.Error(err)
				}
			}
		}
	}
}
