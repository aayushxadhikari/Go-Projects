package main

import (
	"fmt"
	"log"
	"net/http"
)

// Form Handler
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("ParseForm() error: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "POST request successful!\n")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)
}

// Hello Handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "Hello!")
}

func main() {
	// Serve static files from the "static" directory
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	// Register handlers
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// Start the server
	port := ":8080"
	fmt.Printf("🚀 Server running at http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server Error:", err)
	}
}
