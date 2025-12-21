package main

import (
	"fmt"
	"net/http"
)

func checkHealthHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "UP",
	}
	writeJSON(w, http.StatusOK, data)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"message": "Hello, World!",
	}
	writeJSON(w, http.StatusOK, data)
}

func userListHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
	}
	writeJSON(w, http.StatusOK, users)
}

func postListHandler(w http.ResponseWriter, r *http.Request) {
	posts := []Post{
		{ID: 1, Title: "First Post", Content: "Hello world", Author: "Alice"},
		{ID: 2, Title: "Go is great", Content: "Concurrency is fun", Author: "Bob"},
	}
	writeJSON(w, http.StatusOK, posts)
}

func userDetailHandler(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	// Mock finding a user
	user := User{
		ID:    id,
		Name:  fmt.Sprintf("User %d", id),
		Email: fmt.Sprintf("user%d@example.com", id),
	}
	writeJSON(w, http.StatusOK, user)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := readJSON(r, &user); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	// Simple validation
	if user.Name == "" || user.Email == "" {
		writeError(w, http.StatusBadRequest, "Name and Email are required")
		return
	}

	// Mock creation - assign an ID
	user.ID = 101 // Simulated DB ID

	writeJSON(w, http.StatusCreated, user)
}

func postDetailHandler(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	// Mock finding a post
	post := Post{
		ID:      id,
		Title:   fmt.Sprintf("Post %d", id),
		Content: "Some content here...",
		Author:  "Unknown",
	}
	writeJSON(w, http.StatusOK, post)
}
