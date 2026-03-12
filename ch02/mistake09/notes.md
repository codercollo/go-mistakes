# Mistake 9: Not Using the Functional Options Pattern

The Problem:
-You're building a library with a function like:

```go
func NewServer(addr string, port int) (\*http.Server, error)
```

-Users are happy, until they need options(timeouts, TLS etc) Adding new parameters breaks every existing caller. You need a better design

# Three Approaches

1 `Config Struct`

```go
type Config struct { Port *int }
func NewServer(addr string, cfg Config) (*http.Server, error)
```

What it fixes: Adding fields to Config doesn't break old callers
The pointer trick: Go zero-initialises int to 0 so you can't tell if the caller set the port to 0 or just forgot to set it. Using *int lets you check nil = "not set"
*Downsides\*
-Callers have to create a throwaway variable:

```go
 port := 0;
 cfg := Config{Port : &port}
```

- Empty default call looks magic

```go
NewServer("localhost", Config{})
```

2 `Builder Pattern`

```go
builder := ConfigBuilder{}
builder.Port(8080)
cfg, _ := builder.Build()
NewServer("localhost", cfg)
```

What it fixes: No pointer variables for callers. Chainable setters look clean
_Downsides_
Errors can only surface in Build() - individual setter methods can't return errors
without breaking method chaining
Still need to pass a Config(even if empty) to NewServer

3 `Functional Options Pattern`
The idiomatic Go answer: Used by gRPC, many std lib wrappers
How it works: three moving parts:

(a) _An unexported settings bag_

```go
type options struct {
    port    *int
    timeout time.Duration
}
```

Unexported = callers can't touch it directly. Only your With\* functions mutate it

(b)`An Option type`
Literally just a function

```go
type Option func(*options) error
```

Each option is a function that reveals the bag and modifies one field

(c) `Public With* constructors return a pre-baked Option`

```go
func WithPort (port int) Option {
  return func(o *options) error { //this is a closure
    if port < 0 {
      return errors.New("port must be positive")
    }
    o.port = &port
    return nil
  }
}
//Closure - an anonymous function that "closes over" (remembers) variables from
//its surrounding scope, here it remembers port even after WithPort returns.
```

(d) NewServer takes variadic ...Option

```go
func NewServer(addr string, opts ..Option)(*http.Server, error) {
  var o options
  for _, opt := range opts {
    if err := opt(&o); err != nil {
      return nil, err
    }
  }
}
```

- Use ...Option (functional options) for any Go API that has optional configuration.
  Each option is a With* function that returns a closure. Callers pass only what they need;
  everything else falls back to sensible defaults. No empty structs, no pointer gymnastics. *

-Variadic (...Option)
A variadic function parameter that allows the function to accept zero, one, or multiple arguments of the specified type.

-Closure
A function that captures and remembers variables from the scope where it was created, allowing it to use those variables even after the outer function has finished executing

-Unexported Struct
A struct whose name starts with a lowercase letter, making it private to the package so external packages cannot directly create or access it.

-Zero Value
The default value Go assigns to a variable when it is declared but not explicitly initialized (e.g., 0 for numbers, "" for strings, nil for pointers, slices, maps, etc.).

-Functional Options Pattern
A design pattern where configuration options are implemented as functions, and these functions are passed as arguments to modify how an object is created or configured.

-Builder Pattern
A design pattern where a separate builder object collects configuration settings step by step, and then constructs the final object when building is complete.
