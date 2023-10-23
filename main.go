package main

import (
	"bufio"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Art struct {
	Result string
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", generateHandler)
	http.HandleFunc("/web.css", cssHandler)
	http.HandleFunc("/404", custom404Handler)

	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if r.URL.Path != "/" {
			// http.Redirect(w, r, "/404", http.StatusNotFound)
			http.ServeFile(w, r, "error404.html")
			return
		}

		// Serve the HTML file
		tmpl, err := template.ParseFiles("web.html")
		if err != nil {
			http.Redirect(w, r, "/404", http.StatusNotFound)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Redirect(w, r, "/404", http.StatusNotFound)
			return
		}
	}
}

func custom404Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "error404.html")
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

		if text == "" || text <= string(32) || text >= string(128) {
			http.ServeFile(w, r, "error400.html")
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
	} else {
		http.Error(w, "405", http.StatusMethodNotAllowed)
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

	bannerFile = bannerType + ".txt"
	_, err := ioutil.ReadFile(bannerFile)
	if err != nil {
		return "Internal Server Error: Error 500"
	}

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

	bannerContent, err = readBannerFile(bannerFile)
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
