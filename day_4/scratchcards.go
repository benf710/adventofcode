package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/benf710/adventofcode/common"
)

func parseLine(line string) ([]string, []string) {
	_, nums, _ := strings.Cut(line, ":")
	allNums := strings.Split(nums, "|")
	winnersStr := strings.Split(allNums[0], " ")
	cardNumsStr := strings.Split(allNums[1], " ")
	return winnersStr, cardNumsStr
}

func main() {
	lines := common.ReadInput("day_4/input.txt")
	points := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		winners, cardNums := parseLine(line)
		winningCardNums := 0
		for _, cardNum := range cardNums {
			for _, winner := range winners {
				if cardNum == winner {
					winningCardNums++
					break
				}
			}
		}

		var cardPoints int
		if winningCardNums == 0 {
			cardPoints = 0
		} else if winningCardNums == 1 {
			cardPoints = 1
		} else {
			cardPoints = int(math.Pow(2, float64(winningCardNums)))
		}
		points += cardPoints
		fmt.Printf("Card points: %d\n", cardPoints)
	}
	fmt.Printf("Total points: %d\n", points)
}
