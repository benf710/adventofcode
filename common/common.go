package common

import (
	"os"
	"strings"
)

func ReadInput(f string) []string {
	bytes, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}
	dat := string(bytes)
	strs := strings.Split(strings.ReplaceAll(dat, "\r\n", "\n"), "\n")
	return strs
}
