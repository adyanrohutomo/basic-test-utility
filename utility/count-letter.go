package utility

import (
	"fmt"
	"strings"
	"unicode"
)

func CountLetter(input string) string {
	letterCount := make(map[string]int64)
	output := ""
	input = strings.ReplaceAll(input, " ", "")
	input = strings.ReplaceAll(input, ".", "")

	for _, v := range input {
		letterCount[string(unicode.ToLower(v))] += 1
	}

	for i, v := range letterCount {
		output += fmt.Sprintf("%v%v", i, v)
	}
	return output
}
