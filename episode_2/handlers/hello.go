package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a simple handler
type Hello struct {
	l *log.Logger
}

// NewHello create a new hello handler with the given logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServeHTTP implements the go http.handler interface
// https://golang.org/pkh/net/http/#Handler
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Helo World")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadRequest)
		return
	}

	// write the response
	fmt.Fprintf(rw, "Hello %s", d)

}
