package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Pair[T any] [2]T
type Position Pair[int]
type Path []Position

type Cave struct {
	m           map[Position]int8
	floorY      int
	floorExists bool
}

var NoMoreInts = errors.New("no more integers found in string")
var SandStartingPosition = Position{500, 0}

const (
	sand = iota + 1
	rock
	air
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	lines := ReadLines()
	paths := ParseInput(lines)

	cave := InitCave(paths)
	highestY := findDeepestRockY(cave)

	var caveFloored Cave
	CopyCave(&caveFloored, &cave)
	caveFloored.floorExists = true
	caveFloored.floorY = highestY + 2

	turns := simulate(cave, highestY)
	turnsFloored := simulateFloored(caveFloored)

	fmt.Println(turns)
	fmt.Println(turnsFloored)
}

func (c *Cave) Demap(pos Position) int8 {
	if c.floorExists && pos[1] == c.floorY {
		return rock
	}
	if v, ok := c.m[pos]; ok {
		return v
	}
	return air
}

func (c *Cave) Map(pos Position, v int8) {
	c.m[pos] = v
}

func CopyCave(dst, src *Cave) {
	dst.floorY = src.floorY
	dst.floorExists = src.floorExists
	dst.m = make(map[Position]int8, len(src.m))
	for k, v := range src.m {
		dst.m[k] = v
	}
}

func simulate(cave Cave, deepestRockY int) int {
	for i := 0; ; i++ {
		curPos := SandStartingPosition
		for {
			nextPos := calculateNextPosition(curPos, cave)

			if nextPos == curPos {
				break
			} else if nextPos[1] > deepestRockY {
				return i
			}
			cave.Map(curPos, air)
			cave.Map(nextPos, sand)
			curPos = nextPos
		}
	}
}

func simulateFloored(cave Cave) int {
	for i := 0; ; i++ {
		curPos := SandStartingPosition
		for {
			nextPos := calculateNextPosition(curPos, cave)
			if nextPos == curPos {
				if curPos == SandStartingPosition {
					return i + 1
				}
				break
			}
			cave.Map(curPos, air)
			cave.Map(nextPos, sand)
			curPos = nextPos
		}
	}
}

func calculateNextPosition(curPos Position, cave Cave) Position {
	possiblePositions := []Position{
		{curPos[0], curPos[1] + 1},
		{curPos[0] - 1, curPos[1] + 1},
		{curPos[0] + 1, curPos[1] + 1},
	}

	for _, pos := range possiblePositions {
		if cave.Demap(pos) == air {
			return pos
		}
	}

	return curPos
}

func InitCave(paths []Path) Cave {
	cave := Cave{
		m:           make(map[Position]int8),
	}
	for pathNum, path := range paths {
		var prev Position
		if len(path) != 0 {
			prev = path[0]
		}
		for i := 1; i < len(path); i++ {
			cur := path[i]
			if cur[0] == prev[0] {
				x := cur[0]
				from := MinInt(cur[1], prev[1])
				to := MaxInt(cur[1], prev[1])
				for y := from; y <= to; y++ {
					cave.Map(Position{x, y}, rock)
				}
			} else if cur[1] == prev[1] {
				y := cur[1]
				from := MinInt(cur[0], prev[0])
				to := MaxInt(cur[0], prev[0])
				for x := from; x <= to; x++ {
					cave.Map(Position{x, y}, rock)
				}
			} else {
				log.Fatal(fmt.Errorf("bad input usage: prev: %v, cur: %v, pathNum: %v, i: %v",
					prev, cur, pathNum, i))
			}
			prev = cur
		}
	}
	return cave
}

func findDeepestRockY(c Cave) int {
	maxY := 0
	for k, v := range c.m {
		if v == rock {
			maxY = MaxInt(maxY, k[1])
		}
	}
	return maxY
}

func ReadLines() []string {
	var lines []string
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			} else {
				log.Fatal(err)
			}
		}
		lines = append(lines, line)
	}
	return lines
}

func ParseInput(lines []string) []Path {
	var paths []Path
	for _, line := range lines {
		var path Path
		var position Position
		var coordinateIndex int8
		pos := 0
		for {
			var intRead int
			var err error
			intRead, pos, err = ReadOneInt(line, pos)
			if err != nil {
				if errors.Is(err, NoMoreInts) {
					break
				} else {
					log.Fatal(err)
				}
			}
			position[coordinateIndex] = intRead
			if coordinateIndex == 1 {
				path = append(path, position)
			}
			coordinateIndex = (coordinateIndex + 1) % 2
		}
		paths = append(paths, path)
	}
	return paths
}

func ReadOneInt(str string, pos int) (intRead, newPos int, err error) {
	var start int
	for start = pos; start < len(str) && (str[start] < byte('0') || str[start] > byte('9')); start++ {
	}

	var end int
	for end = start; end < len(str) && str[end] >= byte('0') && str[end] <= byte('9'); end++ {
	}
	if end == start {
		return 0, 0, NoMoreInts
	}
	i, atoiErr := strconv.Atoi(str[start:end])
	if atoiErr != nil {
		log.Fatal(atoiErr)
	}
	return i, end, nil
}

func MinInt(s ...int) int {
	if len(s) == 0 {
		return -1
	}
	min := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}
	return min
}

func MaxInt(s ...int) int {
	if len(s) == 0 {
		return -1
	}
	max := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}
