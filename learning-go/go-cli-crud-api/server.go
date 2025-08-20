package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func startServer() {
	r := chi.NewRouter()

	r.Post("/items/", createItemHandler)
	r.Get("/items/", getAllItemsHandler)
	r.Get("/items/{id}/",getItemHandler)
	r.Put("items/{id}/",updateItemHandler)
	r.Delete("items/{id}/",deleteItemHandler)

	fmt.Println("Server starting in http://localhost:8080")
	http.ListenAndServe(":8080",r)
}

func createItemHandler(w http.ResponseWriter,r *http.Request){
	var newItem Item
	if err:= json.NewDecoder(r.Body).Decode(&newItem); err != nil{
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	item:= createItem(newItem.Name,newItem.Price)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(item)
}

func getAllItemsHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(listItems())
}

func getItemHandler(w http.ResponseWriter, r *http.Request){
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w,"Invalid ID", http.StatusBadRequest)
		return
	}
	if item, ok:=getItem(id); ok{
		w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(item)
	}else{
		http.NotFound(w,r)
	}
}

func updateItemHandler(w http.ResponseWriter, r *http.Request){
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updated Item
	if err:= json.NewDecoder(r.Body).Decode(&updated); err!=nil{
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if item, ok := updateItem(id, updated.Name,updated.price);ok{
		w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(item)
	}else{
		http.NotFound(w,r)
	}
}

func deleteItemHandler(w http.ResponseWriter,r *http.Request){
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if ok:= deleteItem(id); ok{
		w.WriteHeader(http.StatusNoContent)
		return
	}else{
		http.NotFound(w,r)
	}
	
}
