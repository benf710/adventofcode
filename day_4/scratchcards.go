package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/benf710/adventofcode/common"
)

func removeSpaces(s []string) []string {
	var ret []string
	for _, str := range s {
		if strings.TrimSpace(str) != "" {
			ret = append(ret, str)
		}
	}
	return ret
}

func parseLine(line string) ([]string, []string) {
	_, nums, _ := strings.Cut(line, ":")
	allNums := strings.Split(nums, "|")
	winnersStr := removeSpaces(strings.Split(allNums[0], " "))
	cardNumsStr := removeSpaces(strings.Split(allNums[1], " "))
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
		var matchedNums []string
		for _, cardNum := range cardNums {
			for _, winner := range winners {
				if cardNum == winner {
					winningCardNums++
					matchedNums = append(matchedNums, cardNum)
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
			cardPoints = int(math.Pow(2, float64(winningCardNums)-1))
		}
		points += cardPoints
		fmt.Printf("Matched nums: %q\nCard points: %d\n", matchedNums, cardPoints)
	}
	fmt.Printf("Total points: %d\n", points) // 646754 too high 323378 too high
}
