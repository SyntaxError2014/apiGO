package models

import (
    "apiGO/dbmodels"
    "apiGO/interfaces"
    "apiGO/service"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
    "time"
)

type Endpoint struct {
    Id  bson.ObjectId `json:"id"`

    URLPath        string                               `json:"urlPath"`
    User           dbmodels.User                        `json:"user"`
    Name           string                               `json:"name"`
    Description    string                               `json:"description"`
    Authentication dbmodels.EndpointAuth                `json:"authentication"`
    Enabled        bool                                 `json:"enabled"`
    DateCreated    time.Time                            `json:"dateCreated"`
    REST           map[string]dbmodels.EndpointResponse `json:"rest"`
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
    case endpoint.Enabled != otherEndpoint.Enabled:
        return false
    case !endpoint.DateCreated.Equal(otherEndpoint.DateCreated):
        return false
    case !endpoint.Authentication.Equal(otherEndpoint.Authentication):
        return false
    default:
        for method, response := range endpoint.REST {
            value, found := otherEndpoint.REST[method]

            if !found || !value.Equal(response) {
                return false
            }
        }
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
    endpoint.URLPath = baseEndpoint.URLPath
    endpoint.Name = baseEndpoint.Name
    endpoint.Description = baseEndpoint.Description
    endpoint.Authentication = baseEndpoint.Authentication
    endpoint.Enabled = baseEndpoint.Enabled
    endpoint.REST = baseEndpoint.REST

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
        URLPath:        endpoint.URLPath,
        UserId:         endpoint.User.Id,
        Name:           endpoint.Name,
        Enabled:        endpoint.Enabled,
        Description:    endpoint.Description,
        Authentication: endpoint.Authentication,
        REST:           endpoint.REST,
    }

    return &collapsedEndpoint, nil
}
