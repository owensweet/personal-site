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
		{Name: "Go", Category: "languages", Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg", Color: "#00ADD8"},
		{Name: "JavaScript", Category: "languages", Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/javascript/javascript-original.svg", Color: "#F7DF1E"},
		{Name: "TypeScript", Category: "languages", Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/typescript/typescript-original.svg", Color: "#3178C6"},
		{Name: "Python", Category: "languages", Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/python/python-original.svg", Color: "#3776AB"},
		{Name: "Java", Category: "languages", Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/java/java-original.svg", Color: "#ED8B00"},

		// Frameworks & Libraries
		{Name: "React", Category: "frameworks", Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/react/react-original.svg", Color: "#61DAFB"},
		{Name: "Next.js", Category: "frameworks", Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/nextjs/nextjs-original.svg", Color: "#000000"},
		{Name: "Node.js", Category: "frameworks", Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/nodejs/nodejs-original.svg", Color: "#339933"},

		// DevOps & Tools
		{Name: "Docker", Category: "devops", Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg", Color: "#2496ED"},
		{Name: "Kubernetes", Category: "devops", Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/kubernetes/kubernetes-plain.svg", Color: "#326CE5"},
		{Name: "Terraform", Category: "devops", Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/terraform/terraform-original.svg", Color: "#7B42BC"},
		{Name: "Jenkins", Category: "devops", Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/jenkins/jenkins-original.svg", Color: "#D33833"},
		{Name: "AWS", Category: "devops", Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/amazonwebservices/amazonwebservices-original.svg", Color: "#FF9900"},

		// Databases
		{Name: "PostgreSQL", Category: "databases", Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/postgresql/postgresql-original.svg", Color: "#336791"},
		{Name: "MongoDB", Category: "databases", Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/mongodb/mongodb-original.svg", Color: "#47A248"},
		{Name: "Redis", Category: "databases", Icon: "https://cdn.jsdelivr.net/gh/devicons/devicon/icons/redis/redis-original.svg", Color: "#DC382D"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tech)
}
