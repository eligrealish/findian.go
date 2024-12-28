package main

import (
	"fmt"
	"strings"
)

const I = "i"
const N = "n"
const A = "a"

func main() {
	input, done := scanInput()
	if done {
		return
	}
	// checks string starts with i && ends with n && contains an a
	found := checkContainsChars(input)
	if found {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

func scanInput() (string, bool) {
	// scans user string input
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println(err)
		return "", true
	}
	return input, false
}

func checkContainsChars(rawString string) bool {
	// converts to lower case to handle validation
	lowerCaseString := strings.ToLower(rawString)
	// equivalent of an &&, checks all scenarios are true
	if strings.HasPrefix(lowerCaseString, I) {
		if strings.HasSuffix(lowerCaseString, N) {
			if strings.Contains(lowerCaseString, A) {
				return true
			}
		}
	}
	return false
}
