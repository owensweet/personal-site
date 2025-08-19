package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {
	// Serve static files
	http.Handle("/", http.FileServer(http.Dir("./static/")))
	
	// API endpoints with CORS
	http.HandleFunc("/api/about", corsMiddleware(aboutHandler))
	http.HandleFunc("/api/projects", corsMiddleware(projectsHandler))

	port := ":8080"
	fmt.Println("Server listening on", port)
	fmt.Println("Put your index.html in ./static/ folder")
	log.Fatal(http.ListenAndServe(port, nil))
}

// Middleware setup
func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		if r.Method == "OPTIONS" {
			return
		}
		
		next(w, r)
	}
}

// GET /
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my site")
}

// GET /api/about
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page")
}

// GET /api/projects
func projectsHandler(w http.ResponseWriter, r *http.Request) {
	projects := []Project{
		{Name: "BCIT Accreditation System", Description: "Deployed the backend database and interface system for the engineering departments of BCIT"},
		{Name: "Atmoxhere", Description: "Interactive clothing website built on next.js"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
