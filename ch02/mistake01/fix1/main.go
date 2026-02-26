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
	//Declares the outer client variable
	var client *http.Client

	if tracing {
		//FIX 1: Use a temporary variable c inside the block
		//then assign it to the outer client using =
		c, err := createClientWithTracing()
		if err != nil {
			return err
		}
		//Assigns to the OUTER client - not shadowing
		client = c

	} else {
		c, err := createDefaultClient()
		if err != nil {
			return err
		}
		//Assigns to the OUTER client - not shadowing
		client = c
	}

	//Client is correctly assigned now
	fmt.Println("client after if/else:", client)
	log.Println("FIX 1 works: client is not nil")
	return nil
}

func main() {
	err := run(true)
	if err != nil {
		log.Println("Error:", err)
	}
}
