package main

import "fmt"

//FIX: idiomatic Go - export fields directly when no behavior is needed
// Simple data structs don't need getters/setters
//Direct access is cleaner and more readable

type Customer struct {
	ID      int
	Name    string
	Balance float64
}

func main() {
	c := Customer{
		ID:      1,
		Name:    "Collins",
		Balance: 1000.0,
	}

	//Direct access - clean, simple idiomatic Go
	fmt.Println(c.ID)
	fmt.Println(c.Name)
	fmt.Println(c.Balance)

	//Direct mutation - fine when no validation or logic is needed
	if c.Balance < 0 {
		c.Balance = 0
	}
}
