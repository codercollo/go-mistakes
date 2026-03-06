package app1

import (
	"github.com/codercollo/go-mistakes/ch02/mistake05/store4"
)

//The consumer defines its own interface only if abstrction is actually needed
// eg: for mocking in tests. The producer doesn't force this

// ✅ Consumer defines a minimal interface for its own needs (e.g. unit testing).
// The producer (store package) doesn't know or care about this.
type dataStore interface {
	Save(id string, value int) error // only what this package needs
	Load(id string) (int, error)
}

// Application logic
type App struct {
	store dataStore // injected dependency
}

// Constructor for App
func NewApp(s dataStore) App {
	return App{store: s} // dependency injection
}

// Process saves a value and returns it
func (a App) Process(id string, value int) (int, error) {
	if err := a.store.Save(id, value); err != nil { // save via interface
		return 0, err
	}
	return a.store.Load(id) // load via interface
}

// Wire-up: *store.InMemoryStore satisfies dataStore implicitly — no explicit declaration needed.
func main() {
	s := store4.NewInMemoryStore() // concrete type returned
	app := NewApp(s)               // implicitly satisfies dataStore

	v, _ := app.Process("key1", 42) // use app through interface
	_ = v

	// Also works directly on the concrete type when no abstraction needed:
	_ = s.Dump() // accessible because we hold the concrete type
}
