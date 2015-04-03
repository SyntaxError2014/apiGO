package models

import (
    "apiGO/dbmodels"
    "apiGO/interfaces"
    "apiGO/service"
    "bytes"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
    "time"
)

type RequestHistory struct {
    Id  bson.ObjectId `bson:"_id" json:"id"`

    Endpoint   dbmodels.Endpoint `bson:"endpoint" json:"endpoint"`
    Time       time.Time         `bson:"time" json:"time"`
    HTTPMethod string            `bson:"httpMethod" json:"httpMethod"`
    Header     []byte            `bson:"header" json:"header"`
    Parameters []byte            `bson:"parameters" json:"parameters"`
    Body       []byte            `bson:"body" json:"body"`
}

func (requestHistory *RequestHistory) Equal(otherRequestHistory RequestHistory) bool {
    switch {
    case requestHistory.Id != otherRequestHistory.Id:
        return false
    case !requestHistory.Endpoint.Equal(otherRequestHistory.Endpoint):
        return false
    case !requestHistory.Time.Equal(otherRequestHistory.Time):
        return false
    case bytes.Compare(requestHistory.Header, otherRequestHistory.Header) != 0:
        return false
    case bytes.Compare(requestHistory.Parameters, otherRequestHistory.Parameters) != 0:
        return false
    case bytes.Compare(requestHistory.Body, otherRequestHistory.Body) != 0:
        return false
    }

    return true
}

func (requestHistory *RequestHistory) SerializeJson() ([]byte, error) {
    data, err := json.MarshalIndent(*requestHistory, interfaces.JsonPrefix, interfaces.JsonIndent)

    if err != nil {
        return nil, err
    }

    return data, nil
}

func (requestHistory *RequestHistory) DeserializeJson(obj []byte) error {
    err := json.Unmarshal(obj, requestHistory)

    if err != nil {
        return err
    }

    return nil
}

func (requestHistory *RequestHistory) Expand(baseRequestHistory dbmodels.RequestHistory) error {
    requestHistory.Id = baseRequestHistory.Id
    requestHistory.Time = baseRequestHistory.Time
    requestHistory.HTTPMethod = baseRequestHistory.HTTPMethod
    requestHistory.Header = baseRequestHistory.Header
    requestHistory.Parameters = baseRequestHistory.Parameters
    requestHistory.Body = baseRequestHistory.Body

    endpoint, err := service.GetEndpoint(baseRequestHistory.EndpointId)
    if err != nil {
        return err
    }

    requestHistory.Endpoint = *endpoint

    return nil
}

func (requestHistory *RequestHistory) Collapse() (*dbmodels.RequestHistory, error) {
    var collapsedRequestHistory = dbmodels.RequestHistory{
        Id:         requestHistory.Id,
        EndpointId: requestHistory.Endpoint.Id,
        Time:       requestHistory.Time,
        HTTPMethod: requestHistory.HTTPMethod,
        Header:     requestHistory.Header,
        Parameters: requestHistory.Parameters,
        Body:       requestHistory.Body,
    }

    return &collapsedRequestHistory, nil
}
