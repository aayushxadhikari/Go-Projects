package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

// in-memory storage
var (
	items  = []Item{}
	nextID = 1
	mu     sync.Mutex
)

// CRUD Functions used by both CLI and server
func createItem(name string, price int) Item {
	mu.Lock()
	defer mu.Unlock()
	item := Item{ID: nextID, Name: name, Price: price}
	nextID++
	items = append(items,item)
	return item
}

func listItems() []Item{
	mu.Lock()
	defer mu.Unlock()
	return append([]Item(nil), items...)	
}

func getItem(id int) (Item,bool){
	mu.Lock()
	defer mu.Unlock()
	for _, item := range items{
		if item.ID == id {
			return item, true
		}
	}
	return Item{}, false
}

func updateItem(id int, name string, price int) (Item, bool){
	mu.Lock()
	defer mu.Unlock()
	for i, item := range items{
		if item.ID == id{
			if name != ""{
				items[i].Name= name
			}
			if price!= 0{
				items[i].Price = price
			}
			return items[i], true
		}
	}
	return Item{}, false
}

func deleteItem(id int) bool{
	mu.Lock()
	defer mu.Unlock()
	for i, item := range items{
		if item.ID == id{
			items = append(items[:i], items[i+1:]...)
			return true
		}
	}
	return false
}

// Helper for pretty JSON printing in CLI
func printJSON(v any){
	b,_:=json.MarshalIndent(v,"","")
	fmt.Println(string(b))
}