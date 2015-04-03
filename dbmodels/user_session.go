package dbmodels

import (
    "apiGO/interfaces"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
    "time"
)

type UserSession struct {
    Id  bson.ObjectId `bson:"_id" json:"id"`

    UserId bson.ObjectId `bson:"userId" json:"userId"`
    Token  string        `bson:"token" json:"token"`
    Time   time.Time     `bson:"time" json:"time"`
}
