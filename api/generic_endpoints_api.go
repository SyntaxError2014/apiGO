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
    endpoint := validateAndGetEndpoint(vars, resp)

    if endpoint != nil {

    }

    return nil
}

func (api *Api) GenericPUT(vars *ApiVar, resp *ApiResponse) error {
    endpoint := validateAndGetEndpoint(vars, resp)

    if endpoint != nil {

    }

    return nil
}

func (api *Api) GenericDELETE(vars *ApiVar, resp *ApiResponse) error {
    endpoint := validateAndGetEndpoint(vars, resp)

    if endpoint != nil {

    }

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

    response := endpoint.REST[vars.RequestMethod]

    resp.StatusCode = response.StatusCode
    resp.Message = []byte(response.Response)

    // delay response
    time.Sleep(response.Delay * time.Millisecond)

    return endpoint
}
