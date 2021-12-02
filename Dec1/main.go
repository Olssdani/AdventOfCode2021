package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	lastValue := -1
	numIncrease := 0
	for _, line := range lines {
		value, _ := strconv.Atoi(line)
		if lastValue >= 0 && lastValue < value {
			numIncrease++
		}
		lastValue = value
	}
	fmt.Println("Increase", numIncrease)

	// B
	numIncrease = 0
	for i := 1; i < len(lines)-1; i++ {
		value1, _ := strconv.Atoi(lines[i-1])
		value2, _ := strconv.Atoi(lines[i])
		value3, _ := strconv.Atoi(lines[i+1])
		value := value1 + value2 + value3

		if lastValue >= 0 && lastValue < value {
			numIncrease++
		}
		lastValue = value
	}
	fmt.Println("Increase", numIncrease)
}
