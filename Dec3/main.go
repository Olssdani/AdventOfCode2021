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

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func bitsToUint(slice string, length int) uint {
	var res uint = 0
	for i := 0; i < length; i++ {
		value, _ := strconv.Atoi(string(slice[length-i-1]))
		res |= uint(value) << i
	}
	return res
}

func getMostCommonBit(slice []string, bitIndex int) (int, int) {
	sum := 0
	for _, line := range slice {
		value, _ := strconv.Atoi(string(line[bitIndex]))
		sum += value
	}
	if sum >= len(slice)-sum {
		return 1, sum
	} else {
		return 0, len(slice) - sum
	}
}

func main() {
	lines, err := readLines("data.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// A
	// Size of string = number of bits
	nrBits := len([]rune(lines[0]))
	dataLength := len(lines)

	// Intial slice with values 0
	sum := make([]int, nrBits)
	for i := range sum {
		sum[i] = 0
	}

	for _, line := range lines {
		for i := 0; i < nrBits; i++ {
			value, _ := strconv.Atoi(string(line[i]))
			sum[i] += value
		}
	}

	var gamma uint32 = 0
	for i := 0; i < nrBits; i++ {
		value := uint32(sum[nrBits-i-1] / (dataLength / 2))
		gamma |= value << i
	}

	shift := 32 - nrBits
	epsilon := ^gamma << shift >> shift
	fmt.Println("Gamma rate:", gamma, "Epsilon Rate", epsilon, "Power", epsilon*gamma)

	// B
	var oxygen uint = 0
	var co2 uint = 0

	{
		linesVector := make([]string, len(lines))
		copy(linesVector, lines)
		bitNumber := 0
		for len(linesVector) > 1 {
			index := 0
			compare, sum := getMostCommonBit(linesVector, bitNumber)
			for len(linesVector) > sum {
				v, _ := strconv.Atoi(string(linesVector[index][bitNumber]))
				if v != compare {
					linesVector = remove(linesVector, index)
				} else {
					index++
				}
			}
			bitNumber++
		}
		oxygen = bitsToUint(linesVector[0], nrBits)
	}

	{
		linesVector := make([]string, len(lines))
		copy(linesVector, lines)
		bitNumber := 0

		for len(linesVector) > 1 {
			index := 0
			compare, sum := getMostCommonBit(linesVector, bitNumber)
			sum = len(linesVector) - sum
			if compare > 0 {
				compare = 0
			} else {
				compare = 1
			}

			for len(linesVector) > sum {
				v, _ := strconv.Atoi(string(linesVector[index][bitNumber]))
				if v != compare {
					linesVector = remove(linesVector, index)
				} else {
					index++
				}
			}
			bitNumber++
		}
		co2 = bitsToUint(linesVector[0], nrBits)
	}

	fmt.Println("Oxygen", oxygen, "co2", co2, "Answer", co2*oxygen)
}
