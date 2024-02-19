package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/sultanlinjawi/first/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("PUT")
		// Expect the ID in the url
		regexpr := regexp.MustCompile(`/([0-9]+)`)
		group := regexpr.FindAllStringSubmatch(r.URL.Path, -1)
		// r.URL.Path
		if len(group) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(group[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		// [0] = first element that matches, [1] = second char in string (string is "/id")
		idString := group[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			p.l.Println("Invalid URI cannot convert to number")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.l.Println("Got id", id)

		p.updateProducts(id, rw, r)
	}

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	listProducts := data.GetProducts()
	err := listProducts.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		// return
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	product := &data.Product{}

	err := product.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to Unmarshal JSON", http.StatusBadRequest)
	}

	data.AddProduct(product)

	p.l.Printf("Prod: %#v", product)
}

func (p *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	product := &data.Product{}

	err := product.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to Unmarshal JSON", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, product)
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
