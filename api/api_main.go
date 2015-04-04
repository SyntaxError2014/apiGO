// Package containing the actual API functionality
//
// Each file represents the API functionality for a certain
// database dbmodels (GET, POST, PUT, DELETE)
package api

import (
    "apiGO/config"
    "apiGO/filter"
    "net/http"
    "net/url"
)

const (
    GET    = "GET"
    POST   = "POST"
    PUT    = "PUT"
    DELETE = "DELETE"
)

// Used for defining a type representing the api
type Api int

// Data type containing the basic auth details
type BasicAuthentication struct {
    Username string
    Password string
    OK       bool
}

// Data type containing important data from a HTTP
// request that has been made to the server
type ApiVar struct {
    Route                config.Route
    RequestMethod        string
    RequestHeader        http.Header
    RequestForm          url.Values
    RequestContentLength int64
    RequestBody          []byte
    BasicAuth            BasicAuthentication
}

// Data type containing important data that will be
// sent as response to HTTP requests
type ApiResponse struct {
    Message      []byte
    StatusCode   int
    ErrorMessage string
}

// Responds to the client that made the HTTP request with a status code
// and appropriate response data (usually JSON/XML encoded data)
func GiveApiResponse(statusCode int, message []byte, rw http.ResponseWriter) {
    rw.Header().Add("Content-Type", "application/json")
    rw.Header().Set("WWW-Authenticate", "Basic realm=\"private\"")

    if filter.CheckNotNull(message) {
        rw.WriteHeader(statusCode)
        rw.Write(message)
    } else {
        rw.WriteHeader(http.StatusNoContent)
    }
}

// Responds to the client that made the HTTP request with a
// status code and a simple message
func GiveApiMessage(statusCode int, message string, rw http.ResponseWriter) {
    msg := []byte(message)

    GiveApiResponse(statusCode, msg, rw)
}

// Responds to the client that made the HTTP request with a status code
// and a default message for that status code
func GiveApiStatus(statusCode int, rw http.ResponseWriter) string {
    code := http.StatusNoContent
    msg := http.StatusText(code)

    GiveApiMessage(code, msg, rw)

    return msg
}
