package dbmodels

import (
    "apiGO/interfaces"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
)

type Endpoint struct {
    Id  bson.ObjectId `bson:"_id" json:"id"`

    URLPath        string           `bson:"urlPath" json:"urlPath"`
    UserId         bson.ObjectId    `bson:"userId" json:"userId"`
    Name           string           `bson:"name" json:"name"`
    Description    string           `bson:"description" json:"description"`
    Authentication EndpointAuth     `bson:"authentication" json:"authentication"`
    GET            EndpointResponse `bson:"get" json:"get"`
    POST           EndpointResponse `bson:"post" json:"post"`
    PUT            EndpointResponse `bson:"put" json:"put"`
    DELETE         EndpointResponse `bson:"delete" json:"delete"`
}
