package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestReadOneInt(t *testing.T) {
	var tests = []struct {
		strIn    string
		expected []int
	}{
		{"0", []int{0}},
		{" 9", []int{9}},
		{"9 ", []int{9}},
		{" 5 ", []int{5}},
		{"fdsafd0_1 23 456v 78 98fdsad7", []int{0, 1, 23, 456, 78, 98, 7}},
	}

	for _, test := range tests {
		var out []int
		var pos int
		for {
			var intRead int
			var err error
			intRead, pos, err = ReadOneInt(test.strIn, pos)
			if err != nil {
				if errors.Is(err, NoMoreInts) {
					break
				} else {
					t.Errorf("ReadOneInt: in %v, got error %v, expected %v\n",
						test.strIn, err, test.expected)
					break
				}
			}
			out = append(out, intRead)
		}
		errMessage := fmt.Sprintf("ReadOneInt: in %v, got %v, expected, %v\n",
			test.strIn, out, test.expected)
		if len(out) != len(test.expected) {
			t.Error(errMessage)
		} else {
			for i := 0; i < len(out); i++ {
				if out[i] != test.expected[i] {
					t.Error(errMessage)
				}
			}
		}
	}
}

func TestMinInt(t *testing.T) {
	var tests = []struct {
		in       []int
		expected int
	}{
		{[]int{}, -1},
		{[]int{2}, 2},
		{[]int{1, 2, -1, 7}, -1},
	}
	for _, test := range tests {
		out := MinInt(test.in...)
		if out != test.expected {
			t.Errorf("MinInt: in %v, got %v, expected %v\n",
				test.in,
				out,
				test.expected,
			)
		}
	}
}

func TestMaxInt(t *testing.T) {
	var tests = []struct {
		in       []int
		expected int
	}{
		{[]int{}, -1},
		{[]int{2}, 2},
		{[]int{1, 2, -1, 7}, 7},
	}
	for _, test := range tests {
		out := MaxInt(test.in...)
		if out != test.expected {
			t.Errorf("MaxInt: in %v, got %v, expected %v\n",
				test.in,
				out,
				test.expected,
			)
		}
	}
}
