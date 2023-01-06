package main

import (
	"bufio"
	"log"
	"errors"
	"os"
	"strings"
	"strconv"
	"io"
	"fmt"
)

type pair[T any] [2]T

func main() {
	reader := bufio.NewReader(os.Stdin)
	pairs := make([]pair[string], 0)
	var p pair[string] 

	var interpairIndex int8
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				log.Fatal(err)
			}
		}

		line = strings.TrimSpace(line)
		if line != "" {
			p[interpairIndex] = line
			if interpairIndex == 1 {
				pairs = append(pairs, p)
			}
			interpairIndex = (interpairIndex + 1) % 2
		}
	}

	indexSum := 0;
	for i, p := range pairs {
		if (isInRightOrder(p)) {
			indexSum += i + 1
		}
	}
	
	fmt.Println(indexSum)
}

func isInRightOrder(p pair[string]) bool {
	return Compare(p[0], p[1]) >= 0
}

// a positive return value means the right is greater than the left
// a negative return value means the right is less than the left
// an equeal reuturn value means left and right are the same
func Compare(l, r string) int {
	if isList(l) && isList(r) {
		lSlice := ListToSlice(l)
		rSlice := ListToSlice(r)
		minLength := MinInt(len(lSlice), len(rSlice))
		
		for i := 0; i < minLength; i++ {
			cValue := Compare(lSlice[i], rSlice[i])
			switch {
			case cValue > 0:
				return 1
			case cValue < 0:
				return -1
			}
		}
		if len(lSlice) < len(rSlice) {
			return 1
		}
		if len(lSlice) > len(rSlice) {
			return -1
		}
		return 0
	} else if !isList(l) && !isList(r) {
		lInt, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}
		rInt, err := strconv.Atoi(r)
		if err != nil {
			log.Fatal(err)
		}
		switch {
		case lInt < rInt:
			return 1
		case lInt > rInt:
			return -1
		}
		return 0
	} else {
		switch {
		case !isList(l):
			l = listify(l)
		case !isList(r):
			r = listify(r)
		}
		return Compare(l, r)
	}
}


func isList(v string) bool {
	return strings.HasPrefix(v, "[") && strings.HasSuffix(v, "]")
}

func listify(v string) string {
	v = "[" + v + "]"
	return v
}

func ListToSlice(v string) []string {
	v = strings.TrimPrefix(v, "[")
	v = strings.TrimSuffix(v, "]")

	var slice []string
	startIndex := 0
	var i int
	for i = 0; i < len(v); i++ {
		switch v[i] {
		case '[':
			i++
			for opens := 1; opens > 0 && i < len(v); i++ {
				switch v[i] {
				case '[':
					opens++
				case ']':
					opens--
				}
			}
			i--

		case ',':
			slice = append(slice, v[startIndex:i])
			startIndex = i + 1
		}
	}

	if len(v) > 0 {
		slice = append(slice, v[startIndex:i])
	}
	return slice
}

func MinInt(a, b int) int {
	if (a < b) {
		return a
	}
	return b
}
