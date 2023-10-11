package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", generateHandler)

	http.HandleFunc("/web.css", cssHandler)

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Serve the HTML file
		tmpl, err := template.ParseFiles("web.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web.css")
}

func generateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		text := r.Form.Get("text")
		text = strings.TrimSpace(text)
		bannerType := r.Form.Get("bannerType")

		asciiArt := GenerateASCIIArt(text, bannerType)

		// Serve the ASCII art as the response
		fmt.Fprint(w, asciiArt)
	}
}

func CharConvert(arr []string, char rune, line int) string {
	ascii := (int(char) - 32) * 9 + line
	if ascii >= 0 && ascii < len(arr) {
		return arr[ascii]
	}
	return ""
}

func GenerateASCIIArt(text string, bannerType string) string {
	var bannerContent []string
	var bannerFile string

	switch bannerType {
	case "standard":
		bannerFile = "standard.txt"
	case "thinkertoy":
		bannerFile = "thinkertoy.txt"
	case "shadow":
		bannerFile = "shadow.txt"
	default:
		bannerFile = "standard.txt"
	}

	bannerContent, err := readBannerFile(bannerFile)
	if err != nil {
		return fmt.Sprintf("Error reading %s file: %v", bannerFile, err)
	}

	splitted := strings.Split(text, "\n")
	var output string

	for _, line := range splitted {
		if line == "" {
			output += "\n"
			continue
		}

		for j := 1; j < 9; j++ {
			for _, char := range line {
				output += CharConvert(bannerContent, rune(char), j)
			}
			output += "\n"
		}
	}

	return output
}

func readBannerFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}