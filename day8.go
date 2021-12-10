package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Display struct {
	noise []string
	data  []string
}

var Sizes [10]int = [10]int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}

var dataDay8 = func() (dataDay8 []Display) {
	file, _ := os.Open("data/day8.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")
		dataDay8 = append(dataDay8, Display{
			noise: strings.Split(line[0], " "),
			data:  strings.Split(line[1], " "),
		})
	}

	return
}

func equal(source, target string) bool {
	for _, s := range source {
		if !strings.ContainsRune(target, s) {
			return false
		}
	}

	return true
}

func substractString(s, m string) string {
	for _, c := range m {
		s = strings.Replace(s, string(c), "", 1)
	}
	return s
}

func decode(noise []string, lengthMap map[int][]string) map[string]int {
	parsing := map[string]int{
		lengthMap[2][0]: 1,
		lengthMap[3][0]: 7,
		lengthMap[4][0]: 4,
		lengthMap[7][0]: 8,
	}

	one := lengthMap[2][0]
	four := lengthMap[4][0]

	var nine string
	for _, sixLetter := range lengthMap[6] {
		if len(substractString(sixLetter, one)) == 5 {
			parsing[sixLetter] = 6
		} else if len(substractString(sixLetter, four)) == 3 {
			parsing[sixLetter] = 0
		} else {
			nine = sixLetter
			parsing[sixLetter] = 9
		}
	}

	for _, fiveLetter := range lengthMap[5] {
		if len(substractString(fiveLetter, one)) == 3 {
			parsing[fiveLetter] = 3
		} else if len(substractString(fiveLetter, nine)) == 0 {
			parsing[fiveLetter] = 5
		} else {
			parsing[fiveLetter] = 2
		}
	}

	return parsing
}

func day8part2(data []Display) int {
	var total int
	for _, line := range data {
		var number string

		mapByLen := map[int][]string{}

		for _, in := range line.noise {
			mapByLen[len(in)] = append(mapByLen[len(in)], in)
		}

		table := decode(line.noise, mapByLen)

		for _, d := range line.data {
			if n, ok := table[d]; ok {
				number = fmt.Sprintf("%v%v", number, n)
				continue
			}

			for _, combination := range mapByLen[len(d)] {
				if equal(d, combination) {
					number = fmt.Sprintf("%v%v", number, table[combination])
					break
				}
			}
		}

		n, _ := strconv.Atoi(number)
		total += n
	}

	return total
}

func day8part1(data []Display) int {
	var count int

	for _, d := range data {
		for _, n := range d.data {
			switch len(n) {
			case Sizes[1], Sizes[4], Sizes[7], Sizes[8]:
				count++
			}
		}
	}

	return count
}

func day8() {
	data := dataDay8()
	fmt.Printf("Part 1: %v\n", day8part1(data))
	fmt.Printf("Part 2: %v\n", day8part2(data))
}
