package main

import (
	"errors"
	"fmt"
	"net/http"
)

//Still not IDEAL : Builder Pattern
//Better than Config struct - no pointer variables for callers
//But: Error handing is delayed to Build(), and you still need
//to pass a (possibly empty) Config to NewServer

const defaultHTTPPort = 8080

// The final config the server acctually use
type Config struct {
	Port int
}

// Builder accumulates settings before creating Config
type ConfigBuilder struct {
	port *int //nil means "Caller never set it"
}

// Port is a chainable setter, Returns *ConfigBuilder so you can chain calls
func (b *ConfigBuilder) Port(port int) *ConfigBuilder {
	b.port = &port
	return b
}

// Build valdates everything and produces the real Config
// Validation is centralised here because chainable methods can't return errors
func (b *ConfigBuilder) Build() (Config, error) {
	cfg := Config{}

	switch {
	case b.port == nil:
		cfg.Port = defaultHTTPPort
	case *b.port < 0:
		return Config{}, errors.New("port must be positive")
	case *b.port == 0:
		cfg.Port = randomPort()
	default:
		cfg.Port = *b.port
	}

	return cfg, nil
}

func NewServer(addr string, cfg Config) (*http.Server, error) {
	fmt.Printf("Starting server at %s:%d\n", addr, cfg.Port)
	return &http.Server{Addr: fmt.Sprintf("%s:%d", addr, cfg.Port)}, nil
}

func randomPort() int { return 49152 }

func main() {
	//  Explicit port via builder — nicer than &port trick
	builder := ConfigBuilder{}
	builder.Port(8080)
	cfg, err := builder.Build()
	if err != nil {
		panic(err)
	}
	srv, err := NewServer("localhost", cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("server addr:", srv.Addr)

	//  Default config — still need an empty Config{}
	defaultCfg, _ := (&ConfigBuilder{}).Build()
	srv2, _ := NewServer("localhost", defaultCfg)
	fmt.Println("server addr:", srv2.Addr)
}
