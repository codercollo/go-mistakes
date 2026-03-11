package main

import (
	"fmt"
	"os"

	"github.com/codercollo/go-mistakes/ch02/mistake08/fix/inmem"
	"github.com/codercollo/go-mistakes/ch02/mistake08/fix/logger"
)

func main() {
	//FIX 1: Mutex is now a named (unexported) field
	//External callers can no longer call Lock and Unlock, it's fully encapsulated
	m := inmem.New()
	m.Set("score", 42)

	val, ok := m.Get("score")
	fmt.Printf("inmem.Get(\"score\") = %d, found=%v\n", val, ok)

	//FIX 2: Embedding io.Writecloser in logger is appropriate here
	//Write and Close are *meant* to be public, embedding just removes boilerplate
	l := logger.New(os.Stdout)
	_, _ = l.Write([]byte("hello from logger\n"))
	_ = l.Close()

}
