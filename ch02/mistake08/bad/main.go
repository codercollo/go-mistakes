package main

import (
	"fmt"

	"github.com/codercollo/go-mistakes/ch02/mistake08/bad/inmem"
)

func main() {
	//PROBLEM 1: sync.Mutex is embedded in InMem
	//This unintentionally exposes Lock/Unlock to external callers
	m := inmem.New()
	m.Lock()
	m.Unlock()

	_ = m
	fmt.Println("External code can call m.Lock() - this should Not be possible")
}
