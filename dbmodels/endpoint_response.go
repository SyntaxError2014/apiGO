package dbmodels

import (
    "bytes"
)

type EndpointResponse struct {
    StatusCode int    `bson:"statusCode" json:"statusCode"`
    Delay      int    `bson:"delay" json:"delay"`
    Response   []byte `bson:"response" json:"response"`
}

func (endpointResponse *EndpointResponse) Equal(otherEndpointResponse EndpointResponse) bool {
    switch {
    case endpointResponse.StatusCode != otherEndpointResponse.StatusCode:
        return false
    case endpointResponse.Delay != otherEndpointResponse.Delay:
        return false
    case bytes.Compare(endpointResponse.Response, otherEndpointResponse.Response) != 0:
        return false
    }

    return true
}

func NewEndpointResponse() EndpointResponse {
    basicEndpointResponse := EndpointResponse{
        StatusCode: 200,
        Delay:      0,
        Response:   []byte("Hello world!"),
    }

    return basicEndpointResponse
}
