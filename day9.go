package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Floor [][]int

func (f Floor) Locate(pt Loc) int {
	return f[pt.Y][pt.X]
}

func (f Floor) InBounds(pt Loc) bool {
	if pt.X < 0 || pt.Y < 0 || pt.X == len(f[0]) || pt.Y == len(f) {
		return false
	}

	return true
}

func (f Floor) Flood(pt Loc, lastVal int) int {
	if !f.InBounds(pt) {
		return 0
	}

	curr := f.Locate(pt)
	flow := curr-lastVal > 0

	if !flow || curr >= 9 {
		return 0
	}

	f[pt.Y][pt.X] = 10

	return 1 +
		f.Flood(pt.Move(0, -1), curr) +
		f.Flood(pt.Move(0, 1), curr) +
		f.Flood(pt.Move(-1, 0), curr) +
		f.Flood(pt.Move(1, 0), curr)
}

func (f Floor) BasinSize(pt Loc) int {
	start := f.Locate(pt)

	return 1 +
		f.Flood(pt.Move(0, -1), start) +
		f.Flood(pt.Move(0, 1), start) +
		f.Flood(pt.Move(-1, 0), start) +
		f.Flood(pt.Move(1, 0), start)
}

type Loc struct {
	image.Point
}

func NewLocation(x, y int) Loc {
	return Loc{image.Point{x, y}}
}

func (p Loc) Move(n, m int) Loc {
	return NewLocation(p.X+n, p.Y+m)
}

var dataDay9 = func() (dataDay9 Floor) {
	file, _ := os.Open("data/day9.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var lineData []int
		line := strings.Split(scanner.Text(), "")
		for _, c := range line {
			n, _ := strconv.Atoi(c)
			lineData = append(lineData, n)
		}
		dataDay9 = append(dataDay9, lineData)
	}

	return
}

func day9part1(fl Floor) int {
	var total int

	for y := range fl {
		for x := range fl[y] {
			pt := NewLocation(x, y)
			if isvent(fl, pt) {
				total += fl.Locate(pt) + 1
			}
		}
	}

	return total
}

func day9part2(floor Floor) int {
	var basins []int

	for y := range floor {
		for x := range floor[y] {
			loc := NewLocation(x, y)
			if isvent(floor, loc) {
				basins = append(basins, floor.BasinSize(loc))
			}
		}
	}

	sort.Ints(basins)
	return basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3]
}

func isvent(floor Floor, pt Loc) bool {
	var min bool = true

	if pt.Y == 0 {
		min = min && (floor.Locate(pt) < floor.Locate(pt.Move(0, 1)))
	} else if pt.Y == len(floor)-1 {
		min = min && (floor.Locate(pt) < floor.Locate(pt.Move(0, -1)))
	} else {
		min = min && (floor.Locate(pt) < floor.Locate(pt.Move(0, -1)))
		min = min && (floor.Locate(pt) < floor.Locate(pt.Move(0, 1)))
	}

	if !min {
		return false
	}

	if pt.X == 0 {
		min = min && (floor.Locate(pt) < floor.Locate(pt.Move(1, 0)))
	} else if pt.X == len(floor[0])-1 {
		min = min && (floor.Locate(pt) < floor.Locate(pt.Move(-1, 0)))
	} else {
		min = min && (floor.Locate(pt) < floor.Locate(pt.Move(-1, 0)))
		min = min && (floor.Locate(pt) < floor.Locate(pt.Move(1, 0)))
	}

	return min
}

func day9() {
	data := dataDay9()
	fmt.Printf("Part 1: %v\n", day9part1(data))
	fmt.Printf("Part 2: %v\n", day9part2(data))
}
