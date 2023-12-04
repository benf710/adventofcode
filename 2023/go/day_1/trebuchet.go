package main

import (
    "os"
    "fmt"
    "strings"
    "errors"
)


func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func isNumber(r rune) bool {
    if r >= 48 && r <= 57 {
        return true
    } else {
        return false
    }
}

func GetFirstNum(s string) (int, error) {
    for _, char := range s {
        if isNumber(char) {
            return int(char-'0'), nil
        }
    }
    return -1, errors.New("no numbers found")
}

func GetLastNum(s string) (int, error) {
    rs := []rune(s)
    for i := len(rs)-1; i >= 0; i-- {
        char := rs[i]
        if isNumber(char) {
            return int(char-'0'), nil
        }
    }
    return -1, errors.New("no numbers found")
}

func convertNums(s string) string {
    spelledNums := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

    var firstNum string
    lowestIndex := len([]rune(s))
    for name := range spelledNums {
        index := strings.Index(s, name)
        if index != -1 && index < lowestIndex {
            firstNum = name
            lowestIndex = index
        }
    }

    var lastNum string
    highestIndex := 0
    for name := range spelledNums {
        index := strings.LastIndex(s, name)
        if index != -1 && index > highestIndex {
            lastNum = name
            highestIndex = index
        }
    }

    if lowestIndex > highestIndex {
        return s
    }

    rs := []rune(s)
    fmt.Printf("Low index: %v, High index: %v on %v\n", lowestIndex, highestIndex, s)
    s = string(rs[:lowestIndex]) + spelledNums[firstNum] + string(rs[lowestIndex:highestIndex]) + spelledNums[lastNum] + string(rs[highestIndex:])

    return s
}

func readInput() []string {
    bytes, err := os.ReadFile("day_1/input.txt")
    if err != nil {
        panic(err)
    }
    dat := string(bytes)
    strs := strings.Split(strings.ReplaceAll(dat, "\r\n", "\n"), "\n")
    return strs
}

func main() {
    strs := readInput()
    sums := 0
    for _, str := range strs {
        normalized_str := convertNums(str)
        first, err := GetFirstNum(normalized_str)
        if err == nil {
            last, _ := GetLastNum(normalized_str)
            val := 10*first + last
            fmt.Printf("First %v Last %v on %v is %v\n", first, last, normalized_str, val)
            sums += val
        } else {
            fmt.Printf("ERROR! No digit for '%v'\n", normalized_str)
        }
    }
    fmt.Printf("Sum is: %v\n", sums)
}