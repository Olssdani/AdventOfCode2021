package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	// A
	x := 0
	y := 0
	for _, line := range lines {
		res := strings.Split(line, " ")
		value, _ := strconv.Atoi(res[1])
		if res[0] == "forward" {
			x += value
		} else if res[0] == "up" {
			y -= value
		} else if res[0] == "down" {
			y += value
		}
	}
	fmt.Println("A: x", x, "y", y, "mult", x*y)

	// B
	x = 0
	y = 0
	aim := 0
	for _, line := range lines {
		res := strings.Split(line, " ")
		value, err := strconv.Atoi(res[1])
		if err != nil {
			fmt.Println("Error in parsing string to int")
			return
		}
		if res[0] == "forward" {
			x += value
			y += value * aim
		} else if res[0] == "up" {
			aim -= value
		} else if res[0] == "down" {
			aim += value
		}
	}
	fmt.Println("B: x", x, "y", y, "mult", x*y)
}
