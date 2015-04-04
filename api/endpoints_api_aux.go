package api

import (
    "apiGO/config"
    "apiGO/dbmodels"
    "apiGO/models"
    "apiGO/random"
    "strings"
)

func generateNewEndpoint() *dbmodels.Endpoint {
    endpoint := &dbmodels.Endpoint{
        URLPath:     strings.Join([]string{"/", random.RandomString(8)}, ""),
        Name:        "Untitled",
        Description: "-",
        Enabled:     true,
        REST:        make(map[string]dbmodels.EndpointResponse, 4),
    }

    endpoint.REST[GET] = dbmodels.NewEndpointResponse(GET)
    endpoint.REST[POST] = dbmodels.NewEndpointResponse(POST)
    endpoint.REST[PUT] = dbmodels.NewEndpointResponse(PUT)
    endpoint.REST[DELETE] = dbmodels.NewEndpointResponse(DELETE)

    return endpoint
}

func generateNewRoute(endpoint *models.Endpoint, resp *ApiResponse) error {
    route := &config.Route{
        Id:       random.RandomString(5),
        Pattern:  endpoint.URLPath,
        Handlers: make(map[string]string, len(endpoint.REST)),
    }

    for method, endpointResponse := range endpoint.REST {
        route.Handlers[method] = endpointResponse.Function
    }

    return config.AddRoute(route, true)
}

func updateRoute(routePath string, endpoint *dbmodels.Endpoint) error {
    route := config.GetRouteByPattern(routePath)
    route.Pattern = endpoint.URLPath
    route.Handlers = make(map[string]string, len(endpoint.REST))

    for method, endpointResponse := range endpoint.REST {
        route.Handlers[method] = endpointResponse.Function
    }

    return config.ModifyRoute(route.Id, *route, true)
}
