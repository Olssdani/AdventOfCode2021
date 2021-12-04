package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BingoBoard struct {
	board  [5][5]int
	marked [5][5]bool
}

func (board *BingoBoard) markValue(value int) {
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if board.board[r][c] == value {
				board.marked[r][c] = true
			}
		}
	}
}

func (board BingoBoard) sumUnmark() int {
	sum := 0
	for r := 0; r < 5; r++ {
		for c := 0; c < 5; c++ {
			if !board.marked[r][c] {
				sum += board.board[r][c]
			}
		}
	}
	return sum
}

func (board BingoBoard) hasBingo() (bool, int) {
	for r := 0; r < 5; r++ {
		bingo := 0
		for c := 0; c < 5; c++ {
			if board.marked[r][c] {
				bingo++
			}
		}
		if bingo == 5 {
			sum := board.sumUnmark()
			return true, sum
		}
	}

	for c := 0; c < 5; c++ {
		bingo := 0
		for r := 0; r < 5; r++ {
			if board.marked[r][c] {
				bingo++
			}
		}
		if bingo == 5 {
			sum := board.sumUnmark()
			return true, sum
		}
	}
	return false, 0
}

func New(values []string) BingoBoard {
	var board [5][5]int
	var marked [5][5]bool

	for r := 0; r < 5; r++ {
		res := strings.Split(values[r], " ")
		c := 0
		for _, data := range res {
			value, err := strconv.Atoi(string(data))
			if err == nil {
				board[r][c] = value
				marked[r][c] = false
				c++
			}

		}
	}

	b := BingoBoard{board, marked}
	return b
}

func remove(slice []BingoBoard, s int) []BingoBoard {
	return append(slice[:s], slice[s+1:]...)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, err := readLines("data.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Setup data
	bingoBoards := make([]BingoBoard, 0)
	bingoNumbers := make([]int, 0)
	{
		line := lines[0]
		res := strings.Split(line, ",")

		for _, c := range res {
			value, _ := strconv.Atoi(c)
			bingoNumbers = append(bingoNumbers, value)
		}
	}

	for i := 2; i < len(lines); i = i + 6 {
		board := New(lines[i : i+5])
		bingoBoards = append(bingoBoards, board)
	}

	// A
	for _, value := range bingoNumbers {
		hasBingo := false
		for c := range bingoBoards {
			bingoBoards[c].markValue(value)
			bingo, sum := bingoBoards[c].hasBingo()
			if bingo {
				fmt.Println("Number drawn:", value, "Sum of unmarked:", sum, "Result", sum*value)
				hasBingo = true
				break
			}
		}

		if hasBingo {
			break
		}
	}

	// B
	for _, value := range bingoNumbers {
		removeIndex := make([]int, 0)
		for c := range bingoBoards {
			bingoBoards[c].markValue(value)
			bingo, sum := bingoBoards[c].hasBingo()
			if bingo {
				if len(bingoBoards) == 1 {
					fmt.Println("Number drawn:", value, "Sum of unmarked:", sum, "Result", sum*value)
					return
				}
				removeIndex = append(removeIndex, c)

			}
		}
		for i := (len(removeIndex) - 1); i >= 0; i-- {
			bingoBoards = remove(bingoBoards, removeIndex[i])
		}
	}
}
