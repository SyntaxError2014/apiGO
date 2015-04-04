package dbmodels

import (
    "strings"
    "time"
)

type EndpointResponse struct {
    StatusCode int           `bson:"statusCode" json:"statusCode"`
    Delay      time.Duration `bson:"delay" json:"delay"`
    Response   string        `bson:"response" json:"response"`

    function string
}

func (endpointResponse EndpointResponse) Equal(otherEndpointResponse EndpointResponse) bool {
    switch {
    case endpointResponse.StatusCode != otherEndpointResponse.StatusCode:
        return false
    case endpointResponse.Delay != otherEndpointResponse.Delay:
        return false
    case endpointResponse.Response != otherEndpointResponse.Response:
        return false
    case endpointResponse.function != otherEndpointResponse.function:
        return false
    }

    return true
}

func (endpointResponse EndpointResponse) GetApiFunction() string {
    return endpointResponse.function
}

func NewEndpointResponse(method string) EndpointResponse {
    basicEndpointResponse := EndpointResponse{
        StatusCode: 200,
        Delay:      0,
        Response:   "Hello world!",
        function:   strings.Join([]string{"Api.Generic", method}, ""),
    }

    return basicEndpointResponse
}
