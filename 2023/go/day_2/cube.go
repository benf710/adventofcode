package main

import (
    "fmt"
    "os"
    "strings"
    "regexp"
    "strconv"
)

func readInput() []string {
    bytes, err := os.ReadFile("day_2/input.txt")
    if err != nil {
        panic(err)
    }
    dat := string(bytes)
    strs := strings.Split(strings.ReplaceAll(dat, "\r\n", "\n"), "\n")
    return strs
}

func getSums(b [][]byte) int {
    sum := 0
    for _, bytes := range b {
        num, _ := strconv.Atoi(string(bytes))
        sum += num
    }
    return sum
}

func roundValid(b []byte) bool {
    maxRed := 12
    maxGreen := 13
    maxBlue := 14

    redsRex := regexp.MustCompile("([0-9]+) red")
    bluesRex := regexp.MustCompile("([0-9]+) blue")
    greensRex := regexp.MustCompile("([0-9]+) green")

    reds := getSums(redsRex.FindSubmatch(b))
    blues := getSums(bluesRex.FindSubmatch(b))
    greens := getSums(greensRex.FindSubmatch(b))

    if reds <= maxRed && blues <= maxBlue && greens <= maxGreen {
        return true
    }
    return false

}

func main() {

    gameNumSums := 0
    gameNumRex := regexp.MustCompile(`Game (\d+):`)

    lines := readInput()
    for _, line := range lines {
        if line == "" {
            continue
        }

        bline := []byte(line)
        gameNum := gameNumRex.FindSubmatch(bline)[1]

        rounds := strings.Split(line, ";")
        valid := true

        for _, round := range rounds {
            if !roundValid([]byte(round)) {
                valid = false
            }
        }

        if valid {
            gameNumInt, _ := strconv.Atoi(string(gameNum))
            gameNumSums += gameNumInt
        }
        fmt.Printf("Game: %q valid: %v\n", gameNum, valid)
    }

    fmt.Println(gameNumSums)
}