package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"errors"
	"io"
	"strings"
	"strconv"
)
type item int64

type monkey struct {
	items []item
	operation string
	testFunc func(old item) bool
	monkeyOnTrue int
	monkeyOnFalse int
	activityCount int
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	monkeys := make([]monkey, 0)
	var mnk monkey
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}
		line = strings.TrimSpace(line)
		lineErr := errors.New(fmt.Sprint("bad usage, got: '", line, "'"))

		components := strings.Split(line, ":")
		for len(components) != 0 && components[len(components) - 1] == "" {
			components = components[:len(components) - 1]
		}

		switch len(components) {
		case 1:
			mnk = newMonkey()
		case 2:
			words := strings.FieldsFunc(components[1], func(r rune) bool {
				switch r {
				case ' ', ',':
					return true
				}
				return false
			})
			switch components[0] {
			case "Starting items":
				for _, word := range words {
					num, err := strconv.Atoi(word)
					if err != nil {
						log.Fatal(err)
					}
					mnk.items = append(mnk.items, item(num))
				}
			case "Operation":
				mnk.operation = components[1]
			case "Test":
				if len(words) != 3 || words[0] != "divisible" || words[1] != "by" {
					log.Fatal(lineErr)
				}
				num, err := strconv.Atoi(words[2])
				if err != nil {
					log.Fatal(err)
				}
				mnk.testFunc = func(n item) bool {
					return int(n) % num == 0
				}
			case "If true", "If false":
				if len(words) != 4 || words[0] != "throw" || words[1] != "to" || words[2] != "monkey" {
					log.Fatal(lineErr)
				}
				num, err := strconv.Atoi(words[3])
				if err != nil {
					log.Fatal(err)
				}
				if components[0] == "If true" {
					mnk.monkeyOnTrue = num
				} else {
					mnk.monkeyOnFalse = num
				}
			default:
				log.Fatal(lineErr)
			}
		case 0:
			monkeys = append(monkeys, mnk)
		default:
			log.Fatal(lineErr)
		}
	}
	monkeys = append(monkeys, mnk)

	// fmt.Println("before round 0: ")
	// for j := 0; j < len(monkeys); j++ {
	// 	fmt.Println("monkey", j, ": ", monkeys[j].items)
	// }

	for i := 0; i < 20; i++ {
		for j, _ := range monkeys {
			for _, itm := range monkeys[j].items {
				monkeys[j].activityCount++;
				(&itm).operate(monkeys[j].operation)
				itm /= 3
				if monkeys[j].testFunc(itm) {
					monkeys[monkeys[j].monkeyOnTrue].addItem(itm)
				} else {
					monkeys[monkeys[j].monkeyOnFalse].addItem(itm)
				}
				monkeys[j].items = monkeys[j].items[1:]
			}
		}
		// fmt.Println("after round", i, ": ")
		// for j := 0; j < len(monkeys); j++ {
		// 	fmt.Println("monkey", j, ": ", monkeys[j].items)
		// }
		// fmt.Println()
	}

	var maxActivityCount1, maxActivityCount2 int
	for i := 0; i < len(monkeys); i++ {
		ac := monkeys[i].activityCount
		if ac > maxActivityCount1 {
			maxActivityCount2 = maxActivityCount1
			maxActivityCount1 = ac
		} else if ac > maxActivityCount2 {
			maxActivityCount2 = ac;
		}
	}
	monkeyBuisness := maxActivityCount1 * maxActivityCount2
	fmt.Println(monkeyBuisness)
}


func (itm *item) operate(op string) {
	//TODO
	oldItm := int64(*itm)
	var newItm int64

	words := strings.Fields(op)
	operationErr := errors.New(fmt.Sprint("bad operation usage, got: '", op, "'"))
	if len(words) < 3 || words[0] != "new" || words[1] != "=" || len(words) % 2 == 0 {
		log.Fatal(operationErr)
	}
	switch words[2] {
	case "old":
		newItm = oldItm
	default:
		tmp, err := strconv.Atoi(words[2])
		if err != nil {
			log.Fatal(err)
		}
		newItm = int64(tmp)
	}

	var operatorFunc func (o1, o2 int64) int64
	for i := 3; i < len(words); i += 2 {
		switch words[i] {
		case "*":
			operatorFunc = func (o1, o2 int64) int64 {
				return o1 * o2
			}
		case "+":
			operatorFunc = func (o1, o2 int64) int64 {
				return o1 + o2
			}
		default:
			log.Fatal(operationErr)
		}
		var o2 int64
		switch words[i + 1] {
		case "old":
			o2 = int64(oldItm)
		default:
			tmp, err := strconv.Atoi(words[i + 1])
			if err != nil {
				log.Fatal(operationErr)
			}
			o2 = int64(tmp)
		}
		newItm = operatorFunc(newItm, o2)
	}
	*itm = item(newItm)
}


func (mnk *monkey) addItem(itm item) {
	mnk.items = append(mnk.items, itm)
}


func newMonkey() monkey {
	var m monkey
	m.items = make([]item, 0)
	return m
}
