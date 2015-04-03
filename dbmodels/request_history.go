package dbmodels

import (
    "apiGO/interfaces"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
    "time"
)

type RequestHistory struct {
    Id  bson.ObjectId `bson:"_id" json:"id"`

    EndpointId bson.ObjectId `bson:"endpointId" json:"endpointId"`
    Time       time.Time     `bson:"time" json:"time"`
    HTTPMethod string        `bson:"httpMethod" json:"httpMethod"`
    Header     []byte        `bson:"header" json:"header"`
    Parameters []byte        `bson:"parameters" json:"parameters"`
    Body       []byte        `bson:"body" json:"body"`
}
