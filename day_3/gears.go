package main

import (
	"fmt"
	"strconv"

	"github.com/benf710/adventofcode/common"
)

func getSymbolIndecies(line string) []int {
	symbols := []string{"@", "#", "$", "%", "&", "*", "/", "=", "+", "-"}
	var SymbolIndecies []int
	for j, char := range line {
		for _, symbol := range symbols {
			if string(char) == symbol {
				SymbolIndecies = append(SymbolIndecies, j)
			}
		}
	}
	return SymbolIndecies
}

func inSlice(r rune, s []rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}

func getNumRanges(line string) [][]int {
	nums := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	numRanges := [][]int{}

	numStart := -1
	numEnd := -1
	for idx, char := range line {
		if inSlice(rune(char), nums) {
			if numStart == -1 {
				numStart = idx
			} else {
				numEnd = idx
			}
		} else {
			if numStart != -1 && numEnd != -1 {
				numRanges = append(numRanges, []int{numStart, numEnd})
			}
			numStart = -1
		}
	}
	return numRanges
}

func checkSurroundingSymbols(numStart int, numEnd int, symbolIndecies []int) bool {
	for _, symbolIndex := range symbolIndecies {
		if symbolIndex >= numStart-1 && symbolIndex <= numEnd+1 {
			return true
		}
	}
	return false
}

func main() {
	lines := common.ReadInput("day_3/input.txt")

	sum := 0

	for i, line := range lines {

		// Get symbol indecies for surrounding lines
		symbolIndecies := getSymbolIndecies(line)
		if i == 0 {
			symbolIndecies = append(symbolIndecies, getSymbolIndecies(lines[i+1])...)
		} else if i == len(lines)-1 {
			symbolIndecies = append(symbolIndecies, getSymbolIndecies(lines[i-1])...)
		} else {
			symbolIndecies = append(symbolIndecies, getSymbolIndecies(lines[i-1])...)
			symbolIndecies = append(symbolIndecies, getSymbolIndecies(lines[i+1])...)
		}
		fmt.Printf("Symbol Indecies: %v\n", symbolIndecies)

		// Get number indecies for current line
		numRanges := getNumRanges(line)
		fmt.Printf("Num Ranges: %v\n", numRanges)
		for _, numRange := range numRanges {
			numStart := numRange[0]
			numEnd := numRange[1]
			adjacent := checkSurroundingSymbols(numStart, numEnd, symbolIndecies)
			if adjacent {
				fmt.Printf("Num: %q\n", line[numStart:numEnd+1])
				num, err := strconv.Atoi(line[numStart : numEnd+1])
				if err != nil {
					panic(err)
				}
				sum += num
			}
		}
	}

	fmt.Printf("Sum: %d\n", sum) // 458562 too low, 519,840 too low
}
