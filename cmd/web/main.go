package main

import (
	"net/http"
	"github.com/TendonT52/SimpleWebApplication/pkg/handlers"
)

const portNumber string = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home) 
	http.HandleFunc("/about", handlers.About) 
	http.ListenAndServe(portNumber, nil)	
}