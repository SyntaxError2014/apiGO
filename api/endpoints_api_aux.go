package api

import (
    "apiGO/config"
    "apiGO/dbmodels"
    "apiGO/filter"
    "apiGO/models"
    "apiGO/random"
    "apiGO/service"
    "strings"
    "time"
)

// Create a random endpoint with unique path
func generateNewEndpoint(user *dbmodels.User) *dbmodels.Endpoint {
    endpoint := &dbmodels.Endpoint{
        UserId:      user.Id,
        URLPath:     strings.Join([]string{"/", random.RandomString(8)}, ""),
        Name:        "Untitled",
        Description: "-",
        Enabled:     true,
        DateCreated: time.Now().Local(),
        REST:        make(map[string]dbmodels.EndpointResponse, 4),
    }

    endpoint.REST[GET] = dbmodels.NewEndpointResponse(GET)
    endpoint.REST[POST] = dbmodels.NewEndpointResponse(POST)
    endpoint.REST[PUT] = dbmodels.NewEndpointResponse(PUT)
    endpoint.REST[DELETE] = dbmodels.NewEndpointResponse(DELETE)

    return endpoint
}

// Create a new route with random Id, for a certaing path,
// based on its associated endpoint
func generateNewRoute(endpoint *models.Endpoint, resp *ApiResponse) error {
    route := &config.Route{
        Id:       random.RandomString(5),
        Pattern:  endpoint.URLPath,
        Handlers: make(map[string]string, len(endpoint.REST)),
    }

    for method, endpointResponse := range endpoint.REST {
        route.Handlers[method] = endpointResponse.GetApiFunction(method)
    }

    return config.AddRoute(route, true)
}

// Update a certain Route using the details from an Endpoint
func updateRoute(routePath string, endpoint *dbmodels.Endpoint) error {
    route := config.GetRouteByPattern(routePath)
    route.Pattern = endpoint.URLPath
    route.Handlers = make(map[string]string, len(endpoint.REST))

    for method, endpointResponse := range endpoint.REST {
        if endpointResponse.StatusCode != 0 {
            route.Handlers[method] = endpointResponse.GetApiFunction(method)
        }
    }

    return config.ModifyRoute(route.Id, *route, true)
}

// Get the User entity for the session that has the
// entered token
func extractUserUsingToken(token string, resp *ApiResponse) *dbmodels.User {
    session, err := service.GetUserSessionByToken(token)
    if err != nil || session == nil {
        notFound(resp, "No used is logged in with these credentials")
        return nil
    }

    user, err := service.GetUser(session.UserId)
    if err != nil || user == nil {
        notFound(resp, "The user that was logged in no longer exists")
        return nil
    }

    return user
}

func authenticateUsingToken(vars *ApiVar, resp *ApiResponse) *dbmodels.User {
    token, err, found := filter.GetStringValueFromParams("token", vars.RequestForm)

    if !found {
        badRequest(resp, "The authentication token was not specified")
        return nil
    }

    if err != nil || token == "" {
        badRequest(resp, err.Error())
        return nil
    }

    return extractUserUsingToken(token, resp)
}
