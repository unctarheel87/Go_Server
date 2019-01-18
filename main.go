package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Name string
	Age  int
}

func main() {
	addr := ":8080"

	http.HandleFunc("/api", apiHandler)

	http.Handle("/", http.FileServer(http.Dir("./src")))

	log.Println("Server is running on PORT", addr)

	http.ListenAndServe(addr, nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	data := person{Name: "John", Age: 31}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}
