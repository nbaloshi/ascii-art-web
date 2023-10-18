package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Art struct {
	Result string
}

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

		if text == "" {
			http.Error(w, "Bad Request: (error 400)", http.StatusBadRequest)
			return
		}

		asciiArt := GenerateASCIIArt(text, bannerType)
		// Create an instance of Art struct to hold the ASCII art
		art := Art{Result: asciiArt}

		// tmpl, err := template.New("result").Parse(`{{.Result}}`)
		tmpl, err := template.ParseFiles("web.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, &art)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func CharConvert(arr []string, char rune, line int) string {
	ascii := (int(char)-32)*9 + line
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
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(fileContent), "\n")
	return lines, nil
}
