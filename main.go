package main

import (
	"fmt"
	"basic-test-utility/utility"
)

func main() {
	inputCase := "Lorem Ipsum Dolor Sit Amet"
	resCase := utility.CaseConverter(inputCase)
	fmt.Println(resCase)

	inputCount := "Team Engineering PT. Raksasa Laju Lintang"
	resCount := utility.CountLetter(inputCount)
	fmt.Println(resCount)
}
