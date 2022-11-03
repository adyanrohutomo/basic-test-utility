package utility

import (
	"unicode"
)

func CaseConverter(input string) string {
	output := ""
	for _, v := range input {
		if unicode.IsUpper(v) {
			output += string(unicode.ToLower(v))
		} else {
			output += string(unicode.ToUpper(v))
		}
	}
	return output
}
