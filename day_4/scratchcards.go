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

func main_part_1() {
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

func main_part_2() {
	lines := common.ReadInput("day_4/input.txt")
	cards := 0
	var nextCardCoppies []int
	for index, line := range lines {
		if line == "" {
			continue
		}
		winners, cardNums := parseLine(line)

		// Get the count of winning numbers
		winningCardNums := 0
		for _, cardNum := range cardNums {
			for _, winner := range winners {
				if cardNum == winner {
					winningCardNums++
					break
				}
			}
		}

		coppies := 0
		if len(nextCardCoppies) != 0 {
			coppies = nextCardCoppies[0]
			nextCardCoppies = nextCardCoppies[1:]
		}

		for i := 0; i < coppies+1; i++ {
			for j := 0; j < winningCardNums; j++ {
				if j < len(nextCardCoppies) {
					nextCardCoppies[j]++
				} else {
					nextCardCoppies = append(nextCardCoppies, 1)
				}
			}
		}

		cards += coppies + 1
		fmt.Printf("Card %d: (%d coppies) %d winners, %v nextCardCoppies\n", index+1, coppies, winningCardNums, nextCardCoppies)
	}
	fmt.Printf("Total cards: %d\n", cards) // 646754 too high 323378 too high
}

func main() {
	//main_part_1()
	main_part_2()
}
