# GENERICS : Being Confused About When to Use Generics

_Generics_
Generics let you write functions and types where the type is specified later,
at the point of use not when you write the code

```go
// T is a type parameter — the caller decides what T is
func foo[T any](t T) { ... }
foo[int](42)     // T = int
foo[string]("hi") // T = string
```

The type is resolved at compile time, full type safety no runtime overhead

# The Problem Generics Solves

Before generic, writing one function that works across multiple map types required
`any` + type switches

```go
// ❌ Pre-generics: boilerplate, runtime errors, forced type assertions
func getKeys(m any) ([]any, error) {
    switch t := m.(type) {
    case map[string]int:
        // ... range loop
    case map[int]string:
        // ... same loop duplicated
    default:
        return nil, fmt.Errorf("unknown type: %T", t) // runtime surprise
    }
}
```

Problems:

1. Duplicated Logic for every new map type
2. Type errors caught at runtime, not compile time
3. Caller must type-assert every returned value

# Fix : Type Parameters

```go
// ✅ One function, any map type, compile-time safe
func getKeys[K comparable, V any](m map[K]V) []K {
    var keys []K
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}
keys := getKeys(map[string]int{"a": 1}) // []string — no assertion needed
keys := getKeys(map[int]string{1: "a"}) // []int   — correct type returned
```

`K comparable` map keys must be comparable (==, !=) slices can't be map keys
`V any` values can be any type
Go infers the type argument — you rarely need to write getKeys[string](m) explicitly

# Constraints

A constraint restricts what types a type parameter can be. It's just an interface

```go
// Constraint using a type set — only int or string allowed
type intOrString interface {
    ~int | ~string
}

func printKey[K intOrString](k K) { fmt.Println(k) }
```

`int` -Only the exact built-in int type
`~int` int plus any custom type whose underlying type is int

```go
type customInt int // underlying type is int

// ~int accepts customInt ✅
// int  rejects customInt ❌ — compilation error
```

# Generic Data Structures

Type parameters work on structs too — useful for collections like linked lists, trees, heaps:

```go
type Node[T any] struct {
    Val  T
    next *Node[T]
}

func (n *Node[T]) Add(next *Node[T]) {
    n.next = next
}

n := &Node[int]{Val: 1}
```

_Note: methods cannot have their own type parameters — generics go on the receiver (the struct), not the method:_

```go
// ❌ Won't compile
func (Foo) bar[T any](t T) {}

// ✅ Correct — type parameter on the struct
type Foo[T any] struct{}
func (f Foo[T]) bar(t T) {}
```

# When TO Use Generics

1. Generic data structures
   Any time you'd write a binary tree, linked list, heap, or stack — use a type parameter instead of hardcoding the element type.

2. Functions over slices, maps, or channels of any type

```go
// ✅ Merge any two channels of the same type
func merge[T any](ch1, ch2 <-chan T) <-chan T { ... }
```

3. Factoring out behavior (not just types)

```go
// ✅ One sort wrapper for any type — avoids one function per type
type SliceFn[T any] struct {
    S       []T
    Compare func(T, T) bool
}

func (s SliceFn[T]) Len() int           { return len(s.S) }
func (s SliceFn[T]) Less(i, j int) bool { return s.Compare(s.S[i], s.S[j]) }
func (s SliceFn[T]) Swap(i, j int)      { s.S[i], s.S[j] = s.S[j], s.S[i] }
```

# When Not to Use Generics

1. When you just call a method on the type parameter

```go
// ❌ Generics add nothing here — T is only used to call Write
func writeData[T io.Writer](w T, data []byte) {
    w.Write(data)
}

// ✅ Just use the interface directly — simpler, clearer
func writeData(w io.Writer, data []byte) {
    w.Write(data)
}
```

If you only need a type to call its methods, use an interface — not generics.

2. When it makes the code more complex
   Generics are an abstraction. Like interfaces, unnecessary abstractions make code harder to read. If the generic version isn't obviously cleaner than the concrete version, skip it.

# Generics vs. Interfaces

1. If you need to call methods on a type, use an interface

```go
// Define interface
type Stringer interface {
	String() string
}

// Implement for multiple types
type Person struct{ Name string }
func (p Person) String() string { return p.Name }

type Animal struct{ Species string }
func (a Animal) String() string { return a.Species }

// Function works on any type that implements Stringer
func printString(s Stringer) {
	fmt.Println(s.String())
}

func main() {
	printString(Person{"Alice"})
	printString(Animal{"Dog"})
}
//Explanation: You only care about the behavior (String()), not the underlying type.
```

2. If you need to work with the type itself(store it, return it, pass it around) use generics

```go
// Generic function to return the first element of a slice
func first[T any](slice []T) T {
	return slice[0]
}

func main() {
	ints := []int{1, 2, 3}
	strs := []string{"a", "b", "c"}

	fmt.Println(first(ints)) // 1
	fmt.Println(first(strs)) // "a"
}
//Explanation: Generics preserve the actual type (int, string), no type assertions needed.
```

3. If you have one implementation but many types sharing behavior, use an interface

4. If you have one implementation but many types sharing structure, use generics
