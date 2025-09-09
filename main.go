package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!, This Go application runs on Clouddley")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3007", nil)
}
