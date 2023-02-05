package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pair[T any] [2]T
type Position Pair[int]

type SensorRecord struct {
	pos Position
	// distClosestBeacon int
	detectedBeaconPos Position
}

type Field struct {
	row int
	ranges []Pair[int]
}

type Node[T any] struct {
	next *Node[T]
	value T
}

const defaultRow = 2000000
// const defaultRow = 10

func main() {
	lines := ReadLines()
	sbPairs := ParseInput(lines)
	records := ExtractRecords(sbPairs)
	field := Field{row: defaultRow}
	field.InitRanges(records)
	
	cnt := field.CountPositions(records)
	fmt.Println(cnt)
}

func (f *Field) CountPositions(records []SensorRecord) int {
	count := 0
	for _, r := range f.ranges {
		count += r[1] - r[0]
	}
	return count
}

func (f *Field) AddRange(newR Pair[int]) {
	if len(f.ranges) == 0 {
		f.ranges = append(f.ranges, newR)
		return
	}
	
	i := 0
	for ; i < len(f.ranges) && f.ranges[i][1] < newR[0]; i++ {}
	j := i
	for ; j < len(f.ranges) && f.ranges[j][0] <= newR[1]; j++ {}
	var newRanges []Pair[int]
	newRanges = append(newRanges, f.ranges[:i]...)
	//assert j >= i TODO
	var mergedRange Pair[int]
	if i != j {
		mergedRange = Pair[int]{MinInt(newR[0], f.ranges[i][0]), MaxInt(newR[1], f.ranges[j - 1][1])}
	} else {
		mergedRange = newR
	}
	newRanges = append(newRanges, mergedRange)
	newRanges = append(newRanges, f.ranges[j:]...)
	f.ranges = newRanges
}

func (f *Field) InitRanges(records []SensorRecord) {
	for _, r := range records {
		dist := GetManhattanDistance(r.pos, r.detectedBeaconPos)
		l := dist - AbsInt(r.pos[1] - f.row)
		if l <= 0 {
			continue
		}
		rng := Pair[int]{r.pos[0] - l, r.pos[0] + l}
		f.AddRange(rng)
	}
}

func ExtractRecords(sbPairs []Pair[Position]) []SensorRecord {
	var records []SensorRecord
	for _, sbPair := range sbPairs {
		records = append(records, SensorRecord{
			pos: sbPair[0],
			// distClosestBeacon: GetManhattanDistance(sbPair[0], sbPair[1]),
			detectedBeaconPos: sbPair[1],
		})
	}
	return records
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

func ParseInput(lines []string) []Pair[Position] {
	pairs := make([]Pair[Position], 0, len(lines))
	for _, line := range lines {
		ints := make([]int, 0, 4)
		for i := 0; i < 4; i++ {
			from := strings.Index(line, "=")
			from++
			to := from + 1
			for ; to < len(line) && line[to] != ',' && line[to] != ':'; to++ {}
			num, err := strconv.Atoi(strings.TrimSpace(line[from:to]))
			if err != nil {
				panic(fmt.Errorf("ParseInput: %v", err))
			}
			ints = append(ints, num)
			line = line[to:]
		}
		pair := Pair[Position]{{ints[0], ints[1]}, {ints[2], ints[3]}}
		pairs = append(pairs, pair)
	}
	return pairs
}

func GetManhattanDistance(p1, p2 Position) int {
	var dist int
	for i := 0; i < 2; i++ {
		addedDist := AbsInt(p1[i] - p2[i])
		dist += addedDist
	}
	return dist
}

func AbsInt(s int) int {
	if s < 0 {
		s = -s
	}
	return s
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
