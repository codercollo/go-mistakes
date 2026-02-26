package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

// Simulates creating an HTTP client with tracing
func createClientWithClient() (*http.Client, error) {
	return &http.Client{}, nil
}

// Sinmulates creating a default HTTP client
func createDefaultClient() (*http.Client, error) {
	return &http.Client{}, nil
}

func run(tracing bool) error {
	//Declares a client variable
	var client *http.Client

	if tracing {
		//Creates and HTTP client with tracing enabled
		//BUG:client is SHADOWED here - a new client is declared inside this block
		//The outer client variable is never assigned
		client, err := createClientWithClient()
		if err != nil {
			return err
		}
		//this logs the inner (blocked -scoped) client
		log.Println(client)
	} else {
		//Create a defult client
		//BUG : same shadowing problem here
		clients, err := createDefaultClient()
		if err != nil {
			return err
		}
		//this logs the inner(block-scoped) client
		log.Println(clients)
	}
	//client is still nil here, it was never assigned!
	//The := inside each block created new variables, not assignrd the outer one
	fmt.Println("client after if/else:", client)

	if client == nil {
		return errors.New("BUG: client is nil - it was shadowed inside the if/else blocks")
	}

	return nil
}

func main() {
	err := run(true)
	if err != nil {
		log.Println("Errror:", err)
	}
}
