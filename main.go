package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(converter("this is the#$*%$## sample text @@#$&&&&&&for you buddy$$$&#^#*&*"))
	fmt.Println(calculation(100))
}

// first function

func converter(input string) (string, error) {
	if len(input) == 11 {
		return "", fmt.Errorf("input length is 11 and it is not allowed")
	}

	splitter := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	splits := strings.FieldsFunc(input, splitter)
	out := strings.Join(splits, "__")

	return out, nil
}

// second function

func calculation(x int) (result int) {
	result = x * 2
	return result
}
