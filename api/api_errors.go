package api

import (
    "net/http"
)

// Return a status and message that signals the API client
// about an 'internal server error' that has occured
func internalServerError(resp *ApiResponse, msg string) error {
    if msg == "" {
        msg = http.StatusText(http.StatusInternalServerError)
    }

    resp.StatusCode = http.StatusInternalServerError
    resp.Message = []byte(msg)
    resp.ErrorMessage = msg

    return nil
}

// Return a status and message that signals the API client
// about a 'bad request' that the client has made to the server
func badRequest(resp *ApiResponse, msg string) error {
    resp.StatusCode = http.StatusBadRequest
    resp.Message = []byte(msg)
    resp.ErrorMessage = msg

    return nil
}

// Return a status and message that signals the API client
// that the searched resource was not found on the server
func notFound(resp *ApiResponse, msg string) error {
    resp.StatusCode = http.StatusNotFound
    resp.Message = []byte(msg)
    resp.ErrorMessage = http.StatusText(http.StatusNotFound)

    return nil
}
