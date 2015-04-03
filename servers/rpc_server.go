package servers

import (
    "apiGO/api"
    "apiGO/config"
    "log"
    "net"
    "net/http"
    "net/rpc"
)

// Starts an RPC (remote process call) server that is used internally
// by the HTTP server.
// This is used for easier mapping between routes and functions that
// are used by the roots in order to respond to the requests.
// Using the RPC server this way results in great horizontal scalability
func StartRPCServer() {
    rpc.Register(new(api.Api))
    rpc.HandleHTTP()

    listener, err := net.Listen(config.Protocol, config.RpcServerAddress)
    if err != nil {
        log.Fatal("Error starting RPC server (listen error): ", err)
    }

    log.Println("RPC Server STARTED! Listening at:", config.RpcServerAddress)

    err = http.Serve(listener, nil)
    if err != nil {
        log.Fatal("RPC server error: ", err)
    }
}
