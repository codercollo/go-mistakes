# MISTAKE-05 Interface Pollution

`What is an Interface`
An interface defines behavior - a set of method signatures a type must satisfy. Go interfaces are satisfied implicitly: no _implements_ keywords, no declaration. If a type has the methods, it satisfies the interface automatically

```go
type Reader interface {
  Read(p []byte) (n int, err error)
}
```

`Abstractions should be discovered, not created` - Don't design with interfaces upfront. Solve the real problem now. If an abstraction genuinely emerges from the code, introduce it then.

# When Interfaces Are Used:

1. _Common Behavior_
   Multiple unrelated types share the same operations
   factor it into an interface

```go
//sort.Interface - works for any index-based collection
type Interface interface{
 Len() int
 Less (i, j int) bool
 Swap(i, j int)
}
```

2. _Decoupling (for testing)_
   Replace a concrete dependency(eg a database) with a mockable abstraction

```go
//Without interface: forced to spin up a real MySQL instance to test
//If you want to write a unit test for AddCustomer, you have to connect to a real database
type CustomerService struct {
  store mysql.Store
}

//With interface: can inject a mock in unit tests
//Instead of depending directly on mysql.Store, define a small interface that only exposes the methods you need
type customerStorer interface {
  StoreCustomer(Customer) struct
}
type CustomerService struct {
  storer customerStorer
}
```

3. _Restricting Behavior_
   Expose only a subset of a type's
   capabilities to a consumer

```go
//IntConfig has both Get() and Set(), but this consumer should only read
type intConfigGetter interface {
  Get() int
}

//Bar() can only read - Set() is inaccessible by design
type Foo struct {
  threshold intConfigGetter
}
```

`The bigger the interface, the weaker the abstraction` - Rob Pike
Prefer small, focused interfaces. They compose easily and stay reusable

```go
//Standard library pattern - Combine small interfaces when needed
type ReadWriter interface {
  Reader
  Writer
}
```

# Interface Pollution

Creating interfaces speculatively, before
any concrete need exists. One implementation, no tests requiring a mock, no second consumer - yet an interface layer is added anyway.
_Why it happens_
Habit from C#/Java where _interface_ declarations are the norm before any implementation. Go works the other way around.
The problem:
-Adds an indirection layer with no payoff
-Makes the code harder to read and trace
-Abstractions created in anticipation are usually wrong anyway

```go
// ❌ One implementation exists. No second consumer. No test double needed.
// The interface adds a layer of indirection that helps nobody right now.
type CustomerStorage interface {
    StoreCustomer(customer Customer) error
    GetCustomer(id string) (Customer, error)
    UpdateCustomer(customer Customer) error
    GetAllCustomers() ([]Customer, error)
    GetCustomersWithoutContract() ([]Customer, error)
    GetCustomersWithNegativeBalance() ([]Customer, error)
}
```

Fix:

```go
// ✅ Use the concrete type directly. If a mock or second implementation
// is needed later, introduce the interface at that point.
type CustomerService struct {
    store *MySQLStore
}
```

Rule: `if  you can't point to a concrete reason and interface improves the code today remove it`

# Interface on the Producer Side

_Producer Side_ - Interface defined in the same package as the concrete implementation
_Consumer side_ - Interface defined in the package that uses the implementation

Defining an interface alongside its implementation and expecting all consummers to use it, is wrong; heres why:
-The producer is deciding the abstraction level for every consumer, but different consumers need different subsets
-It violates the Interface Segregation Principle: clients shouldn't depend on methods they don't use
-It creates awkward package dependencies(the producer knows about the consumer's abstraction shape)

```go
// store package — interface lives next to its implementation ❌
package store

type CustomerStorage interface { // 6 methods — every consumer gets all 6
    StoreCustomer(...) error
    GetCustomer(...) (Customer, error)
    UpdateCustomer(...) error
    GetAllCustomers() ([]Customer, error)
    GetCustomersWithoutContract() ([]Customer, error)
    GetCustomersWithNegativeBalance() ([]Customer, error)
}
```

Fix:

```go
// store package — expose only the concrete type ✅
package store
//Concrete type
type InMemoryStore struct {
	data map[string]int
}

// Constructor
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		data: make(map[string]int),
	}
}   // concrete return

// client package — define only what this package actually needs ✅
package client
type customersGetter interface {         // 1 method, unexported, perfectly scoped
    GetAllCustomers() ([]store.Customer, error)
}
```

_Constructor_ Is just a function that creates and initializes a struct.
_Concrete Type_ Is the actual implementation of something.

# Why this works in Go but not C#/Java`

Because Go interfaces are implicitly satisfied, `*store.InMemoryStore` satisfies `customersGetter` automatically, the store package has zero knowledge of the client's 3 interface. No circular dependecy, no explicit wiring.

# Exception

The standard library's encoding package defines interfaces (encoding.BinaryMarshaler etc.) on the producer side — but only because the language designers knew, not guessed, these abstractions would be universally useful.

# Rule

Interface on the consumer side by default. Producer-side interfaces only when you can prove the abstraction serves all known consumers well.

# Returning Interfaces

A constructor or factory function returns an interface type instead of a concrete struct

_Why is it wrong_
-Forces all callers into a fixed abstraction, they loase access to methods not in the interface
-To access concrete-only methods, callers need type assertions which is verbose and fragile
-Create a backwards dependency-the implementation package
references the package that owns the interface
-Can cause cyclic imports if the interface lives in a consumer package

```go
package store

// Interface
type Store interface {
	Save(id string, value int)
}
// ❌ Constructor returns the interface
func NewInMemoryStore() Store {
	return &InMemoryStore{data: make(map[string]int)}
}

func (s *InMemoryStore) Save(id string, value int) {
	s.data[id] = value
}

// Method NOT in the interface
func (s *InMemoryStore) Dump() map[string]int {
	return s.data
}

```

FIX:

```go
package store

//Concrete type
type InMemoryStore struct {
	data map[string]int
}

// ✅ Constructor returns the concrete type
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{data: make(map[string]int)}
}

//Methods
func (s *InMemoryStore) Save(id string, value int) {
	s.data[id] = value
}

func (s *InMemoryStore) Dump() map[string]int {
	return s.data
}
//---------------------------------------------------
package app

import "store"

// Consumer defines its own interface
type dataStore interface {
	Save(id string, value int)
	Load(id string) int
}

// Function accepts the interface
func Process(store dataStore) int {
	store.Save("a", 100)
	return store.Load("a")
}

```

\*InmemoryStore automatically satisfies because it has those methods
