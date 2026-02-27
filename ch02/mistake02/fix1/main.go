package main

import (
	"errors"
	"fmt"
	"log"
)

// FIX: align the happy path to the left
// Handle edge cases (non-happy path) early with guard clauses and return
// Scan down the LEFT column = happy path
// Scan down the SECOND column = edge cases
func join(s1, s2 string, max int) (string, error) {
	//Guard clause: handle non-happy path early, return immediately
	if s1 == "" {
		return "", errors.New("s1 is empty")
	}

	//Guard clause: same pattern - flip condition, return early
	if s2 == "" {
		return "", errors.New("s2 is empty")
	}

	//Happy path continues on the left edge - no else needed after a return
	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}

	if len(concat) > max {
		return concat[:max], nil
	}

	//Finally happy path return, clearly visible at the left edge
	return concat, nil
}

func concatenate(s1 string, s2 string) (string, error) {
	return s1 + s2, nil
}

func main() {
	//Test normal case
	result, err := join("Hello, ", "World", 20)
	if err != nil {
		log.Println("Error:", err)
	}
	fmt.Println(result)

	//Test max truncation
	result, err = join("Hello, ", "World", 5)
	if err != nil {
		log.Println(result)
	}
	fmt.Println(result)

	//Test empty s1
	_, err = join("", "World", 20)
	if err != nil {
		log.Println("Error:", err)
	}

	//Test empty s2
	_, err = join("Hello", "", 20)
	if err != nil {
		log.Println("Error:", err)
	}
}
