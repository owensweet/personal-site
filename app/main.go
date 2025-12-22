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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "./static/index.html")
		} else {
			http.NotFound(w, r)
		}
	})

	// API endpoints with CORS
	http.HandleFunc("/api/about", corsMiddleware(aboutHandler))
	http.HandleFunc("/api/projects", corsMiddleware(projectsHandler))
	http.HandleFunc("/api/tech", corsMiddleware(techHandler))

	port := ":8080"
	fmt.Println("Server listening on", port)
	fmt.Println("Put your index.html in ./static/ folder")
	log.Fatal(http.ListenAndServe(port, nil))
}

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
		{Name: "BCIT Accreditation System", Description: "Deployed the backend database and interface system for the engineering departments of BCIT. Designed the relational database and REST APIs. Used python's Django framework for the backend, and used Docker and Nginx to host on BCIT servers. It is currently being used by almost 100 faculty every semester."},
		{Name: "Personal Site", Description: "Simple dev site, with the static files served from server programmed in Go. Hosted on an AWS EC2 instance and runs docker containers (1 for reverse proxy). CI/CD pipelines using github actions and IaC with Terraform."},
		{Name: "Meetup", Description: "A navigation app built on GoogleMaps API, to calculate a meeting point between two people and taking into account different modes of transportation"},
		{Name: "Golden Gaming", Description: "A gaming social platform built for use in senior homes. Under agile development, created the fullstack application along with 4 other team members. Used MongoDB for storage, Websockets for chat system, and REST APIs with Express.js."},
		{Name: "Ada Analytics", Description: "Helped a startup create a new stock scoring feature. Used python libraries to web scrape stock ticker data, and then feed it to a genetic algorithm."},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

// GET /api/tech
func techHandler(w http.ResponseWriter, r *http.Request) {
	tech := []TechItem{
		// Languages
		{Name: "Go", Category: "languages", Icon: "/static/icons/go-original.svg", Color: "#00ADD8"},
		{Name: "JavaScript", Category: "languages", Icon: "/static/icons/javascript-original.svg", Color: "#F7DF1E"},
		{Name: "TypeScript", Category: "languages", Icon: "/static/icons/typescript-original.svg", Color: "#3178C6"},
		{Name: "Python", Category: "languages", Icon: "/static/icons/python-original.svg", Color: "#3776AB"},
		{Name: "Java", Category: "languages", Icon: "/static/icons/java-original.svg", Color: "#ED8B00"},
		{Name: "R", Category: "languages", Icon: "/static/icons/r-original.svg", Color: "#DEDEDE"},
		{Name: "C", Category: "languages", Icon: "/static/icons/c-original.svg", Color: "#0077ffff"},
		{Name: "Kotlin", Category: "languages", Icon: "/static/icons/kotlin-original.svg", Color: "#9000ffff"},
		{Name: "HTML5", Category: "languages", Icon: "/static/icons/html5-original.svg", Color: "#ff0000"},
		{Name: "Ocaml", Category: "languages", Icon: "/static/icons/ocaml-original.svg", Color: "#ffff00"},
		{Name: "C#", Category: "languages", Icon: "/static/icons/csharp-original.svg", Color: "#0000ff"},
		{Name: "Elixir", Category: "languages", Icon: "/static/icons/elixir-original.svg", Color: "#a929f4ff"},
		{Name: "Rust", Category: "languages", Icon: "/static/icons/rust-original.svg", Color: "#000000"},

		// Frameworks & Libraries
		{Name: "React", Category: "frameworks", Icon: "/static/icons/react-original.svg", Color: "#61DAFB"},
		{Name: "Next.js", Category: "frameworks", Icon: "/static/icons/nextjs-original.svg", Color: "#000000"},
		{Name: "Django", Category: "frameworks", Icon: "/static/icons/django-plain.svg", Color: "#000000"},
		{Name: "Node.js", Category: "frameworks", Icon: "/static/icons/nodejs-original.svg", Color: "#339933"},
		{Name: "Tailwind css", Category: "frameworks", Icon: "/static/icons/tailwindcss-original.svg", Color: "#2afcfcff"},
		{Name: "Pandas", Category: "frameworks", Icon: "/static/icons/pandas-original.svg", Color: "#000000"},

		// DevOps &
		{Name: "Linux", Category: "devops", Icon: "/static/icons/linux-original.svg", Color: "#ffffff"},
		{Name: "Docker", Category: "devops", Icon: "/static/icons/docker-original.svg", Color: "#2496ED"},
		{Name: "Kubernetes", Category: "devops", Icon: "/static/icons/kubernetes-plain.svg", Color: "#326CE5"},
		{Name: "Terraform", Category: "devops", Icon: "/static/icons/terraform-original.svg", Color: "#7B42BC"},
		{Name: "Jenkins", Category: "devops", Icon: "/static/icons/jenkins-original.svg", Color: "#D33833"},
		{Name: "Github Actions", Category: "devops", Icon: "/static/icons/githubactions-original.svg", Color: "#326Ce5"},
		{Name: "AWS", Category: "devops", Icon: "/static/icons/amazonwebservices-original-wordmark.svg", Color: "#FF9900"},
		{Name: "Selenium", Category: "devops", Icon: "/static/icons/selenium-original.svg", Color: "#000000"},

		// Databases
		{Name: "PostgreSQL", Category: "database", Icon: "/static/icons/postgresql-original.svg", Color: "#336791"},
		{Name: "MongoDB", Category: "database", Icon: "/static/icons/mongodb-original.svg", Color: "#47A248"},
		{Name: "mySQL", Category: "database", Icon: "/static/icons/mysql-original.svg", Color: "#2d4adcff"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tech)
}
