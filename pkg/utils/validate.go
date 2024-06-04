package utils

import (
	"fmt"
	"regexp"
)

func Validate(value string, field string) {
	secure := true
	key := field
	var tests []string
	switch key {
	case "username":
		tests = []string{".{7,}", "[a-z]", "[A-Z]", "[^\\d\\w]"}
		fmt.Println(value)
	case "password":
		tests = []string{".{7,}", "[a-z]", "[A-Z]", "[0-9]", "[^\\d\\w]"}
		fmt.Println(value)
	default:
		fmt.Println("switch error")
	}

	for _, test := range tests {
		t, _ := regexp.MatchString(test, value)
		if !t {
			secure = false
			break
		}
	}
	fmt.Print(secure)
}
