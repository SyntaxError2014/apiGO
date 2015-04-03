package dbmodels

import (
    "gopkg.in/mgo.v2/bson"
    "time"
)

type EndpointResponse struct {
    StatusCode int       `bson:"statusCode" json:"statusCode"`
    Delay      time.Time `bson:"delay" json:"delay"`
    Response   []byte    `bson:"response" json:"response"`
}
