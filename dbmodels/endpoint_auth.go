package dbmodels

import (
    "gopkg.in/mgo.v2/bson"
)

type EndpointAuth struct {
    UserId   bson.ObjectId `bson:"userId" json:"userId"`
    Password string        `bson:"password" json:"password"`
}
