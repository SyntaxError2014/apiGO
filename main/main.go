package main

import (
    "apiGO/config"
    "apiGO/dbmodels"
    "apiGO/servers"
    "apiGO/service"
    "gopkg.in/mgo.v2/bson"
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

func initBasicEndpoint() {
    id := bson.ObjectIdHex(basicObjectId)

    endpoint, _ := service.GetEndpoint(id)

    if endpoint == nil {
        endpoint = &dbmodels.Endpoint{
            Id:      id,
            URLPath: bson.NewObjectId().Hex(),
            GET:     dbmodels.NewEndpointResponse("GET"),
            POST:    dbmodels.NewEndpointResponse("POST"),
            PUT:     dbmodels.NewEndpointResponse("PUT"),
            DELETE:  dbmodels.NewEndpointResponse("DELETE"),
        }

        service.CreateEndpoint(endpoint)
    }
}

// Application entry point - sets the behaviour for the app
func main() {
    initApplicationConfiguration()
    initBasicEndpoint()

    runtime.GOMAXPROCS(2) // in order for the rpc and http servers to work in parallel

    go servers.StartRPCServer()
    servers.StartHTTPServer()
}
