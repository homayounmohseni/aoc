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
// An InRange is defined as a pair of integers such that the second element is always larger
// than the first one. This range includes the numbers starting from r[0] to r[1] - 1 that is
// r[1] is not included in the range
type IntRange Pair[int]

type SensorRecord struct {
	pos Position
	detectedBeaconPos Position
}

type Field struct {
	row int
	xRanges map[int][]IntRange
	locatedBeacons map[Position]struct{}
}

type Node[T any] struct {
	next *Node[T]
	value T
}

// const defaultRow = 10
// const minY = 0
// const minX = 0
// const maxY = 20 + 1
// const maxX = 20 + 1
const defaultRow = 2_000_000
const minY = 0
const minX = 0
const maxY = 4_000_000 + 1
const maxX = 4_000_000 + 1

func main() {
	lines := ReadLines()
	sbPairs := ParseInput(lines)
	records := ExtractRecords(sbPairs)
	field := NewField()
	field.InitRanges(records)
	
	cnt := field.CountPositions(records)
	distressBeaconPos, err := field.FindDistressBeacon()
	distressBeaconFrequency := CalculateTuningFrequency(distressBeaconPos)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(cnt)
	fmt.Println(distressBeaconFrequency)
}

func NewField() *Field {
	return &Field{
		row: defaultRow,
		xRanges: make(map[int][]IntRange),
		locatedBeacons: make(map[Position]struct{}),
	}
}

func CalculateTuningFrequency(p Position) int64 {
	return int64(p[0]) * int64(4_000_000) + int64(p[1])
}

func (f *Field) FindDistressBeacon() (Position, error) {
	for y := 0; y < maxY; y++ {
		compRanges := ComplementRanges(f.xRanges[y], minX, maxX)
		if CountRangeElements(compRanges) == 1 {
			pos := Position{compRanges[0][0], y}
			if _, ok := f.locatedBeacons[pos]; !ok {
				return Position{compRanges[0][0], y}, nil
			}
		}
	}
	return Position{}, errors.New("distress beacon not found")
}

func (f *Field) CountPositions(records []SensorRecord) int {
	return CountRangeElements(f.xRanges[f.row])
}

func (f *Field) AddRange(newR IntRange, y int) {
	f.xRanges[y] = AddIntRange(f.xRanges[y], newR)
}

func (f *Field) InitRanges(records []SensorRecord) {
	for y := minY; y < maxY; y++ {
		for _, r := range records {
			f.locatedBeacons[r.detectedBeaconPos] = struct{}{}

			dist := GetManhattanDistance(r.pos, r.detectedBeaconPos)
			l := dist - AbsInt(r.pos[1] - y)
			if l <= 0 {
				continue
			}
			rngStart := r.pos[0] - l
			rngEnd := r.pos[0] + l + 1
			if r.detectedBeaconPos[1] == y {
				if r.detectedBeaconPos[0] == r.pos[0] - l {
					rngStart++;
				} else if r.detectedBeaconPos[0] == r.pos[0] + l {
					rngEnd--;
				}
			}
			rng := IntRange{rngStart, rngEnd}
			f.AddRange(rng, y)
		}
	}
}

func AddIntRange (slice []IntRange, elms ...IntRange) []IntRange{
	for _, elm := range elms {
		if len(slice) == 0 {
			slice = append(slice, elm)
			continue
		}
		
		var i int
		for i = 0; i < len(slice) && slice[i][1] < elm[0]; i++ {}
		var j int
		for j = i; j < len(slice) && slice[j][0] <= elm[1]; j++ {}
		var newRanges []IntRange
		newRanges = append(newRanges, slice[:i]...)
		var mergedRange IntRange
		if i != j {
			mergedRange = IntRange{MinInt(elm[0], slice[i][0]), MaxInt(elm[1], slice[j - 1][1])}
		} else {
			mergedRange = elm
		}
		newRanges = append(newRanges, mergedRange)
		newRanges = append(newRanges, slice[j:]...)
		slice = newRanges
	}
	return slice
}

func ComplementRanges(ranges []IntRange, min, max int) []IntRange {
	var compRanges []IntRange
	var i, j int
	for i = 0; i < len(ranges) && ranges[i][0] <= min; i++ {}
	for j = i; j < len(ranges) && ranges[j][1] < max; j++ {}

	var start int
	if i == 0 {
		start = min
	} else {
		start = MaxInt(min, ranges[i - 1][1])
		if start > max {
			start = max
		}
	}
	for k := i; k < j; k++ {
		compRanges = append(compRanges, IntRange{start, ranges[k][0]})
		start = ranges[k][1]
	}
	var end int
	if j == len(ranges) {
		end = max
	} else {
		end = MinInt(max, ranges[j][0])
	}
	if start != end {
		compRanges = append(compRanges, IntRange{start, end})
	}
	return compRanges
}

func CountRangeElements(ranges []IntRange) int {
	count := 0
	for _, r := range ranges {
		count += r[1] - r[0]
	}
	return count
}

func ExtractRecords(sbPairs []Pair[Position]) []SensorRecord {
	var records []SensorRecord
	for _, sbPair := range sbPairs {
		records = append(records, SensorRecord{
			pos: sbPair[0],
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
