package domain

import "errors"

type ErrorResponse struct {
	Message string `json:"message"`
}

// Define standard domain errors
var (
	ErrNotFound       = errors.New("requested item not found")
	ErrDuplicateEntry = errors.New("an entry with the given details already exists")
)
