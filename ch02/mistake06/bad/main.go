package main

import "fmt"

type Customer struct {
	Name string
}

type Contract struct {
	Title string
}

type Store struct{}

// ❌ PROBLEM: Using `any` as parameter and return type.
//
// `any` means "this function accepts literally ANY type".
// The compiler cannot enforce correctness anymore.
//
// This removes Go's biggest advantage: **compile-time type safety**.
//
// Callers cannot know what this function actually returns
// without reading documentation or source code.
func (s *Store) Get(id string) (any, error) {

	// ❌ The function returns `any`, so callers do not know
	// whether the result is a Customer, Contract, or something else.
	return Customer{Name: "Alice"}, nil
}

// ❌ PROBLEM: Parameter type is `any`
//
// The method accepts **anything**:
// int, string, bool, struct, slice, etc.
//
// The compiler cannot prevent incorrect usage.
func (s *Store) Set(id string, v any) error {

	// At runtime we simply print whatever was passed.
	fmt.Printf("storing: %v\n", v)

	return nil
}

func main() {
	s := Store{}

	// ❌ These are clearly wrong values for a "Store",
	// but the compiler cannot stop them because `any` accepts everything.

	_ = s.Set("foo", 42)         // int — accepted silently
	_ = s.Set("bar", true)       // bool — accepted silently
	_ = s.Set("baz", "a string") // string — accepted silently

	// ❌ Another problem: callers must perform a **type assertion**
	// to recover the real value.

	val, _ := s.Get("foo")

	// If the type assertion is wrong, the program can panic
	// at runtime instead of failing at compile time.
	customer, ok := val.(Customer) // "hope this is actually a Customer"

	if !ok {
		// This error only appears **at runtime**
		fmt.Println("wrong type — runtime surprise")
		return
	}

	fmt.Println(customer.Name)
}
