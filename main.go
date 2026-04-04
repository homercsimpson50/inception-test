package main

import (
	"encoding/json"
	"net/http"
)

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var items = []Item{
	{ID: 1, Name: "Widget", Price: 999},
	{ID: 2, Name: "Gadget", Price: 2499},
	{ID: 3, Name: "Doohickey", Price: 149},
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/items", itemsHandler)
	http.ListenAndServe(":8080", nil)
}
