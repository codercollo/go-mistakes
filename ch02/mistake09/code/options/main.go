package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

//BEST: Functional Options Pattern
//Idiomatic Go, used by gRPC and many stdlib libraries

//Core idea:
//Keep an unexported `options` struct ("the bag of settings")
//Each setting is a function: type Option func(*Options) error
//Public With* functions return a pre-filled Option closure
//NewServer takes ...Option (variadic) so callers pass 0 or more options

//Benefits:
//Zero-value default: NewServer("localhost") just works
//Each option validates itself immediately no delayed Build() needed
//Adding new optins never breaks existing callers
//Highly readable at the call site

const defaultHTTPPort = 8080

// The Unexported settings bag
type options struct {
	port    *int          // nil = not set by caller
	timeout time.Duration // 0 = not set by caller
}

// The Option type - just a function that mutates *options
type Option func(*options) error

// Public With* constructors - one per configuration field
// WithPort validates the port and stores it in options
// WithPort a closure lets WithPort capture the `port` argument cleanly
func WithPort(port int) Option {
	return func(o *options) error {
		if port < 0 {
			return errors.New("port must be positive")
		}
		o.port = &port
		return nil
	}
}

// WithTimeout is another option - adding it didn't touch any existing code
func WithTimeout(d time.Duration) Option {
	return func(o *options) error {
		if d < 0 {
			return errors.New("timeout must be non-negative")
		}
		o.timeout = d
		return nil
	}
}

// NewServer - accepts variadic Option, applies each one
func NewServer(addr string, opts ...Option) (*http.Server, error) {
	//Start with zero-value options (all nil/zero)
	var o options

	//Apply every option the caller passed
	for _, opt := range opts {
		if err := opt(&o); err != nil {
			return nil, err //one bad option stops everything
		}
	}

	//Resolve final port value using our business rules
	var port int
	switch {
	case o.port == nil:
		port = defaultHTTPPort //caller passed nothing > default
	case *o.port == 0:
		port = randomPort() //caller explicitly wants a random port
	default:
		port = *o.port
	}
	fmt.Printf("Startung server at %s:%d(timeout-%s)\n", addr, port, o.timeout)
	return &http.Server{
		Addr:         fmt.Sprintf("%s:%d", addr, port),
		WriteTimeout: o.timeout,
	}, nil
}

func randomPort() int { return 49152 }

//Caller examples

func main() {
	//Explicit port + timeout - reads like a sentence
	srv, err := NewServer("localhost",
		WithPort(9090),
		WithTimeout(5*time.Second),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("addr", srv.Addr)

	//Default everything - no empty struct, no nil, just omit options
	srv2, err := NewServer("localhost")
	if err != nil {
		panic(err)
	}
	fmt.Println("addr", srv2.Addr)

	//Random Port
	srv3, err := NewServer("localhost", WithPort(0))
	if err != nil {
		panic(err)
	}
	fmt.Println("addr", srv3.Addr)

	//Invalid port - error surfaced immediately, not at Build() time
	_, err = NewServer("localhost", WithPort(-1))
	fmt.Println("expected error", err)
}
