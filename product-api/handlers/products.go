package handlers

import (
	"log"
	"net/http"
	"github.com/jedrekdomanski/microservices/product-api/data"
)

// Products is a http.Handler
type Products struct {
	logger *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(logger *log.Logger) *Products {
	return &Products { logger }
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// interface
func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	// catch all
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
	// handle the request for a list of products
	productList := data.GetProducts()
	err := productList.ToJSON(w)

	if err != nil {
		http.Error(w, "Unable to parse JSON", http.StatusInternalServerError)
	}
}
