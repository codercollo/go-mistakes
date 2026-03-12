# Go Mistakes: #12–14 — Project & Package Organization

_Project Misorganization_

- Recommended Layout (project-layout)
  /cmd/foo/main.go ← app entry points (one dir per binary)
  /internal/ ← private code, not importable externally
  /pkg/ ← public libraries for others to import
  /test/ ← integration/API tests (unit tests live next to source)
  /configs/ ← config files
  /api/ ← Swagger, Protobuf, etc.
  /docs/ ← design docs
  /examples/ ← usage examples
  /build/ ← CI, packaging
  /scripts/ ← tooling scripts
  /vendor/ ← vendored deps
  `No /src — too generic. Go favors intent-revealing top-level dirs.`

_Go Package Design Rules_
-Avoid premature packaging
Start simple and refactor the package structure as the project grows.
-Avoid nano packages
Do not create a package for every 1–2 files unless there is a clear responsibility.
-Avoid huge packages
Very large packages dilute meaning. Split them based on cohesive functionality.
-Name packages by what they provide
Example: stringset instead of stringutils.
-Use single lowercase words for package names
Keep names short, concise, and expressive.
-Minimize exports
Default to unexported identifiers. Export only what must be accessed outside the package.
-Stay consistent
An imperfect but consistent structure is better than an inconsistent one.
`Organize by context (customer, contract) or layer (hex architecture) — pick one and stick with it.`

# Creating Utility Packages

The Problem:

```go
//BAD — meaningless package name
package util
func NewStringSet(...string) map[string]struct{} { ... }
func SortStringSet(map[string]struct{}) []string  { ... }

// client
set := util.NewStringSet("c", "a", "b")
util.SortStringSet(set)
```

-util, common, base, shared → meaningless names that reveal nothing.

The Fix — Expressive Package + Type

```go
// ✅ GOOD — package named after what it provides
package stringset

type Set map[string]struct{}

func New(...string) Set      { ... }
func (s Set) Sort() []string { ... }

// client:
set := stringset.New("c", "a", "b")
set.Sort()
```

Package name reads like a sentence: stringset.New, stringset.Sort
Only one reference to the package on the client
Exposes a real type instead of loose functions
``if client, server, and shared types all exist — consider merging them into one package rather than a common dump.`

# Package Name Collisions

The Problem

```go
BAD — variable shadows the package
redis := redis.NewClient()   // redis package now inaccessible in this scope
redis.Get("foo")             // ambiguous: var or package?
```

Fix 1: Rename the Variable

```go
//  Clearest solution
redisClient := redis.NewClient()
redisClient.Get("foo")
```

FIX 2: Import Alias

```go
//When you must keep the variable name
import redisapi "mylib/redis"

redis := redisapi.NewClient()
redis.Get("foo")

```

Also Avoid Built-in Shadowing

```go
//  Shadows built-in copy()
copy := copyFile(src, dst)
```

`Avoid dot imports (import . "pkg") — they remove the qualifier entirely and cause confusion.`
