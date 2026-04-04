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

func versionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"version":  "1.0.0",
		"built_by": "containerized-gt",
	})
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/items", itemsHandler)
	http.HandleFunc("/version", versionHandler)
	http.ListenAndServe(":8080", nil)
}
