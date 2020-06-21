package main

import (
	"fmt"
	"log"
	"net/http"
)

const message = "This is a long piece of text"

// Hello - just a little piece of text
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(message))
	fmt.Println(message)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hello)
	err := http.ListenAndServeTLS("127.0.0.1:443", "server.crt", "server.key", mux)
	if err != nil {
		log.Fatal("Failed to create the server", err)
	}
}
