# `any` Says Nothing

_What is any_
any - is an alias for the empty interface
interface{} - an interface with zero methods.

```go
// These are identical
var i interface{}
var i any
```

Because any has no methods, every type in Go satisfies it: int, string, structs, etc

```go
var i any
i = 42
i = "foo"
i = struct{ s string }{s: "bar"}
i = someFunc
```

_The Problem_
Accepting or returning any loses all type information
-When you assign to any, Go forgets the original type. To
use the value again, you need to type assert.

```go
val, _ := store.Get("foo")
customer, ok := val.(Customer) // hope the type is right — runtime panic if wrong
```

This throws away one of Go's biggest strengths: static typing.

# Using `any` in Method Signatures

Using any as a parameter or return type
when the actual types are known just to avoid writing separate methods
_The Problem_
-Method signatures convey no information, callers must read docs or source code to understand valid types
-No compile-time safeguard wrong types (eg: passing int where Customer is expected) silently compile
-Callers need type assertions to use return values

- You lose the benefits of Go
  WRONG:

```go
// ❌ What does Get return? What does Set accept? Nobody knows from the signature.
func (s *Store) Get(id string) (any, error) { ... }
func (s *Store) Set(id string, v any) error { ... }

// ❌ Compiler won't stop this — wrong types accepted silently
s.Set("foo", 42)       // int
s.Set("bar", true)     // bool
s.Set("baz", "oops")   // string

```

FIX:

```go
// ✅ One explicit method per type — expressive, safe, self-documenting
func (s *Store) GetCustomer(id string) (Customer, error)          { ... }
func (s *Store) SetCustomer(id string, customer Customer) error   { ... }

func (s *Store) GetContract(id string) (Contract, error)          { ... }
func (s *Store) SetContract(id string, contract Contract) error   { ... }
```

More methods isn't a problem — consumers can define their own minimal interface if needed:

```go
// Consumer only needs contract methods — define a scoped interface
type ContractStorer interface {
    GetContract(id string) (store.Contract, error)
    SetContract(id string, contract store.Contract) error
}
```

RULE: Make signatures as explicit as possible. A little duplication is worth it for expressiveness and type safety.

# WHEN `any` is Acceptable

any is valid when there is a genuine need to handle truly unknown types typically in generic infrastucture code eg:encoding/json and database/sql packages

```go
// ✅ Can marshal any type by design — any is correct here
func Marshal(v any) ([]byte, error) { ... }

// ✅ Query parameters could be any type — any is correct here
func (c *Conn) QueryContext(ctx context.Context, query string, args ...any) (*Rows, error) { ... }
```

This are justified because the functions genuinely cannot know the type in advance. Your application code almost never has this problem, the types are known just write them out
