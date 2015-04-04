package dbmodels

import (
    "apiGO/interfaces"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
)

type Endpoint struct {
    Id  bson.ObjectId `bson:"_id" json:"id"`

    URLPath        string                      `bson:"urlPath" json:"urlPath"`
    UserId         bson.ObjectId               `bson:"userId,omitempty" json:"userId"`
    Name           string                      `bson:"name" json:"name"`
    Description    string                      `bson:"description" json:"description"`
    Authentication EndpointAuth                `bson:"authentication" json:"authentication"`
    Enabled        bool                        `bson:"enabled" json:"enabled"`
    REST           map[string]EndpointResponse `bson:"rest" json:"rest"`
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
    case endpoint.Enabled != otherEndpoint.Enabled:
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
