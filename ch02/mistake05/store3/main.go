package store3

//BAD: A constructor returns an interface instead of a concrete type
//This forces all callers to work through that abstrction and creates
//an awkward package dependency direction

// ❌ Returning an interface from a constructor.
// Any package that wants to call NewInMemoryStore now depends on
// wherever Store is defined. If Store lives in a client package,
// you get a cyclic dependency. If it lives here, the producer is
// dictating the consumer's abstraction — mistake06 all over again.
type Store interface {
	Save(id string, value int) error
	Load(id string) (int, error)
}

// Concrete in-memory store implementation
type InMemoryStore struct {
	data map[string]int // actual data storage
}

// ❌ Returns the interface — callers can't access any InMemoryStore-specific
// methods without a type assertion, and the dependency graph gets tangled.
func NewInMemoryStore() Store {
	return &InMemoryStore{data: make(map[string]int)} // initialize storage
}

// Save a key-value pair in memory
func (s *InMemoryStore) Save(id string, value int) error {
	s.data[id] = value
	return nil
}

// Load a value by key from memory
func (s *InMemoryStore) Load(id string) (int, error) {
	return s.data[id], nil
}

// ❌ Hypothetical extra method only on the concrete type.
// Callers using Store interface can NEVER reach this without a type assertion.
func (s *InMemoryStore) Dump() map[string]int {
	return s.data // return all data
}
