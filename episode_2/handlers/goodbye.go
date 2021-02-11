package handlers

import (
	"log"
	"net/http"
)

// Goodbye is a simple handler
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye create a new goodbye handler with the given logger
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

// ServeHTTP implements the go http.handler interface
// https://golang.org/pkh/net/http/#Handler
func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("GoodBye"))
}
