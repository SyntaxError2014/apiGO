package api

import (
    "apiGO/dbmodels"
    py "apiGO/python_integration"
    "apiGO/service"
    "net/http"
    "time"
)

func (api *Api) GenericGET(vars *ApiVar, resp *ApiResponse) error {
    endpoint := validateAndGetEndpoint(vars, resp)

    if endpoint != nil {
        // executeSourceCode(endpoint, vars, resp)
    }

    return nil
}

func (api *Api) GenericPOST(vars *ApiVar, resp *ApiResponse) error {
    endpoint := validateAndGetEndpoint(vars, resp)

    if endpoint != nil {
        // executeSourceCode(endpoint, vars, resp)
    }

    return nil
}

func (api *Api) GenericPUT(vars *ApiVar, resp *ApiResponse) error {
    endpoint := validateAndGetEndpoint(vars, resp)

    if endpoint != nil {
        // executeSourceCode(endpoint, vars, resp)
    }

    return nil
}

func (api *Api) GenericDELETE(vars *ApiVar, resp *ApiResponse) error {
    endpoint := validateAndGetEndpoint(vars, resp)

    if endpoint != nil {
        // executeSourceCode(endpoint, vars, resp)
    }

    return nil
}

func executeSourceCode(endpoint *dbmodels.Endpoint, vars *ApiVar, resp *ApiResponse) bool {
    endpointResponse := endpoint.REST[vars.RequestMethod]

    if len(endpointResponse.SourceCode) > 0 {
        output, err := py.ExecuteCommand(endpointResponse.SourceCode, vars.RequestForm)

        if err != nil {
            internalServerError(resp, err.Error())
            return false
        }

        statusCode, responseMessage := py.ParseOutput(output)

        resp.StatusCode = statusCode
        resp.Message = []byte(responseMessage)
    }

    return true
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
        requestHistory.ResponseMessage = msg
        requestHistory.ResponseContentType = "text/plain"
        service.CreateRequestHistory(requestHistory)

        serviceUnavailable(resp, msg)
        return nil
    }

    if !performBasicAuth(endpoint, vars) {
        msg := "Basic authentication failed!"

        requestHistory.ResponseStatusCode = http.StatusServiceUnavailable
        requestHistory.ResponseMessage = msg
        requestHistory.ResponseContentType = "text/plain"
        service.CreateRequestHistory(requestHistory)

        unauthorized(resp, msg)
        return nil
    }

    // delay the response
    time.Sleep(endpoint.REST[vars.RequestMethod].Delay * time.Millisecond)

    return parseEndpoint(endpoint, requestHistory, vars, resp)
}

func parseEndpoint(endpoint *dbmodels.Endpoint, requestHistory *dbmodels.RequestHistory, vars *ApiVar, resp *ApiResponse) *dbmodels.Endpoint {
    endpointResponse := endpoint.REST[vars.RequestMethod]

    // Set the response
    resp.StatusCode = endpointResponse.StatusCode
    resp.Message = []byte(endpointResponse.Response)
    resp.ContentType = endpointResponse.ContentType

    // Set the response history and add it to the database
    requestHistory.ResponseStatusCode = resp.StatusCode
    requestHistory.ResponseMessage = string(resp.Message)
    requestHistory.ResponseContentType = resp.ContentType
    service.CreateRequestHistory(requestHistory)

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
        Body:                string(vars.RequestBody),
        ResponseContentType: vars.RequestHeader.Get("Content-Type"),
    }

    return &dbHistory
}
