package main

import (
	"log"
	"bufio"
	"os"
	"fmt"
	"io"
	"errors"
	"strings"
)

type square struct {
	height byte
	seen bool
	distance int
}

type position struct {
	x int
	y int
}

type queue[T any] []T

type heightmap [][]square


func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	reader := bufio.NewReader(os.Stdin)

	hmSlice := make([][]square, 0)
	for {
		row := make([]square, 0)
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}
		for i := 0; i < len(line); i++ {
			row = append(row, square{height: line[i], distance: 1 << 31 - 1})
		}
		hmSlice = append(hmSlice, row)
	}

	var startPos, endPos position
	for i, row := range hmSlice {
		for j, sq := range row {
			if sq.height == 'S' {
				hmSlice[i][j].height = 'a'
				startPos = position{j, i}
			} else if sq.height == 'E' {
				hmSlice[i][j].height = 'z'
				endPos = position{j, i}
			}
		}
	}

	hm := heightmap(hmSlice)

	q := newQueue[position]()
	endSq := hm.at(endPos)
	endSq.seen = true
	endSq.distance = 0

	q.enqueue(endPos)
	for !q.isEmpty() {
		pos := q.dequeue()
		
		sq := hm.at(pos)
		adjs := []position{pos.up(), pos.down(), pos.left(), pos.right()}
		for _, adjPos := range adjs {
			if hm.contains(adjPos) {
				adjSq := hm.at(adjPos)
				if canUnclimb(sq, adjSq) && !adjSq.seen {
					adjSq.seen = true
					adjSq.distance = sq.distance + 1
					q.enqueue(adjPos)
				}
			}
		}
	}

	startToEndDistance := hm.at(startPos).distance
	var minLowestToEndDistance int = 1 << 31 - 1
	for _, row := range hm {
		for _, sq := range row {
			if sq.height == 'a' {
				if sq.distance < minLowestToEndDistance {
					minLowestToEndDistance = sq.distance
				}
			}
		}
	}
	fmt.Println(startToEndDistance)
	fmt.Println(minLowestToEndDistance)
}


func newQueue[T any]() *queue[T] {
	q := queue[T](make([]T, 0))
	return &q
}

func (q *queue[T]) enqueue(v T) {
	*q =append(*q, v)
}

func (q *queue[T]) dequeue() T {
	v := (*q)[0];
	*q = (*q)[1:];
	return v
}

func (q *queue[T]) isEmpty() bool {
	return len(*q) == 0
}

func (hm heightmap) at(pos position) *square {
	return &hm[pos.y][pos.x]
}

func (hm heightmap) contains(pos position) bool {
	if pos.y < 0 || pos.y >= len(hm) {
		return false
	}
	if pos.x < 0 || pos.x >= len(hm[0]) {
		return false
	}
	return true
}

func (hm heightmap) display() {
	for _, row := range hm {
		for _, sq := range row {
			var r rune = '.'
			if sq.seen {
				r = '#'
			}
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
}

func (pos position) up() position {
	return position{pos.x, pos.y + 1}
}

func (pos position) down() position {
	return position{pos.x, pos.y - 1}
}

func (pos position) right() position {
	return position{pos.x + 1, pos.y}
}

func (pos position) left() position {
	return position{pos.x - 1, pos.y}
}

func canClimb(from, to *square) bool {
	return int16(to.height) - int16(from.height) <= 1
}

func canUnclimb(from, to *square) bool {
	return int16(to.height) - int16(from.height) >= -1
}
