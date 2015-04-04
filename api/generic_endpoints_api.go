package api

import (
    "apiGO/dbmodels"
    "apiGO/service"
    "time"
)

func (api *Api) GenericGET(vars *ApiVar, resp *ApiResponse) error {
    validateAndGetEndpoint(vars, resp)

    return nil
}

func (api *Api) GenericPOST(vars *ApiVar, resp *ApiResponse) error {
    validateAndGetEndpoint(vars, resp)

    return nil
}

func (api *Api) GenericPUT(vars *ApiVar, resp *ApiResponse) error {
    validateAndGetEndpoint(vars, resp)

    return nil
}

func (api *Api) GenericDELETE(vars *ApiVar, resp *ApiResponse) error {
    validateAndGetEndpoint(vars, resp)

    return nil
}

func validateAndGetEndpoint(vars *ApiVar, resp *ApiResponse) *dbmodels.Endpoint {
    endpoint, err := service.GetEndpointByURLPath(vars.Route.Pattern)
    if err != nil || endpoint == nil {
        notFound(resp, err.Error())
        return nil
    }

    if endpoint.Enabled == false {
        serviceUnavailable(resp, "This endpoint is not enabled")
        return nil
    }

    if !performBasicAuth(endpoint, vars) {
        unauthorized(resp, "Basic authentication failed!")
        return nil
    }

    endpointResponse := endpoint.REST[vars.RequestMethod]

    resp.StatusCode = endpointResponse.StatusCode
    resp.Message = []byte(endpointResponse.Response)

    // delay the response
    time.Sleep(endpointResponse.Delay * time.Millisecond)

    return endpoint
}

func performBasicAuth(endpoint *dbmodels.Endpoint, vars *ApiVar) bool {
    switch {
    case !vars.BasicAuth.OK:
        return false
    case vars.BasicAuth.Username != endpoint.Authentication.Username:
        return false
    case vars.BasicAuth.Password != endpoint.Authentication.Password:
        return false
    }

    return true
}
