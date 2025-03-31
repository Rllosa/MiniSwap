package main

import (
	"fmt"
	"log"
	"net/http"

	api "github.com/Rllosa/miniSwap/backend/rest/API"
	"github.com/rs/cors"
)

func main() {
	// Initialize CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Replace with specific origins in production
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		Debug:          true,
	})

	// Create a new mux for better route handling
	mux := http.NewServeMux()

	mux.Handle("/swap", corsHandler.Handler(http.HandlerFunc(api.Swap)))
	mux.Handle("/getBalance", corsHandler.Handler(http.HandlerFunc(api.GetBalance)))

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "WebSocket endpoint moved to port 9000", http.StatusTemporaryRedirect)
	})

	// Wrap the mux with CORS middleware
	handler := corsHandler.Handler(mux)

	// Start HTTP server
	fmt.Printf("Starting HTTP server at port 8080\n")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
