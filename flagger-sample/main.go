package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "VERSION:"+os.Getenv("VERSION"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
