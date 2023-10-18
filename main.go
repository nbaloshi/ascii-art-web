package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/generate", generateHandler)

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
