package main

import (
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
	writeJSON(w, http.StatusOK, users)
}

func postListHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, posts)
}

func userDetailHandler(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	for _, u := range users {
		if u.ID == id {
			writeJSON(w, http.StatusOK, u)
			return
		}
	}

	writeError(w, http.StatusNotFound, "User not found")
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

	// Find max ID to auto-increment
	maxID := 0
	for _, u := range users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	user.ID = maxID + 1

	// Store in memory
	users = append(users, user)

	writeJSON(w, http.StatusCreated, user)
}

func postDetailHandler(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	for _, p := range posts {
		if p.ID == id {
			writeJSON(w, http.StatusOK, p)
			return
		}
	}

	writeError(w, http.StatusNotFound, "Post not found")
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	for i, u := range users {
		if u.ID == id {
			var updatedUser User
			if err := readJSON(r, &updatedUser); err != nil {
				writeError(w, http.StatusBadRequest, "Invalid JSON body")
				return
			}

			// Update the actual slice element using index
			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email

			writeJSON(w, http.StatusOK, users[i])
			return
		}
	}

	writeError(w, http.StatusNotFound, "User not found")
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			writeJSON(w, http.StatusOK, map[string]string{"message": "User deleted"})
			return
		}
	}
	writeError(w, http.StatusNotFound, "User not found")
}
