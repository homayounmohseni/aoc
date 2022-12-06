package main

import (
	"fmt"
)

const vsize = 14;

func main() {
	var line string
	fmt.Scanln(&line);
	
	var prevChars [vsize]rune
	for i := 0; i < vsize; i++ {
		prevChars[i] = 0;
	}

	for ind, c := range line {
		prevChars = shiftRight(prevChars, c);
		if ind > 3 && isVectorOK(prevChars) {
			fmt.Println(ind + 1)
			return
		}
	}
}

func shiftRight(charVector [vsize]rune, sin rune) [vsize]rune {
	for i := 0; i < vsize - 1; i++ {
		charVector[i] = charVector[i + 1];
	}
	charVector[vsize - 1] = sin;
	return charVector
}

func isVectorOK(charVector [vsize]rune) bool {
	m := make(map[rune]bool, vsize)

	for _, v := range charVector {
		if m[v] {
			return false;
		}
		m[v] = true;
	}
	return true;
}


		
