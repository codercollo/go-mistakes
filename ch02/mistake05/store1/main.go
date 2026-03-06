package store1

///BAD: Interface defined on the PRODUCER side (same package as the implementation)
//This forces every consumer to use the full abstrctive regardless of their needs
//It also creates an akward dependency direction

//Interface lives in the same package as its implementation
//The producer is deciding the abstraction for all consumers
//A consumer that only needs GetAllCustomers is still couple to 6 methods

type CustomerStorage interface {
	StoreCustomer(customer Customer) error
	GetCustomer(id string) (Customer, error)
	UpdateCustomer(customer Customer) error
	GetAllCustomers() ([]Customer, error)
	GetCustomersWithoutContract() ([]Customer, error)
	GetCustomersWithNegativeBalance() ([]Customer, error)
}

// Customer domain model
type Customer struct {
	ID       string
	Name     string
	Balance  float64
	Contract bool
}

// Concrete implementation in the same package as the interface above.
type InMemoryStore struct {
	customers map[string]Customer // in-memory customer storage
}

// Constructor for in-memory store
func NewInMemoryStore() CustomerStorage { // ❌ returns the interface — see mistake07
	return &InMemoryStore{customers: make(map[string]Customer)} // initialize map storage
}

// Save customer in memory
func (s *InMemoryStore) StoreCustomer(c Customer) error {
	s.customers[c.ID] = c
	return nil
}

// Retrieve customer by ID
func (s *InMemoryStore) GetCustomer(id string) (Customer, error) {
	return s.customers[id], nil
}

// Update customer record
func (s *InMemoryStore) UpdateCustomer(c Customer) error {
	s.customers[c.ID] = c
	return nil
}

// Return all customers
func (s *InMemoryStore) GetAllCustomers() ([]Customer, error) {
	all := make([]Customer, 0, len(s.customers)) // allocate slice
	for _, c := range s.customers {              // iterate over map
		all = append(all, c)
	}
	return all, nil
}

// Placeholder: customers without contracts
func (s *InMemoryStore) GetCustomersWithoutContract() ([]Customer, error) { return nil, nil }

// Placeholder: customers with negative balances
func (s *InMemoryStore) GetCustomersWithNegativeBalance() ([]Customer, error) { return nil, nil }
