package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the home page.")
}

func About(w http.ResponseWriter, r *http.Request){
	sum := addValues(3,4)
	fmt.Fprintf(w, fmt.Sprintf("This is the about page and sum of the 3 + 4 is %d", sum))
}

func addValues(x,y int) int{
	return  x+y
}

func Divide(w http.ResponseWriter, r *http.Request){
	f, err := divideValues(100.0, -10.0)
	if err != nil{
		fmt.Fprintf(w, "Cannot Divide by 0")
		return 
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0,10.0,f))
}

func divideValues(x, y float32)(float32, error){
	if y <=0{
		err := errors.New("Cannot Divide by 0.")
		return 0, err
	}
	result := x/y
	return result, nil

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/About", About)
	http.HandleFunc("/Divide", Divide)
	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
	
}