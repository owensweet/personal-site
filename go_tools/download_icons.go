package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	urls := []string{
		"https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon/icons/javascript/javascript-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon/icons/typescript/typescript-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon/icons/python/python-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon/icons/java/java-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon/icons/react/react-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon/icons/nextjs/nextjs-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon/icons/nodejs/nodejs-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon/icons/kubernetes/kubernetes-plain.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon/icons/terraform/terraform-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon/icons/jenkins/jenkins-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/amazonwebservices/amazonwebservices-original-wordmark.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon/icons/postgresql/postgresql-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon/icons/mongodb/mongodb-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon/icons/redis/redis-original.svg",

		"https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/tailwindcss/tailwindcss-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/mysql/mysql-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/linux/linux-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/r/r-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/c/c-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/kotlin/kotlin-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/django/django-plain.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/html5/html5-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/selenium/selenium-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/ocaml/ocaml-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/csharp/csharp-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/pandas/pandas-original.svg",
		"https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/elixir/elixir-original.svg",
	}

	outputDir := "../static/icons"
	os.MkdirAll(outputDir, os.ModePerm)

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Failed to download:", url, err)
			continue
		}

		// Take filename from URL
		file := filepath.Join(outputDir, filepath.Base(url))
		out, err := os.Create(file)
		if err != nil {
			fmt.Println("Failed to create file:", file, err)
			resp.Body.Close()
			continue
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			fmt.Println("Failed to save:", file, err)
			continue
		}

		fmt.Println("Downloaded:", file)
		defer resp.Body.Close()
	}
}
