package main

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)
 
func rateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler{
	limiter := rate.NewLimiter(rate.Every(500*time.Millisecond), 4)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow(){
			message := Message{
				Status : "Request Failed",
				Body : "THe API is at capacity, Try Again.",
			}

			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&message)
			return
		}else{
			next(w,r)
		}
	})
}


