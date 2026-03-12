// Package stringset provides set operations on strings
// Named after what it provides, not what it contains (avoids: "util", "common", etc)
package stringset

import "sort"

//Set is the core type a map with empty struct values (no memory overhead).
type Set map[string]struct{}

// New creates a Set from the given strings.
func New(vals ...string) Set {
	s := make(Set, len(vals))
	for _, v := range vals {
		s[v] = struct{}{}
	}
	return s
}

// Sort returns the set elements in sorted order.
func (s Set) Sort() []string {
	out := make([]string, 0, len(s))
	for k := range s {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
