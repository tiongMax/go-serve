package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", WithDefaults(rootHandler))
	mux.HandleFunc("GET /health", WithDefaults(checkHealthHandler))

	mux.HandleFunc("GET /users", WithDefaults(userListHandler))
	mux.HandleFunc("POST /users", WithDefaults(createUserHandler))
	mux.HandleFunc("GET /users/{id}", WithDefaults(userDetailHandler))
	mux.HandleFunc("PUT /users/{id}", WithDefaults(updateUserHandler))
	mux.HandleFunc("DELETE /users/{id}", WithDefaults(deleteUserHandler))

	mux.HandleFunc("GET /posts", WithDefaults(postListHandler))
	mux.HandleFunc("GET /posts/{id}", WithDefaults(postDetailHandler))

	log.Println("🚀 Server starting on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8083", mux))
}
