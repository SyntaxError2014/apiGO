package dbmodels

import (
    "strings"
    "time"
)

type EndpointResponse struct {
    StatusCode int           `bson:"statusCode" json:"statusCode"`
    Delay      time.Duration `bson:"delay" json:"delay"`
    Response   string        `bson:"response" json:"response"`
    Function   string        `bson:"function" json:"function"`
}

func (endpointResponse EndpointResponse) Equal(otherEndpointResponse EndpointResponse) bool {
    switch {
    case endpointResponse.StatusCode != otherEndpointResponse.StatusCode:
        return false
    case endpointResponse.Delay != otherEndpointResponse.Delay:
        return false
    case endpointResponse.Response != otherEndpointResponse.Response:
        return false
    case endpointResponse.Function != otherEndpointResponse.Function:
        return false
    }

    return true
}

func NewEndpointResponse(method string) EndpointResponse {
    basicEndpointResponse := EndpointResponse{
        StatusCode: 200,
        Delay:      0,
        Response:   "Hello world!",
        Function:   strings.Join([]string{"Api.Generic", method}, ""),
    }

    return basicEndpointResponse
}
