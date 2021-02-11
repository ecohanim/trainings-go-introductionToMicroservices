package handlers

import (
	"log"
	"net/http"

	"../data"
)

// Products is a http.handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with a given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and satisfied the http.handler interface
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	// handle get
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// catch all
	// if no method satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data store
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the data store
	lp := data.GetProducts()
	// d, err := json.Marshal(lp)

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	// rw.Write(d)
}
