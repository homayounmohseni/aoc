package main

import (
	"testing"
	"fmt"
)

func TestCompare(t *testing.T) {
	var tests = []struct {
		in_l string
		in_r string
		expected int
	}{
		{"1", "2", 1},
		{"2", "1", -1},
		{"1", "1", 0},
		{"[1]", "1", 0},
		{"[1,2]", "1", -1},
		{"[[]]", "[]", -1},
		{
			"[1,2,3]",
			"[1,3,3]",
			1,
		},
		{
			"[2,2,3]",
			"[1,3,3]",
			-1,
		},
	}
	for _, tt := range tests {
		if out := Compare(tt.in_l, tt.in_r); out != tt.expected {
			t.Errorf("Compare: l: %#v r: %#v got: %v expected: %v", tt.in_l, tt.in_r, out, tt.expected)
		}
	}
}
func TestListToSlice(t *testing.T) {
	var tests = []struct {
		in string
		expected []string
	}{
		{"[1]", []string{"1"}},
		{"[[1]]", []string{"[1]"}},
		{"[[1],2,3]", []string{"[1]", "2", "3"}},
		{"[1,2,[3]]", []string{"1", "2", "[3]"}},
		{"[1,[2,3,4],5,[[6]]]", []string{"1", "[2,3,4]", "5", "[[6]]"}},
		{"[1,[2,[3,[4,5]]]]", []string{"1", "[2,[3,[4,5]]]"}},
	}

	for _, tt := range tests {
		testname := fmt.Sprint(tt.in)
		t.Run(testname, func(t *testing.T) {
			out := ListToSlice(tt.in)
			errMessage := fmt.Sprintf("ListToSlice: got %v, expected %v", out, tt.expected)
			if len(tt.expected) != len(out) {
				t.Error(errMessage)
			} else {
				for i := 0; i < len(out); i++ {
					if tt.expected[i] != out[i] {
						t.Error(errMessage)
						break
					}
				}
			}
		})
	}
}
