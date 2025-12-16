package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func extractID(r *http.Request, paramName string) (int, error) {
	idStr := r.PathValue(paramName)
	if idStr == "" {
		return 0, fmt.Errorf("%s is required", paramName)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid %s", paramName)
	}
	return id, nil
}
