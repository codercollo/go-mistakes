package main

import (
	"errors"
	"fmt"
	"log"
)

// BAD: deeply nested if/else - hard to read, hard to build a mental model
// You have to track multiple levels simultaneously to understand the flow
func join(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")
	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatenate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

// concatenate function returns a string of joined s1 and s2 strings
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

	//Test max truncuation
	result, err = join("Hello, ", "World", 5)
	if err != nil {
		log.Println("Error:", err)
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
