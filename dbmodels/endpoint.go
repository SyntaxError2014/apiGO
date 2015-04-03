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

func (endpoint *Endpoint) Equal(otherEndpoint Endpoint) bool {
    switch {
    case endpoint.Id != otherEndpoint.Id:
        return false
    case endpoint.URLPath != otherEndpoint.URLPath:
        return false
    case endpoint.UserId != otherEndpoint.UserId:
        return false
    case endpoint.Name != otherEndpoint.Name:
        return false
    case !endpoint.Authentication.Equal(otherEndpoint.Authentication):
        return false
    case !endpoint.GET.Equal(otherEndpoint.GET):
        return false
    case !endpoint.POST.Equal(otherEndpoint.POST):
        return false
    case !endpoint.PUT.Equal(otherEndpoint.PUT):
        return false
    case !endpoint.DELETE.Equal(otherEndpoint.DELETE):
        return false
    }

    return true
}

func (endpoint *Endpoint) SerializeJson() ([]byte, error) {
    data, err := json.MarshalIndent(*endpoint, interfaces.JsonPrefix, interfaces.JsonIndent)

    if err != nil {
        return nil, err
    }

    return data, nil
}

func (endpoint *Endpoint) DeserializeJson(obj []byte) error {
    err := json.Unmarshal(obj, endpoint)

    if err != nil {
        return err
    }

    return nil
}
