package api

import (
    "apiGO/dbmodels"
    "apiGO/filter"
    "apiGO/interfaces"
    "apiGO/models"
    "apiGO/service"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
    "net/http"
    "strings"
)

func (api *Api) GetEndpoint(vars *ApiVar, resp *ApiResponse) error {
    endpointId, err, found := filter.GetIdFromParams(vars.RequestForm)
    if found {
        if err != nil {
            return badRequest(resp, err.Error())
        }

        return getEndpoint(vars, resp, endpointId)
    }

    limit, err, found := filter.GetIntValueFromParams("limit", vars.RequestForm)
    if found {
        if err != nil {
            return badRequest(resp, err.Error())
        }

        return getAllEndpoints(vars, resp, limit)
    }

    return getAllEndpoints(vars, resp, -1)

}

func (api *Api) PostEndpoint(vars *ApiVar, resp *ApiResponse) error {
    basicEndpoint := &dbmodels.Endpoint{
        URLPath: strings.Join([]string{"/", bson.NewObjectId().Hex()}, ""),
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

    err = service.UpdateEndpoint(endpoint)
    if err != nil {
        return notFound(resp, "The endpoint with the specified id could not be found")
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

        err = service.DeleteEndpoint(endpointId)
        if err != nil {
            return notFound(resp, err.Error())
        }

        resp.StatusCode = http.StatusOK
        return nil
    }

    return badRequest(resp, err.Error())
}

func getAllEndpoints(vars *ApiVar, resp *ApiResponse, limit int) error {
    var endpoints []dbmodels.Endpoint
    var err error

    if limit == 0 {
        return badRequest(resp, "The limit cannot be 0. Use the value -1 for retrieving all the entities")
    }

    if limit != -1 {
        endpoints, err = service.GetAllEndpointsLimited(limit)
    } else {
        endpoints, err = service.GetAllEndpoints()
    }

    if err != nil {
        return internalServerError(resp, err.Error())
    }

    var expandedEndpoints []models.Endpoint
    expandedEndpoints = make([]models.Endpoint, len(endpoints))

    for i := 0; i < len(endpoints); i++ {
        err = expandedEndpoints[i].Expand(endpoints[i])

        if err != nil {
            return internalServerError(resp, err.Error())
        }
    }

    endpointsJson, err := json.MarshalIndent(expandedEndpoints, interfaces.JsonPrefix, interfaces.JsonIndent)

    if err != nil {
        return internalServerError(resp, err.Error())
    }

    resp.StatusCode = http.StatusOK
    resp.Message = endpointsJson

    return nil
}

func getEndpoint(vars *ApiVar, resp *ApiResponse, endpointId bson.ObjectId) error {
    endpoint, err := service.GetEndpoint(endpointId)

    if err != nil {
        if err.Error() == "not found" {
            return notFound(resp, err.Error())
        } else {
            return internalServerError(resp, err.Error())
        }
    }

    if endpoint == nil {
        return notFound(resp, "No Endpoint with the selected id was found")
    }

    var expandedEndpoint models.Endpoint
    err = expandedEndpoint.Expand(*endpoint)

    if err != nil {
        return internalServerError(resp, err.Error())
    }

    endpointJson, err := expandedEndpoint.SerializeJson()
    if err != nil {
        return internalServerError(resp, err.Error())
    }

    resp.StatusCode = http.StatusOK
    resp.Message = endpointJson

    return nil
}
