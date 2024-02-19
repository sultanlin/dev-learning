package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sultanlinjawi/first/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	productsHandler := handlers.NewProducts(l)

	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Get("/", productsHandler.GetProducts)

		r.Group(func(r chi.Router) {
			r.Use(productsHandler.MiddlewareValidateProduct)
			r.Post("/", productsHandler.AddProduct)
			r.Put("/{id:[0-9]+}", productsHandler.UpdateProducts)
		})
	})

	server := &http.Server{
		Addr:         ":3333",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	// http.ListenAndServe(":3333", sm)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	// server.ListenAndServe()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	l.Println("Received terminal, graceful Shutdown", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(timeoutContext)
}
