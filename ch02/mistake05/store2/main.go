package store2

//FIX (producer side): Expose only the concrete implementation
//No interface lives here. Let each consumer define what it needs

// Customer domain model
type Customer struct {
	ID       string
	Name     string
	Balance  float64
	Contract bool
}

// ✅ Concrete struct exported — consumers can use it directly or
// define their own minimal interface against it.
type InMemoryStore struct {
	customers map[string]Customer // in-memory storage map
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

// Return all stored customers
func (s *InMemoryStore) GetAllCustomers() ([]Customer, error) {
	all := make([]Customer, 0, len(s.customers)) // allocate result slice
	for _, c := range s.customers {              // iterate through map
		all = append(all, c)
	}
	return all, nil
}

// Placeholder: customers without contracts
func (s *InMemoryStore) GetCustomersWithoutContract() ([]Customer, error) { return nil, nil }

// Placeholder: customers with negative balances
func (s *InMemoryStore) GetCustomersWithNegativeBalance() ([]Customer, error) { return nil, nil }
