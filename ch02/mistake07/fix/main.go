package main

import (
	"fmt"
	"io"
	"sort"
	"strconv"
)

// GOOD USE: #1: Generic function over maps
// One function works for any map type; no duplication, compile-time safe
// K mut be comparable (Go requirement for map keys). V can be anything
// Returns keys in their original type - no type assertions needed
func getKeys[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// GOOD USE #2: Custom type constraints
// Restrict type parameters to a set of allowed underlying types (~int, ~string)
// Useful when you want type safety but flexible inputs
type customInt int

func (i customInt) String() string {
	return strconv.Itoa(int(1))
}

// intOrString allows only int or string (or types with same underlying type)
type intOrString interface {
	~int | ~string
}

func printKey[K intOrString](k K) {
	fmt.Println(k)
}

// GOOD USE #3: Generic data structure
// Linked list for any value type T
type Node[T any] struct {
	Val  T
	next *Node[T]
}

func (n *Node[T]) Add(next *Node[T]) {
	n.next = next
}

// GOOD USE #4: Factor out behavior with generics
// Generic sort wrapper - avoid one sort function per type
type SliceFn[T any] struct {
	S       []T
	Compare func(T, T) bool
}

func (s SliceFn[T]) Len() int           { return len(s.S) }
func (s SliceFn[T]) Less(i, j int) bool { return s.Compare(s.S[i], s.S[j]) }
func (s SliceFn[T]) Swap(i, j int)      { s.S[i], s.S[j] = s.S[j], s.S[i] }

// GOOD USE #5: Merge two channels of any type
func merge[T any](ch1, ch2 <-chan T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		for v := range ch1 {
			out <- v
		}
		for v := range ch2 {
			out <- v
		}
	}()
	return out
}

// BAD GENERIC — using generics where unnecessary
// Just calling a method of the type parameter adds no value.
// Simpler to accept io.Writer directly.

// func writeData[T io.Writer](w T, data []byte) { w.Write(data) }

// CORRECT: accept interface directly — simpler, clearer
func writeData(w io.Writer, data []byte) {
	_, _ = w.Write(data)
}

func main() {
	// ✅ getKeys works with any map — types inferred automatically
	m1 := map[string]int{"one": 1, "two": 2}
	keys1 := getKeys(m1) // []string — no type assertion needed
	fmt.Println(keys1)

	m2 := map[int]string{1: "one", 2: "two"}
	keys2 := getKeys(m2) // []int — correct type returned
	fmt.Println(keys2)

	// ✅ Custom constraint — only int or string keys allowed
	printKey("hello")
	printKey(customInt(42))

	// ✅ Generic linked list
	n1 := &Node[int]{Val: 1}
	n2 := &Node[int]{Val: 2}
	n1.Add(n2)
	fmt.Println(n1.Val, n1.next.Val)

	// ✅ Generic sort — one implementation works for any type
	s := SliceFn[int]{
		S:       []int{3, 1, 2},
		Compare: func(a, b int) bool { return a < b },
	}
	sort.Sort(s)
	fmt.Println(s.S) // [1 2 3]

	// ✅ writeData — simple, no unnecessary generics
	writeData(io.Discard, []byte("hello"))
}
