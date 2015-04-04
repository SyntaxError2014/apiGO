package api

import (
    "apiGO/config"
    "apiGO/dbmodels"
    "apiGO/filter"
    "apiGO/models"
    "apiGO/random"
    "apiGO/service"
    "net/http"
    "strings"
)

func (api *Api) GetEndpoint(vars *ApiVar, resp *ApiResponse) error {
    resp.StatusCode = 204
    resp.Message = []byte("")
    return nil
}

func (api *Api) PostEndpoint(vars *ApiVar, resp *ApiResponse) error {
    basicEndpoint := &dbmodels.Endpoint{
        URLPath: strings.Join([]string{"/", random.RandomString(8)}, ""),
        GET:     dbmodels.NewEndpointResponse("GET"),
        POST:    dbmodels.NewEndpointResponse("POST"),
        PUT:     dbmodels.NewEndpointResponse("PUT"),
        DELETE:  dbmodels.NewEndpointResponse("DELETE"),
    }

    basicEndpoint, err := service.CreateEndpoint(basicEndpoint)

    if err != nil || basicEndpoint == nil {
        return internalServerError(resp, err.Error())
    }

    endpoint := &models.Endpoint{}
    endpoint.Expand(*basicEndpoint)

    endpointJson, err := endpoint.SerializeJson()

    if err != nil || endpointJson == nil {
        return internalServerError(resp, err.Error())
    }

    err = generateNewRoute(endpoint, resp)

    if err != nil {
        return internalServerError(resp, err.Error())
    }

    resp.StatusCode = http.StatusCreated
    resp.Message = endpointJson

    return nil
}

func (api *Api) PutEndpoint(vars *ApiVar, resp *ApiResponse) error {
    expandedEndpoint := &models.Endpoint{}

    err := expandedEndpoint.DeserializeJson(vars.RequestBody)
    if err != nil {
        return badRequest(resp, "The entity was not in the correct format")
    }

    if expandedEndpoint.Id == "" {
        return badRequest(resp, "No id was specified for the endpoint to be updated")
    }

    if !filter.CheckEndpointIntegrity(expandedEndpoint) {
        return badRequest(resp, "The entity doesn't comply to the integrity requirements")
    }

    endpoint, err := expandedEndpoint.Collapse()
    if err != nil {
        return internalServerError(resp, err.Error())
    }

    initialEndpoint, _ := service.GetEndpoint(endpoint.Id)
    initialRoutePath := initialEndpoint.URLPath

    err = service.UpdateEndpoint(endpoint)
    if err != nil {
        return notFound(resp, "The endpoint with the specified id could not be found")
    }

    route := config.GetRouteByPattern(initialRoutePath)
    route.Pattern = endpoint.URLPath

    err = config.ModifyRoute(route.Id, *route, true)
    if err != nil {
        return internalServerError(resp, err.Error())
    }

    resp.StatusCode = http.StatusOK
    resp.Message, _ = endpoint.SerializeJson()

    return nil
}

func (api *Api) DeleteEndpoint(vars *ApiVar, resp *ApiResponse) error {
    endpointId, err, found := filter.GetIdFromParams(vars.RequestForm)

    if found {
        if err != nil {
            return badRequest(resp, err.Error())
        }

        endpoint, _ := service.GetEndpoint(endpointId)

        err = service.DeleteEndpoint(endpointId)
        if err != nil {
            return notFound(resp, err.Error())
        }

        route := config.GetRouteByPattern(endpoint.URLPath)
        err = config.RemoveRoute(route.Id, true)

        if err != nil {
            return internalServerError(resp, err.Error())
        }

        resp.StatusCode = http.StatusOK
        return nil
    }

    return badRequest(resp, err.Error())
}

func generateNewRoute(endpoint *models.Endpoint, resp *ApiResponse) error {
    route := &config.Route{
        Id:      random.RandomString(5),
        Pattern: endpoint.URLPath,
        Handlers: map[string]string{
            "GET":    endpoint.GET.Function,
            "POST":   endpoint.POST.Function,
            "PUT":    endpoint.PUT.Function,
            "DELETE": endpoint.DELETE.Function,
        },
    }

    return config.AddRoute(route, true)
}
