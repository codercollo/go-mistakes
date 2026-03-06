package main

//BAD Interface pollution - creating interfaces upfront with no real need.
//The developer defined an interface before any concrete type exists
//anticipating future use that may never come. This adds indirection with zero
//benefit right now

//Interface created speculatively, only one concrete type will ever use this
//There is no second implementation, no test double needed, no reason to abstect
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

// The only concrete implementation — this is all that exists.
type mysqlStore struct{} // MySQL storage implementation

// Save customer to storage
func (s *mysqlStore) StoreCustomer(customer Customer) error { return nil }

// Retrieve customer by ID
func (s *mysqlStore) GetCustomer(id string) (Customer, error) { return Customer{}, nil }

// Update existing customer
func (s *mysqlStore) UpdateCustomer(customer Customer) error { return nil }

// Return all customers
func (s *mysqlStore) GetAllCustomers() ([]Customer, error) { return nil, nil }

// Return customers without contracts
func (s *mysqlStore) GetCustomersWithoutContract() ([]Customer, error) { return nil, nil }

// Return customers with negative balance
func (s *mysqlStore) GetCustomersWithNegativeBalance() ([]Customer, error) { return nil, nil }

// ❌ The service depends on the interface, but there's only one implementation.
// Callers must navigate an abstraction layer that provides no value.
type CustomerService struct {
	store CustomerStorage // pointless abstraction — only mysqlStore ever goes here
}

// Constructor for CustomerService
func NewCustomerService(store CustomerStorage) CustomerService {
	return CustomerService{store: store}
}

// Create and store a new customer
func (cs CustomerService) CreateNewCustomer(id string) error {
	customer := Customer{ID: id}            // initialize customer
	return cs.store.StoreCustomer(customer) // store customer
}

func main() {} // program entry point
