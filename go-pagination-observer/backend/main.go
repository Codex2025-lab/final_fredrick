package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var products []Product

func seedData() {
	for i := 1; i <= 500; i++ {
		products = append(products, Product{
			ID:   i,
			Name: fmt.Sprintf("Product %d", i),
		})
	}
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 20
	}

	start := (page - 1) * limit
	end := start + limit

	if start > len(products) {
		start = len(products)
	}

	if end > len(products) {
		end = len(products)
	}

	response := map[string]interface{}{
		"page":  page,
		"limit": limit,
		"total": len(products),
		"data":  products[start:end],
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	seedData()

	http.HandleFunc("/products", productsHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}