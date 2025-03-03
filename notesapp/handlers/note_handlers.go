package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"notesapp/models"
	"notesapp/store"
)

// POST /api/notes - Create a new note
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdNote := store.AddNote(note)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdNote)
}

// GET /api/notes - Retrieve all notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	notes := store.GetAllNotes()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}

// PUT /api/notes/{id} - Update a note by ID
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedNote models.Note
	if err := json.NewDecoder(r.Body).Decode(&updatedNote); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	note, success := store.UpdateNote(id, updatedNote)
	if !success {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// DELETE /api/notes/{id} - Delete a note by ID
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if success := store.DeleteNote(id); !success {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
