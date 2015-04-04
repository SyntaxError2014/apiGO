package api

import (
    "apiGO/dbmodels"
    "apiGO/service"
    "net/http"
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

    requestHistory := generateAccessHistory(endpoint, vars)

    if endpoint.Enabled == false {
        msg := "This endpoint is not enabled"

        requestHistory.ResponseStatusCode = http.StatusServiceUnavailable
        requestHistory.ResponseMessage = []byte(msg)
        requestHistory.ResponseContentType = "text/plain"

        serviceUnavailable(resp, msg)
        return nil
    }

    if !performBasicAuth(endpoint, vars) {
        msg := "Basic authentication failed!"

        requestHistory.ResponseStatusCode = http.StatusServiceUnavailable
        requestHistory.ResponseMessage = []byte(msg)
        requestHistory.ResponseContentType = "text/plain"

        unauthorized(resp, msg)
        return nil
    }

    endpointResponse := endpoint.REST[vars.RequestMethod]

    // Set the response
    resp.StatusCode = endpointResponse.StatusCode
    resp.Message = []byte(endpointResponse.Response)
    resp.ContentType = endpointResponse.ContentType

    // Set the response history and add it to the database
    requestHistory.ResponseStatusCode = resp.StatusCode
    requestHistory.ResponseMessage = resp.Message
    requestHistory.ResponseContentType = resp.ContentType
    service.CreateRequestHistory(requestHistory)

    // delay the response
    time.Sleep(endpointResponse.Delay * time.Millisecond)

    return endpoint
}

func performBasicAuth(endpoint *dbmodels.Endpoint, vars *ApiVar) bool {
    var usr string = endpoint.Authentication.Username
    var pass string = endpoint.Authentication.Password

    if len(usr) == 0 && len(pass) == 0 {
        return true
    }

    switch {
    case !vars.BasicAuth.OK:
        return false
    case vars.BasicAuth.Username != usr:
        return false
    case vars.BasicAuth.Password != pass:
        return false
    }

    return true
}

func generateAccessHistory(endpoint *dbmodels.Endpoint, vars *ApiVar) *dbmodels.RequestHistory {
    dbHistory := dbmodels.RequestHistory{
        EndpointId:          endpoint.Id,
        RequestDate:         time.Now().Local(),
        HTTPMethod:          vars.RequestMethod,
        Header:              vars.RequestHeader,
        Parameters:          vars.RequestForm,
        Body:                vars.RequestBody,
        ResponseContentType: vars.RequestHeader.Get("Content-Type"),
    }

    return &dbHistory
}
