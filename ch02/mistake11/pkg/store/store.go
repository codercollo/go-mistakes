package store

// DefaultPermission is the default permission used by the store engine.
const DefaultPermission = 0o644 // Need read and write accesses.

// Customer is a customer representation.
type Customer struct {
	id   string
	name string
}

// ID returns the customer identifier.
func (c Customer) ID() string { return c.id }

// Name returns the customer's full name.
func (c Customer) Name() string { return c.name }

// NewCustomer creates a new Customer with the given id and name.
func NewCustomer(id, name string) Customer {
	return Customer{id: id, name: name}
}

// ComputePath returns the fastest path between two points.
// Deprecated: This function uses a deprecated algorithm. Use ComputeFastestPath instead.
func ComputePath() string { return "" }

// ComputeFastestPath returns the optimal path between two points using the A* algorithm.
func ComputeFastestPath() string { return "" }
