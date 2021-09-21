package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if result, err := Concat(args...); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Concatenated string: '%s'\n", result)
	}

	// streamlined version of above
	result, _ := Concat(args...)
	fmt.Printf("Concatenated string: '%s'\n", result)

	// version that prints the error
	result, err := Concat(args...)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func Concat(parts ...string) (string, error) {
	if len(parts) == 0 {
		return "", errors.New("No strings supplied")
	}
	return strings.Join(parts, " "), nil
}
