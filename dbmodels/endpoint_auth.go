package dbmodels

import (
    "gopkg.in/mgo.v2/bson"
)

type EndpointAuth struct {
    UserId   bson.ObjectId `bson:"userId,omitempty" json:"userId"`
    Password string        `bson:"password" json:"password"`
}

func (endpointAuth *EndpointAuth) Equal(otherEndpointAuth EndpointAuth) bool {
    switch {
    case endpointAuth.UserId != otherEndpointAuth.UserId:
        return false
    case endpointAuth.Password != otherEndpointAuth.Password:
        return false
    }

    return true
}
