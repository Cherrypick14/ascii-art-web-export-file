package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"artweb/handlers"
)


func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
		return
	}
	// http.Handle("/templates/stats/", http.StripPrefix("/templates/stats/", http.FileServer(http.Dir("templates/stats"))))
	// Handle all requests to the root path
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)
	http.HandleFunc("/400", handlers.BadRequestHandler)
	http.HandleFunc("/404", handlers.NotFoundHandler)
	http.HandleFunc("/405", handlers.FourohFiveHandler)
	http.HandleFunc("/500", handlers.InternalServerHandler)
	port := ":8000"
	// Start the server on port 8080
	log.Printf("Server started on http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
