package main

import "fmt"

type Customer struct {
	Name string
}

type Contract struct {
	Title string
}

type Store struct {
	customers map[string]Customer
	contracts map[string]Contract
}

// Constructor: Creates and initializes a Store
// Returning the concrete type (*Store) keeps the API simple and flexibel
func NewStore() *Store {
	return &Store{
		customers: make(map[string]Customer),
		contracts: make(map[string]Contract),
	}
}

// Explicit method for Customer operations
// The method name and types clearly communicate intent
// Callers immediately know what type is expected and returned
// The compiler prevents passing the wrong type
func (s *Store) GetCustomer(id string) (Customer, error) {
	c, ok := s.customers[id]
	if !ok {
		return Customer{}, fmt.Errorf("customer %s not found", id)
	}
	return c, nil
}

//Parameter type is Customer, so only valid data can be stored.
//Trying to pass int, string etc would cause a compile-time error
func (s *Store) SetCustomer(id string, customer Customer) error {
	s.customers[id] = customer
	return nil
}

//Separate method for Contract operations
//Each method handles exactly one type
//This keeps the API clear and avoids unsafe patterns like `any`
func (s *Store) GetContract(id string) (Contract, error) {
	c, ok := s.contracts[id]
	if !ok {
		return Contract{}, fmt.Errorf("contract %s not found", id)
	}
	return c, nil
}

//Explicit parameter type again provides compile-time safety
func (s *Store) SetContract(id string, contract Contract) error {
	s.contracts[id] = contract
	return nil
}

func main() {
	s := NewStore()

	// ✅ Compiler enforces correct types.
	// Only a Customer can be passed to SetCustomer.
	// Only a Contract can be passed to SetContract.

	_ = s.SetCustomer("c1", Customer{Name: "Alice"})
	_ = s.SetContract("k1", Contract{Title: "Service Agreement"})

	// ✅ Return type is concrete (Customer).
	// No type assertion needed, unlike when returning `any`.
	customer, err := s.GetCustomer("c1")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(customer.Name) // "Alice"

	// ✅ Consumers can still define their own interfaces if needed.
	// The Store doesn't need to know about them.
}

// ✅ Consumer-defined interface.
//
// A consumer package might only need contract operations,
// so it defines a minimal interface containing only those methods.
//
// Go allows this because interfaces are satisfied implicitly.
type ContractStorer interface {
	GetContract(id string) (Contract, error)
	SetContract(id string, contract Contract) error
}
