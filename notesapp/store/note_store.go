package store

import (
	"strconv"
	"sync"
	"time"

	"notesapp/models"
)

// Thread-safe map for storing notes
var (
	noteStore = make(map[string]models.Note)
	id        int
	mu        sync.RWMutex
)

// Create a new note
func AddNote(note models.Note) models.Note {
	mu.Lock()
	defer mu.Unlock()

	id++
	key := strconv.Itoa(id)
	note.CreatedOn = time.Now()
	noteStore[key] = note

	return note
}

// Get all notes
func GetAllNotes() []models.Note {
	mu.RLock()
	defer mu.RUnlock()

	var notes []models.Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}
	return notes
}

// Update a note by ID
func UpdateNote(id string, updatedNote models.Note) (models.Note, bool) {
	mu.Lock()
	defer mu.Unlock()

	if note, exists := noteStore[id]; exists {
		updatedNote.CreatedOn = note.CreatedOn
		noteStore[id] = updatedNote
		return updatedNote, true
	}
	return models.Note{}, false
}

// Delete a note by ID
func DeleteNote(id string) bool {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := noteStore[id]; exists {
		delete(noteStore, id)
		return true
	}
	return false
}
