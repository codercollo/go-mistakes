package main

import (
	"fmt"
	"sync"
)

//VALID USE: getters and setters make sense here because they ADD behavior
// 1. Balance() - validation logic in setter (negative balance guard)
// 2. SafeCounter - mutex wrapping for concurrent access

// EXAMPLE 1: Setter with validation logic
type Customer struct {
	balance float64
}

// Balance - correct Go naming for getter (Not GetBalance)
func (c *Customer) Balance() float64 {
	return c.balance
}

// SetBalance - correct Go naming for setter
// Adds value: validates that balance can't go below 0
func (c *Customer) SetBalance(amount float64) {
	if amount < 0 {
		amount = 0 //validation logic this is why the setter exists
	}
	c.balance = amount
}

// EXAMPLE 2: Getter/setter wrapping a mutex for concurrency
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// Count - getter acquires lock before reading
func (s *SafeCounter) Count() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.count
}

// SetCount - setter acquires lock before writing
func (s *SafeCounter) SetCount(n int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count = n
}

func main() {
	//Example 1 - validation in setter
	c := &Customer{}
	c.SetBalance(1000.0)
	fmt.Println(c.Balance())

	c.SetBalance(-500.0)
	fmt.Println(c.Balance()) //setter blocked negative value

	//Example 2 - mutex-wrapped access
	counter := &SafeCounter{}
	counter.SetCount(10)
	fmt.Println(counter.Count()) //safeley read with lock
}
