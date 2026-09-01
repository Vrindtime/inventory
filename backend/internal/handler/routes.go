package handler

import (
	"fmt"
	"net/http"
)

func DemoHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w,"Demo Handler")
}

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"pong")
	})

	//Auth
	mux.HandleFunc("POST /api/login",DemoHandler)
	mux.HandleFunc("POST /api/signup",DemoHandler) //In a way will also call create User

	// Users API
	mux.HandleFunc("GET /api/users", DemoHandler)
	mux.HandleFunc("GET /api/users/{id}", DemoHandler)
	mux.HandleFunc("POST /api/users", DemoHandler)
	mux.HandleFunc("PUT /api/users/{id}", DemoHandler)
	mux.HandleFunc("DELETE /api/users/{id}", DemoHandler)
	mux.HandleFunc("GET /api/users/{id}/activity", DemoHandler)
	mux.HandleFunc("POST /api/users/{id}/activity", DemoHandler)

	// Products
	mux.HandleFunc("GET /api/products", DemoHandler)
	mux.HandleFunc("GET /api/products/{id}", DemoHandler) // will fetch data from stock
	mux.HandleFunc("POST /api/products", DemoHandler)     // will also create stock as well stock log
	mux.HandleFunc("PUT /api/products/{id}", DemoHandler)
	mux.HandleFunc("DELETE /api/products/{id}", DemoHandler)

	// Stock
	mux.HandleFunc("PUT /api/stocks/{id}", DemoHandler) // Increment or Decrement a stock value, this will affect stock log as well

	// Warehouses
	mux.HandleFunc("GET /api/warehouses", DemoHandler)
	mux.HandleFunc("GET /api/warehouses/{id}", DemoHandler)
	mux.HandleFunc("POST /api/warehouses", DemoHandler)
	mux.HandleFunc("PUT /api/warehouses/{id}", DemoHandler)
	mux.HandleFunc("DELETE /api/warehouses/{id}", DemoHandler)
	mux.HandleFunc("GET /api/warehouses/{id}/dashboard", DemoHandler) // Basically Reports we can do like ?=sales, ?=expense etc.

	// Customers
	mux.HandleFunc("GET /api/customers", DemoHandler)
	mux.HandleFunc("GET /api/customers/{id}", DemoHandler)
	mux.HandleFunc("POST /api/customers", DemoHandler)
	mux.HandleFunc("PUT /api/customers/{id}", DemoHandler)
	mux.HandleFunc("DELETE /api/customers/{id}", DemoHandler)
	mux.HandleFunc("GET /api/customers/{id}/dashboard", DemoHandler)
	mux.HandleFunc("GET /api/customers/{id}/history", DemoHandler)

	// Orders
	mux.HandleFunc("GET /api/orders", DemoHandler)
	mux.HandleFunc("GET /api/orders/{id}", DemoHandler) // Shows Order Status as well
	mux.HandleFunc("POST /api/orders", DemoHandler)     // the form in here will call from products and stock and in case a stock is low we call allocation
	mux.HandleFunc("PUT /api/orders/{id}", DemoHandler) // have to create a middleware to let only admin void or edit this, the void will be recorded to order log as well
	mux.HandleFunc("PATCH /api/orders/{id}", DemoHandler) // for status changes
	mux.HandleFunc("DELETE /api/orders/{id}", DemoHandler)
	mux.HandleFunc("GET /api/orders/{id}/dashboard", DemoHandler)
	mux.HandleFunc("GET /api/orders/{id}/history", DemoHandler)

	// Allocations
	mux.HandleFunc("GET /api/allocations/history", DemoHandler)
}
