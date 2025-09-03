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

type TechItem struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Icon     string `json:"icon"`
	Color    string `json:"color"`
}

func main() {
	// Serve static files
	http.Handle("/", http.FileServer(http.Dir("./static/")))

	// API endpoints with CORS
	http.HandleFunc("/api/about", corsMiddleware(aboutHandler))
	http.HandleFunc("/api/projects", corsMiddleware(projectsHandler))
	http.HandleFunc("/api/tech", corsMiddleware(techHandler))

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

// GET /api/tech
func techHandler(w http.ResponseWriter, r *http.Request) {
	tech := []TechItem{
		// Languages
		{Name: "Go", Category: "languages", Icon: "/icons/go-original.svg", Color: "#00ADD8"},
		{Name: "JavaScript", Category: "languages", Icon: "/icons/javascript-original.svg", Color: "#F7DF1E"},
		{Name: "TypeScript", Category: "languages", Icon: "/icons/typescript-original.svg", Color: "#3178C6"},
		{Name: "Python", Category: "languages", Icon: "/icons/python-original.svg", Color: "#3776AB"},
		{Name: "Java", Category: "languages", Icon: "/icons/java-original.svg", Color: "#ED8B00"},
		{Name: "R", Category: "languages", Icon: "/icons/r-original.svg", Color: "#DEDEDE"},
		{Name: "C", Category: "languages", Icon: "/icons/c-original.svg", Color: "#0077ffff"},
		{Name: "Kotlin", Category: "languages", Icon: "/icons/kotlin-original.svg", Color: "#9000ffff"},
		{Name: "HTML5", Category: "languages", Icon: "/icons/html5-original.svg", Color: "#ff0000"},
		{Name: "Ocaml", Category: "languages", Icon: "/icons/ocaml-original.svg", Color: "#ffff00"},
		{Name: "C#", Category: "languages", Icon: "/icons/csharp-original.svg", Color: "#0000ff"},
		{Name: "Elixir", Category: "languages", Icon: "/icons/elixir-original.svg", Color: "#a929f4ff"},

		// Frameworks & Libraries
		{Name: "React", Category: "frameworks", Icon: "/icons/react-original.svg", Color: "#61DAFB"},
		{Name: "Next.js", Category: "frameworks", Icon: "/icons/nextjs-original.svg", Color: "#000000"},
		{Name: "Django", Category: "frameworks", Icon: "/icons/django-plain.svg", Color: "#000000"},
		{Name: "Node.js", Category: "frameworks", Icon: "/icons/nodejs-original.svg", Color: "#339933"},
		{Name: "Tailwind css", Category: "frameworks", Icon: "/icons/tailwindcss-original.svg", Color: "#2afcfcff"},
		{Name: "Pandas", Category: "frameworks", Icon: "/icons/pandas-original.svg", Color: "#000000"},

		// DevOps &
		{Name: "Linux", Category: "devops", Icon: "/icons/linux-original.svg", Color: "#ffffff"},
		{Name: "Docker", Category: "devops", Icon: "/icons/docker-original.svg", Color: "#2496ED"},
		{Name: "Kubernetes", Category: "devops", Icon: "/icons/kubernetes-plain.svg", Color: "#326CE5"},
		{Name: "Terraform", Category: "devops", Icon: "/icons/terraform-original.svg", Color: "#7B42BC"},
		{Name: "Jenkins", Category: "devops", Icon: "/icons/jenkins-original.svg", Color: "#D33833"},
		{Name: "AWS", Category: "devops", Icon: "/icons/amazonwebservices-original-wordmark.svg", Color: "#FF9900"},
		{Name: "Selenium", Category: "devops", Icon: "/icons/selenium-original.svg", Color: "#000000"},

		// Databases
		{Name: "PostgreSQL", Category: "database", Icon: "/icons/postgresql-original.svg", Color: "#336791"},
		{Name: "MongoDB", Category: "database", Icon: "/icons/mongodb-original.svg", Color: "#47A248"},
		{Name: "mySQL", Category: "database", Icon: "/icons/mysql-original.svg", Color: "#2d4adcff"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tech)
}
