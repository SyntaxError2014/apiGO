package api

import (
    "apiGO/config"
    "apiGO/filter"
    "apiGO/interfaces"
    "apiGO/models"
    "apiGO/service"
    "encoding/json"
    "net/http"
)

func (api *Api) GetEndpoint(vars *ApiVar, resp *ApiResponse) error {
    endpoints, err := service.GetAllEndpoints()
    if err != nil {
        return internalServerError(resp, err.Error())
    }

    expandedEndpoints := make([]models.Endpoint, len(endpoints))
    for i := 0; i < len(endpoints); i++ {
        expandedEndpoints[i].Expand(endpoints[i])
    }

    endpointsJson, err := json.MarshalIndent(expandedEndpoints, interfaces.JsonPrefix, interfaces.JsonIndent)

    if err != nil {
        return internalServerError(resp, err.Error())
    }

    resp.StatusCode = http.StatusOK
    resp.Message = endpointsJson

    return nil
}

func (api *Api) PostEndpoint(vars *ApiVar, resp *ApiResponse) error {
    basicEndpoint := generateNewEndpoint()

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

    err = updateRoute(initialRoutePath, endpoint)
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
