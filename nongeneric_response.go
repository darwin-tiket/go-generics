package main

// START OMIT
type FindOrderAPIResponse struct {
	Data    FindOrderData `json:"data"`
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Errors  []string      `json:"errors"`
}

type CreateOrderAPIResponse struct {
	Data    CreateOrderData `json:"data"`
	Code    string          `json:"code"`
	Message string          `json:"message"`
	Errors  []string        `json:"errors"`
}

// END OMIT
