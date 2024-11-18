package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// ParseStrURLParam parses a string URL parameter
func ParseStrURLParam(key string, r *http.Request) string {
	return chi.URLParam(r, key)
}

// ParseStrQuery parses a string query parameter
func ParseStrQuery(key string, r *http.Request) string {
	return r.URL.Query().Get(key)
}

// ParseIntQuery parses an integer query parameter
func ParseIntQuery(key string, r *http.Request, defaultValue int, required bool) (int, error) {
	value := ParseStrQuery(key, r)
	if value == "" {
		if required {
			return 0, fmt.Errorf("%s is required", key)
		}
		return defaultValue, nil
	}

	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be an integer", key)
	}
	return v, nil
}

// ParsePaginationParams parses the offset and limit query parameters
func ParsePaginationParams(r *http.Request) (int, int) {
	offset, _ := ParseIntQuery("offset", r, 0, false)
	limit, _ := ParseIntQuery("limit", r, 10, false)
	return offset, limit
}
