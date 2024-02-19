package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sultanlinjawi/first/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	listProducts := data.GetProducts()
	err := listProducts.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		// return
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	product := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&product)

	p.l.Printf("Prod: %#v", product)
}

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	idString := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(rw, "Bad Parameter Value", http.StatusBadRequest)
		return
	}

	product := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &product)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product Not Found", http.StatusInternalServerError)
		return
	}

	p.l.Printf("Prod: %#v", product)
}

type KeyProduct struct{}

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		product := data.Product{}

		err := product.FromJSON(r.Body)

		if err != nil {
			p.l.Println("[ERROR] deserializating product", err)
			http.Error(rw, "Unable to Unmarshal JSON", http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, product)
		req := r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, req)
	})
}
