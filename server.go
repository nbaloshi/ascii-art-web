package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", localserver)
	http.ListenAndServe(":8080", nil)
}

func localserver(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "web.html")
	} else if r.URL.Path == "/web.css" {
		http.ServeFile(w, r, "web.css")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error 404: Page not found")
	}
}