package servers

import (
    "apiGO/api"
    "apiGO/config"
    "log"
    "net/http"
    "time"
)

// Starts a HTTP server that listens for REST requests
func StartHTTPServer() {
    http.HandleFunc("/", api.ApiHandler)
    server := &http.Server{
        Addr:           config.HttpServerAddress,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }

    log.Println("HTTP Server STARTED! Listening at:", config.HttpServerAddress)
    log.Fatal(server.ListenAndServe())
}
