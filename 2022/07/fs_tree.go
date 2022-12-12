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

type node struct {
	IsDir bool
	Links map[string]*node
	Size  int
}

const DirSizeLimit = 100000
const TotalSpace = 70000000
const UnusedSpaceNeeded = 30000000
const MaximumUsedSpace = TotalSpace - UnusedSpaceNeeded

func main() {
	root := newDir(nil)
	root.Links = make(map[string]*node)

	var pwd *node

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}

		words := strings.Fields(line)
		if len(words) == 0 {
			continue
		}
		switch words[0] {
		case "$":
			//maybe check words' length
			switch words[1] {
			case "ls":
			case "cd":
				switch newDir := words[2]; newDir {
				case "/":
					pwd = root
				default:
					pwd = pwd.Links[newDir]
				}
			default:
				panic("unsupported command")
			}
		default:
			switch words[0] {
			case "dir":
				pwd.Links[words[1]] = newDir(pwd)
			default:
				size, _ := strconv.Atoi(words[0])
				pwd.Links[words[1]] = &node{IsDir: false, Size: size}
			}
		}
	}

	_, acceptedDirsSize := calculateSizes(root)
	fmt.Println(acceptedDirsSize)

	usedSpace := calculateSize(root)
	_, selectedDirSize, err := findDirToDelete(root, usedSpace-MaximumUsedSpace)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(selectedDirSize)
}

func newDir(parent *node) *node {
	dir := &node{IsDir: true, Size: 0}
	dir.Links = make(map[string]*node)

	dir.Links[".."] = parent
	return dir
}

func calculateSize(topDir *node) int {
	var size int
	for k, v := range topDir.Links {
		if k == ".." {
			continue
		}
		if v.IsDir {
			size += calculateSize(v)
		} else {
			size += v.Size
		}
	}
	return size
}

func calculateSizes(topDir *node) (int, int) {
	var thisSize, totalSize int
	for k, v := range topDir.Links {
		if k == ".." {
			continue
		}
		if v.IsDir {
			dirSize, acceptedDirsSize := calculateSizes(v)
			thisSize += dirSize
			totalSize += acceptedDirsSize
		} else {
			thisSize += v.Size
		}
	}
	if thisSize <= DirSizeLimit {
		totalSize += thisSize
	}
	return thisSize, totalSize
}

func findDirToDelete(topDir *node, spaceToFree int) (size int, selectedDirSize int, err error) {
	if spaceToFree < 0 {
		err = errors.New("spaceToFree cannot be negative")
		return
	}

	selectedDirSize = TotalSpace
	for k, v := range topDir.Links {
		if k == ".." {
			continue
		}
		if v.IsDir {
			_size, _selectedDirSize, e := findDirToDelete(v, spaceToFree)
			if e != nil {
				err = e
				return
			}
			size += _size
			if _selectedDirSize < selectedDirSize {
				selectedDirSize = _selectedDirSize
			}
		} else {
			size += v.Size
		}
	}

	if size >= spaceToFree {
		if size <= selectedDirSize {
			selectedDirSize = size
		}
	}
	return size, selectedDirSize, nil
}
