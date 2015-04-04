package main

import (
    "apiGO/config"
    "apiGO/servers"
    "runtime"
)

const basicObjectId = "507f1f77bcf86cd799439011"

// Function for performing automatic initializations at application startup
func initApplicationConfiguration() {
    var emptyConfigParam string = ""

    config.InitApp(emptyConfigParam)
    config.InitDatabase(emptyConfigParam)
    config.InitRoutes(emptyConfigParam)
}

// Application entry point - sets the behaviour for the app
func main() {
    initApplicationConfiguration()

    runtime.GOMAXPROCS(2) // in order for the rpc and http servers to work in parallel

    go servers.StartRPCServer()
    servers.StartHTTPServer()
}
