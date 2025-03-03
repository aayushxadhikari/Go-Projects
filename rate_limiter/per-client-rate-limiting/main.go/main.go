package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"golang.org/x/time/rate"
)

// Message struct for JSON response
type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

// Rate limiter per client
var clients = make(map[string]*rate.Limiter)
var mu sync.Mutex

func getClientLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := clients[ip]
	if !exists {
		limiter = rate.NewLimiter(1, 3) // Allow 1 request per second with a burst of 3
		clients[ip] = limiter
	}

	return limiter
}

func perClientRateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		limiter := getClientLimiter(ip)
		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func endpointHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	message := Message{
		Status: "Successful",
		Body:   "Hi! You've reached the API. How may I help you?",
	}

	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Println("Error encoding JSON:", err)
	}
}

func main() {
	http.Handle("/ping", perClientRateLimiter(http.HandlerFunc(endpointHandler)))

	fmt.Println("Server is running on port 8080....")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error listening on port :8080", err)
	}
}
