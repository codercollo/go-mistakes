# Mistake 08 - Type Embedding Pitfalls

_Type Embedding_
A struct field declared without a name is an embedded field
The embedded type's fields and methods get promoted to the outer struct

```go
//Embedded - no name
type Foo struct {
  Bar
}

type Bar struct {
  Baz int
}

foo := Foo{}
foo.Baz = 42 //promoted - works directy
foo.Bar.Baz = 42 //also works

```

_Bad Embedding: `sync.Mutex`_

```go
//bad/inmem/inmem.go ; embedded
type InMem struct {
  sync.Mutex
  m map[string]int
}
```

_What goes wrong_
sync.Mutex - is embedde, so Lock() and Unlocke() are promoted to InMem.
External callers can now call them directly, which they should never be able to do

```go
m := inmem.New()
m.Lock  //external code can now break your locking logic
```

_FIX: Give the mutex a name, keeps it unexported_

```go
//fix/inmem/inmem.go
type InMem struct {
  mu sync.Mutex  // named + unexported — invisible outside the package
  m map[string]int
}
```

# Good, Embedding `io.WriteCloser`

```go
type Logger struct {
  io.WriterCloser  // embedding is intentional here
}
```

`Why this is fine`
Write and Close are meant to be public. Embedding just removes the need to write
forwading methods ie:

```go
// Without embedding you'd need this boilerplate:
func (l Logger) Write(p []byte) (int, error) { return l.wc.Write(p) }
func (l Logger) Close() error                { return l.wc.Close() }
```

Embedding also makes logger automatically satisfy io.WriteCloser

# Embedding vs OOP Inheritance — one key difference

With embedding, the receiver of a method stays the inner type:

```go
// X is embedded in Y
// When Foo() is called on Y, the receiver is still X — not Y
y.Foo()  // receiver inside Foo = X


type Engine struct{}

func (e *Engine) Start() {
	fmt.Printf("Start called on receiver type: %T\n", e)
}

type Car struct {
	Engine // embedded
}

```

With OOP subclassing, the receiver becomes the subclass Y.
Embedding is composition, not inheritance.

```java

class Engine {
    void start() {
        System.out.println(this.getClass()); // 'this' becomes Car because Car extends Engine
    }
}

class Car extends Engine {}

public class Main {
    public static void main(String[] args) {
        Car c = new Car();
        c.start();
    }
}
```
