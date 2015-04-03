package models

import (
    "apiGO/dbmodels"
    "apiGO/interfaces"
    "apiGO/service"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
)

type Endpoint struct {
    Id  bson.ObjectId `bson:"_id" json:"id"`

    URLPath        string                    `bson:"urlPath" json:"urlPath"`
    User           dbmodels.User             `bson:"user" json:"user"`
    Name           string                    `bson:"name" json:"name"`
    Description    string                    `bson:"description" json:"description"`
    Authentication dbmodels.EndpointAuth     `bson:"authentication" json:"authentication"`
    GET            dbmodels.EndpointResponse `bson:"get" json:"get"`
    POST           dbmodels.EndpointResponse `bson:"post" json:"post"`
    PUT            dbmodels.EndpointResponse `bson:"put" json:"put"`
    DELETE         dbmodels.EndpointResponse `bson:"delete" json:"delete"`
}

func (endpoint *Endpoint) Equal(otherEndpoint Endpoint) bool {
    switch {
    case endpoint.Id != otherEndpoint.Id:
        return false
    case endpoint.URLPath != otherEndpoint.URLPath:
        return false
    case !endpoint.User.Equal(otherEndpoint.User):
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

func (endpoint *Endpoint) Expand(baseEndpoint dbmodels.Endpoint) error {
    endpoint.Id = baseEndpoint.Id
    endpoint.Name = baseEndpoint.Name
    endpoint.Description = baseEndpoint.Description
    endpoint.Authentication = baseEndpoint.Authentication
    endpoint.GET = baseEndpoint.GET
    endpoint.POST = baseEndpoint.POST
    endpoint.PUT = baseEndpoint.PUT
    endpoint.DELETE = baseEndpoint.DELETE

    user, err := service.GetUser(baseEndpoint.UserId)
    if err != nil {
        return err
    }

    endpoint.User = *user

    return nil
}

func (endpoint *Endpoint) Collapse() (*dbmodels.Endpoint, error) {
    var collapsedEndpoint = dbmodels.Endpoint{
        Id:             endpoint.Id,
        UserId:         endpoint.User.Id,
        Name:           endpoint.Name,
        Description:    endpoint.Description,
        Authentication: endpoint.Authentication,
        GET:            endpoint.GET,
        POST:           endpoint.POST,
        PUT:            endpoint.PUT,
        DELETE:         endpoint.DELETE,
    }

    return &collapsedEndpoint, nil
}