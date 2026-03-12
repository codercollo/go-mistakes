package main

import (
	"fmt"
	"net/http"
)

//PROBLEM: Config struct with pointer field- works but is clunky for callers
//Why pointers? Because Go zero-initializes int to o, so we can't tell if the
//caller meant "use port 0 " vs "I didn't set a port" A *int lets us check nil

const defaultHTTPPort = 8080

// pointer so we can detect "not set" (nil) vs "set to 0"
type Config struct {
	Port *int
}

func NewServer(addr string, cfg Config) (*http.Server, error) {
	var port int

	if cfg.Port == nil {
		port = defaultHTTPPort //Not set - use default
	} else if *cfg.Port < 0 {
		return nil, fmt.Errorf("port must be positive, got %d", *cfg.Port)
	} else if *cfg.Port == 0 {
		port = randomPort()
	} else {
		port = *cfg.Port
	}

	fmt.Printf("Starting server at %s:%d\n", addr, port)
	return &http.Server{
		Addr: fmt.Sprintf("%s:%d", addr, port)}, nil
}
func randomPort() int { return 49152 }

func main() {
	//  Explicit port — caller must create a var first (annoying)
	port := 8080
	srv, err := NewServer("localhost", Config{Port: &port})
	if err != nil {
		panic(err)
	}
	fmt.Println("server addr:", srv.Addr)

	//  Default config — empty struct is fine but looks weird to readers
	srv2, err := NewServer("localhost", Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("server addr:", srv2.Addr)
}
