package main

type JsonResponse struct {
	// Reserved field to add some meta information to the API response
	Meta interface{} `json:"meta"`
	Data interface{} `json:"data"`
}

type JsonErrorResponse struct {
	Error *apiError `json:"error"`
}

type apiError struct {
	status int16
	title  string
}
