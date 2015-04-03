package dbmodels

import (
    "bytes"
    "time"
)

type EndpointResponse struct {
    StatusCode int       `bson:"statusCode" json:"statusCode"`
    Delay      time.Time `bson:"delay" json:"delay"`
    Response   []byte    `bson:"response" json:"response"`
}

func (endpointResponse *EndpointResponse) Equal(otherEndpointResponse EndpointResponse) bool {
    switch {
    case endpointResponse.StatusCode != otherEndpointResponse.StatusCode:
        return false
    case !endpointResponse.Delay.Equal(otherEndpointResponse.Delay):
        return false
    case bytes.Compare(endpointResponse.Response, otherEndpointResponse.Response) != 0:
        return false
    }

    return true
}
