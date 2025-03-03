package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"notesapp/handlers"
)

func main() {
	r := mux.NewRouter().StrictSlash(false)

	// Define routes
	r.HandleFunc("/api/notes", handlers.GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", handlers.PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", handlers.PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", handlers.DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Listening on Port 8080...")
	log.Fatal(server.ListenAndServe())
}
