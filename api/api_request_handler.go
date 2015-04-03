package api

import (
    "apiGO/config"
    "net/http"
    "net/url"
)

// All the requests to the server go through this handler.
//
// This parses requests, filters them and takes appropriate actions
// based on the content and configuration of each http request.
//
// The actions to be taken are stored in the Routes configuration
// of the application
func ApiHandler(rw http.ResponseWriter, req *http.Request) {
    path, err := parseRequestURI(req.RequestURI)

    if err != nil {
        GiveApiMessage(http.StatusBadRequest, "The format of the request URL is invalid", rw)
        return
    }

    route := findRoute(path)

    if route == nil {
        GiveApiMessage(http.StatusNotFound, "The requested URL cannot be found", rw)
        return
    }

    handler := findApiMethod(req.Method, route)

    if handler == "" {
        GiveApiMessage(http.StatusBadRequest, "The requested method is either not implemented, or not allowed", rw)
        return
    }

    PerformClientCall(handler, rw, req, route)
}

// Finds the route which has the selected pattern.
// Returns nil in case such a route doesn't exist
func findRoute(pattern string) *config.Route {
    for _, route := range config.Routes {
        if route.Pattern == pattern {
            return &route
        }
    }

    return nil
}

// Searches a certain routes to see whether it accepts a certain
// REST call method or not.
// If the method is allowed at this route, then the function returns
// the name of the endpoint functionality that needs to be used
func findApiMethod(requestMethod string, route *config.Route) string {
    if handler, found := route.Handlers[requestMethod]; found {
        return handler
    }

    return ""
}

// Parse the request URL and returns the trailing path of the request
// Ex: http://something.com/users?id=some_id returns: /users
func parseRequestURI(uri string) (string, error) {
    u, err := url.ParseRequestURI(uri)

    if err != nil {
        return "", err
    }

    return u.Path, nil
}