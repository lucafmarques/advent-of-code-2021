package main

import (
	"bufio"
	"fmt"

	// "time"
	"os"
	"strconv"
	"strings"
)

type Value struct {
	value string
	found bool
}

type Row []*Value

func (r *Row) markFound(target []string) {
	for _, v := range *r {
		if v.found {
			continue
		}

		for _, a := range target {
			if v.value == a {
				v.found = true
			}
		}
	}
}

func (r Row) String() string {
	var values []string
	for _, v := range r {
		if v.found {
			values = append(values, "<>")
		} else {
			values = append(values, v.value)
		}
	}
	return fmt.Sprintf("[%v\t]", strings.Join(values, "\t"))
}

type Board struct {
	rows  []Row
	found bool
}

func (b Board) String() string {
	var output string
	for _, row := range b.rows {
		output = fmt.Sprintf("%v%v\n", output, row)
	}

	return output
}

type Game []Board

func (g Game) String() string {
	var gamestate string
	for i := 0; i < 30; i += 5 {
		var boards string
		for j := 0; j < len(g[0].rows); j++ {
			var rows string
			for _, board := range g[i : i+5] {
				rows = fmt.Sprintf("%v\t%v", rows, board.rows[j])
			}
			boards = fmt.Sprintf("%v\n%v", boards, rows)
		}
		gamestate = fmt.Sprintf("%v\n%v", gamestate, boards)
	}

	return gamestate
}

var bingo = func() (bingo []string) {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	bingo = strings.Split(scanner.Text(), ",")

	return
}()
var loadGame = func() Game {
	var data Game
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()

	for scanner.Scan() {
		board := Board{
			rows: []Row{},
		}
		for i := 0; i < 5; i++ {
			var parsedLine Row
			scanner.Scan()
			line := scanner.Text()
			for i := 0; i < len(line); i += 3 {
				v := Value{
					value: strings.TrimSpace(line[i : i+2]),
					found: false,
				}
				parsedLine = append(parsedLine, &v)
			}
			board.rows = append(board.rows, parsedLine)
		}
		data = append(data, board)
	}

	return data
}

func main() {
	p1 := part1(bingo, loadGame())
	p2 := part2(bingo, loadGame())
	fmt.Printf("Part 1: %v\n", p1)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(bingo []string, data Game) int {
	var total int
	pos, board := findFirstBoardWin(bingo, data)
	num, _ := strconv.Atoi(bingo[pos])

	for _, row := range board.rows {
		for _, v := range row {
			if v.found == false {
				point, _ := strconv.Atoi(v.value)
				total += point
			}
		}
	}

	return total * num
}

func part2(bingo []string, data Game) int {
	var total int
	pos, board := findLastBoardWin(bingo, data)
	num, _ := strconv.Atoi(bingo[pos])

	for _, row := range board.rows {
		for _, v := range row {
			if v.found == false {
				point, _ := strconv.Atoi(v.value)
				total += point
			}
		}
	}

	return total * num
}

func findFirstBoardWin(bingo []string, data Game) (int, *Board) {
	for i := 1; i < len(bingo); i++ {
		for _, board := range data {
			for pos, row := range board.rows {
				row.markFound(bingo[:i+1])

				column := Row{
					board.rows[0][pos],
					board.rows[1][pos],
					board.rows[2][pos],
					board.rows[3][pos],
					board.rows[4][pos],
				}
				column.markFound(bingo[:i+1])
			}

			if board.questionBingo() {
				board.found = true
				return i, &board
			}
		}
		// fmt.Println("Draw:", bingo[i-1])
		// fmt.Printf("%v", data)
		// time.Sleep(500 * time.Millisecond)
		// fmt.Print("\033[H\033[2J")
	}

	return 0, nil
}

func findLastBoardWin(bingo []string, boards Game) (int, *Board) {
	var pos int
	var last *Board
	for i := 5; i < len(bingo); i++ {
		for j := range boards {
			if boards[j].found {
				continue
			}

			for pos, row := range boards[j].rows {
				row.markFound(bingo[:i+1])

				column := Row{
					boards[j].rows[0][pos],
					boards[j].rows[1][pos],
					boards[j].rows[2][pos],
					boards[j].rows[3][pos],
					boards[j].rows[4][pos],
				}
				column.markFound(bingo[:i+1])
			}

			if boards[j].questionBingo() {
				if !boards[j].found {
					pos = i
					last = &boards[j]
				}
				boards[j].found = true
			}
		}
		// fmt.Println("Draw:", bingo[i-5])
		// fmt.Printf("%v", boards)
		// time.Sleep(500 * time.Millisecond)
		// fmt.Print("\033[H\033[2J")
	}

	return pos, last
}

func (b Board) questionBingo() bool {
	for _, row := range b.rows {
		var bingo bool = true

		for _, v := range row {
			bingo = bingo && v.found
		}

		if bingo {
			return bingo
		}
	}

	for i := range b.rows {
		var bingo bool = true

		for j := range b.rows {
			bingo = bingo && b.rows[j][i].found
		}

		if bingo {
			return bingo
		}
	}

	return false
}
