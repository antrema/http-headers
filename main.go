package main

import (
	"fmt"
	"log"
	"net/http"
)

func getHeaders(w http.ResponseWriter, r *http.Request) {
	fmt.Println("got / request")
	for k, v := range r.Header {
		fmt.Fprintf(w, "%v: %v\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr= %q\n", r.RemoteAddr)

	// Write a response
	w.Header().Set("Content-Type", "text/plain")
}

func main() {
	http.HandleFunc("/", getHeaders)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
