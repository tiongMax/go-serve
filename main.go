package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", WithDefaults(rootHandler))

	mux.HandleFunc("/health", WithDefaults(checkHealthHandler))

	mux.HandleFunc("/users", WithDefaults(userListHandler))

	mux.HandleFunc("/posts", WithDefaults(postListHandler))

	mux.HandleFunc("/users/{id}", WithDefaults(userDetailHandler))

	mux.HandleFunc("/posts/{id}", WithDefaults(postDetailHandler))

	log.Println("🚀 Server starting on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
