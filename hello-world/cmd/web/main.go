package main

// packages always live in their own directory
import (
	"fmt"
	"net/http"

	"github.com/aayushxadhikari/go-course/pkg/handlers"
)

const portNumber = ":8080"


// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
	
}