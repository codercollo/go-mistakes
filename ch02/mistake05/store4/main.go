package store4

//FIX: Return the concrete types from constructors
//Callers get full access to all methods. They can define their own
//interface on the consumer side if abstraction is needed

// Concrete in-memory store
type InMemoryStore struct {
	data map[string]int // actual key-value storage
}

// ✅ Returns the concrete *InMemoryStore.
// Callers that want an interface can define one themselves.
// Callers that don't need one aren't forced into indirection.
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{data: make(map[string]int)} // initialize storage
}

// Load a value by key
func (s *InMemoryStore) Load(id string) (int, error) {
	return s.data[id], nil
}

//Save a key-value pair (matches dataStore interface)
func (s *InMemoryStore) Save(id string, value int) error {
	s.data[id] = value
	return nil
}

// ✅ Callers using the concrete type can access this directly — no type assertion needed.
func (s *InMemoryStore) Dump() map[string]int {
	return s.data // return all data
}
