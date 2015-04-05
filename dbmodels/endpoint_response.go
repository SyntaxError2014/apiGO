package dbmodels

import (
    "strings"
    "time"
)

type EndpointResponse struct {
    StatusCode  int           `bson:"statusCode" json:"statusCode"`
    Delay       time.Duration `bson:"delay" json:"delay"`
    Response    string        `bson:"response" json:"response"`
    SourceCode  string        `bson:"sourceCode" json:"sourceCode"`
    ContentType string        `bson:"contentType" json:"contentType"`
}

func (endpointResponse EndpointResponse) Equal(otherEndpointResponse EndpointResponse) bool {
    switch {
    case endpointResponse.StatusCode != otherEndpointResponse.StatusCode:
        return false
    case endpointResponse.Delay != otherEndpointResponse.Delay:
        return false
    case endpointResponse.Response != otherEndpointResponse.Response:
        return false
    case endpointResponse.SourceCode != otherEndpointResponse.SourceCode:
        return false
    case endpointResponse.ContentType != otherEndpointResponse.ContentType:
        return false
    }

    return true
}

func (endpointResponse EndpointResponse) GetApiFunction(method string) string {
    return defineFunction(method)
}

func NewEndpointResponse(method string) EndpointResponse {
    basicEndpointResponse := EndpointResponse{
        StatusCode:  200,
        Delay:       0,
        Response:    "Hello world!",
        ContentType: "text/plain",
    }

    return basicEndpointResponse
}

func defineFunction(method string) string {
    return strings.Join([]string{"Api.Generic", method}, "")
}
