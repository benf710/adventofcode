package main

import (
    "os"
    "fmt"
    "strings"
    "errors"
)

func checkError(e error) {
    if e != nil {
        panic(e)
    }
}


func isNumber(r rune) bool {
    if r >= 48 && r <= 57 {
        return true
    } else {
        return false
    }
}

func getFirstNum(s string, reverse bool) (int, error) {
    rs := []rune(s)
    var n rune

    for i := 0; i < len(rs); i++ {

        if reverse {
            n = rs[ len(rs)-1 - i]
        } else {
            n = rs[i]
        }

        if isNumber(n) {
            return int(n-'0'), nil
        }
    }
    return -1, errors.New("no numbers found")
}

func main() {
    bytes, err := os.ReadFile("day_1/input.txt")
    checkError(err)
    dat := string(bytes)
    // dat := "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet\n"

    strs := strings.Split(strings.ReplaceAll(dat, "\r\n", "\n"), "\n")

    sums := 0
    for _, str := range strs {
        first, err := getFirstNum(str, false)
        if err == nil {
            last, _ := getFirstNum(str, true)
            val := 10*first + last
            fmt.Printf("First %v Last %v on %v is %v\n", first, last, str, val)
            sums += val
        } else {
            fmt.Printf("ERROR! No digit for '%v'\n", str)
        }
    }
    fmt.Printf("Sum is: %v\n", sums)
}