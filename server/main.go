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
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}
