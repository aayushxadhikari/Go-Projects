package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Note struct{
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedOn time.Time `json:"createdon"`
}

// thread safe map for storing notes

var (
	noteStore = make(map[string]Note) // making a map Note for storing strings
	id int 
	mu sync.RWMutex
)

// HTTP REQUEST - Create a new node

func PostNoteHandler(w http.ResponseWriter, r *http.Request){
	var note Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w,"Invalid request Payload", http.StatusBadRequest)
		return
	}
	note.CreatedOn = time.Now()

	mu.Lock() // read lock on NoteStore allows multiple reads prevents writing
	id++ 
	key := strconv.Itoa(id)
	noteStore[key] = note
	mu.Unlock() // ensures the lock is released after the function finishes preventing deadlocks

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)

}

// HTTP - GET - Retrieve all notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request){
	mu.RLock()

	defer mu.RUnlock()

	// creating a slice to hold notes
	var notes []Note
	for _, v := range noteStore{
		notes = append(notes, v)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
	
}

// HTTP - PUT :Update a note by ID
func PutNoteHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	var updateNote Note

	if err:= json.NewDecoder(r.Body).Decode(&updateNote);err!= nil{
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if note, exists := noteStore[key]; exists{
		updateNote.CreatedOn = note.CreatedOn
		noteStore[key] = updateNote
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updateNote)
	}else{
		http.Error(w, "Note not found", http.StatusNotFound)
	}
}

// HTTP - DELETE - Delete a note by ID
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	mu.Lock()
	defer mu.Unlock()

	if _, exists := noteStore[key]; exists{
		delete(noteStore, key)
		w.WriteHeader(http.StatusNoContent)
	}else{
		http.Error(w, "Note not Found", http.StatusNotFound)
	}
}

func main(){
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes",GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes",PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}",PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}",DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr: ":8080",
		Handler: r,
	}

	log.Println("Listening on Port 8080...")
	log.Fatal(server.ListenAndServe())
}