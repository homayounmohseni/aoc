package main

import (
	"fmt"
	"testing"
)

func TestParseInput(t *testing.T) {
	var test = struct {
		in       []string
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

func TestAddIntRange(t *testing.T) {
	var tests = []struct {
		initialRanges       []IntRange
		newRanges           []IntRange
		expectedFinalRanges []IntRange
	}{
		{
			[]IntRange{},
			[]IntRange{{2, 7}},
			[]IntRange{
				{2, 7},
			},
		},
		{
			[]IntRange{},
			[]IntRange{{2, 2}},
			[]IntRange{{2, 2}},
		},
		{
			[]IntRange{{1, 2}, {5, 10}, {25, 25}},
			[]IntRange{{3, 4}},
			[]IntRange{{1, 2}, {3, 4}, {5, 10}, {25, 25}},
		},
		{
			[]IntRange{{1, 2}, {5, 10}, {25, 25}},
			[]IntRange{{3, 5}},
			[]IntRange{{1, 2}, {3, 10}, {25, 25}},
		},
		{
			[]IntRange{{1, 2}, {5, 10}, {25, 25}},
			[]IntRange{{2, 4}},
			[]IntRange{{1, 4}, {5, 10}, {25, 25}},
		},
		{
			[]IntRange{{1, 2}, {5, 10}, {25, 25}},
			[]IntRange{{4, 9}},
			[]IntRange{{1, 2}, {4, 10}, {25, 25}},
		},
		{
			[]IntRange{{1, 2}, {5, 10}, {25, 25}},
			[]IntRange{{6, 11}},
			[]IntRange{{1, 2}, {5, 11}, {25, 25}},
		},
		{
			[]IntRange{{1, 2}, {5, 10}, {25, 25}, {35, 40}},
			[]IntRange{{4, 26}},
			[]IntRange{{1, 2}, {4, 26}, {35, 40}},
		},
		{
			[]IntRange{},
			[]IntRange{{2, 2}, {4, 5}, {11, 13}},
			[]IntRange{{2, 2}, {4, 5}, {11, 13}},
		},
	}
	for _, test := range tests {
		finalRanges := AddIntRange(test.initialRanges, test.newRanges...)
		err := fmt.Errorf(
			"AddIntRange: initialRanges: %v, newRanges: %v, finalRanges: %v, expectedFinalRanges: %v\n",
			test.initialRanges,
			test.newRanges,
			finalRanges,
			test.expectedFinalRanges,
		)
		if len(finalRanges) != len(test.expectedFinalRanges) {
			t.Error(err)
		} else {
			for i := 0; i < len(finalRanges); i++ {
				if finalRanges[i] != test.expectedFinalRanges[i] {
					t.Error(err)
				}
			}
		}
	}
}

func TestComplementRanges(t *testing.T) {
	min := 0
	max := 20
	tests := []struct {
		ranges              []IntRange
		expectedFinalRanges []IntRange
	}{
		{
			[]IntRange{},
			[]IntRange{{0, 20}},
		},
		{
			[]IntRange{{0, 1}, {10, 11}, {19, 20}},
			[]IntRange{{1, 10}, {11, 19}},
		},
		{
			[]IntRange{{0, 20}},
			[]IntRange{},
		},
		{
			[]IntRange{{0, 25}},
			[]IntRange{},
		},
		{
			[]IntRange{{-5, 20}},
			[]IntRange{},
		},
		{
			[]IntRange{{-5, 25}},
			[]IntRange{},
		},
		{
			[]IntRange{{-5, 5}},
			[]IntRange{{5, 20}},
		},
		{
			[]IntRange{{15, 25}},
			[]IntRange{{0, 15}},
		},
		{
			[]IntRange{{-5, 5}, {7, 13}, {15, 25}},
			[]IntRange{{5, 7}, {13, 15}},
		},
	}

	for _, test := range tests {
		finalRanges := ComplementRanges(test.ranges, min, max)
		err := fmt.Errorf(
			"ComplementRanges: initialRanges: %v, min: %v, max: %v finalRanges: %v, expectedFinalRanges: %v\n",
			test.ranges, min, max, finalRanges, test.expectedFinalRanges,
		)
		if len(finalRanges) != len(test.expectedFinalRanges) {
			t.Error(err)
		} else {
			for i := 0; i < len(finalRanges); i++ {
				if finalRanges[i] != test.expectedFinalRanges[i] {
					t.Error(err)
				}
			}
		}
	}
}

func TestCountRangeElements(t *testing.T) {
	tests := []struct {
		ranges        []IntRange
		expectedCount int
	}{
		{
			[]IntRange{{1, 3}, {5, 10}},
			7,
		},
	}
	for _, test := range tests {
		count := CountRangeElements(test.ranges)
		if count != test.expectedCount {
			t.Errorf("CountRange: ranges: %v, count: %v, expectedCount: %v",
				test.ranges, count, test.expectedCount,
			)
		}
	}
}
