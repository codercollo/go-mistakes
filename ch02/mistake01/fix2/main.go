package main

import (
	"fmt"
	"log"
	"net/http"
)

// Simulates creating an HTTP client with tracing
func createClientWithTracing() (*http.Client, error) {
	return &http.Client{}, nil
}

// Simulates creating a default HTTP client
func createDefaultClient() (*http.Client, error) {
	return &http.Client{}, nil
}

func run(tracing bool) error {
	//FIX 2: Pre-declare both client and err outside the if/else
	//This way we can user = instead of := inside the blocks
	var (
		client *http.Client
		err    error
	)

	if tracing {
		// Using = not := so we assign to the outer client and err
		client, err = createClientWithTracing()
	} else {
		client, err = createDefaultClient()
	}

	//Single error handling after the if/else
	if err != nil {
		return err
	}

	//client is correctly assingned
	fmt.Println("client after if/else:", client)
	log.Println("FIX 2 works: client is not nil, error handling is cleaner")
	return nil
}

func main() {
	err := run(true)
	if err != nil {
		log.Println("Error", err)
	}

}
