package main

// START OMIT
type APIResponse[T any] struct {
	Data    T        `json:"data"`
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

// END OMIT
