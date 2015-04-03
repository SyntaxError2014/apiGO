package servers

import (
    "apiGO/api"
    "apiGO/config"
    "log"
    "net"
    "net/http"
    "net/rpc"
)

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
