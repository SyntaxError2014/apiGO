package models

import (
    "apiGO/dbmodels"
    "apiGO/interfaces"
    "apiGO/service"
    "bytes"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
    "net/url"
    "time"
)

type RequestHistory struct {
    Id  bson.ObjectId `json:"id"`

    Endpoint            dbmodels.Endpoint   `json:"endpoint"`
    RequestDate         time.Time           `json:"requestDate"`
    HTTPMethod          string              `json:"httpMethod"`
    Header              map[string][]string `json:"header"`
    Parameters          url.Values          `json:"parameters"`
    Body                []byte              `json:"body"`
    ResponseStatusCode  int                 `bson:"responseStatusCode" json:"responseStatusCode"`
    ResponseMessage     []byte              `bson:"responseMessage" json:"responseMessage"`
    ResponseContentType string              `bson:"responseContentType" json:"responseContentType"`
}

func (requestHistory *RequestHistory) Equal(otherRequestHistory RequestHistory) bool {
    switch {
    case requestHistory.Id != otherRequestHistory.Id:
        return false
    case !requestHistory.Endpoint.Equal(otherRequestHistory.Endpoint):
        return false
    case !requestHistory.RequestDate.Equal(otherRequestHistory.RequestDate):
        return false
    // case bytes.Compare(requestHistory.Header, otherRequestHistory.Header) != 0:
    //     return false
    // case bytes.Compare(requestHistory.Parameters, otherRequestHistory.Parameters) != 0:
    //     return false
    case bytes.Compare(requestHistory.Body, otherRequestHistory.Body) != 0:
        return false
    case requestHistory.ResponseStatusCode != otherRequestHistory.ResponseStatusCode:
        return false
    case bytes.Compare(requestHistory.ResponseMessage, otherRequestHistory.ResponseMessage) != 0:
        return false
    case requestHistory.ResponseContentType != otherRequestHistory.ResponseContentType:
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
    requestHistory.RequestDate = baseRequestHistory.RequestDate
    requestHistory.HTTPMethod = baseRequestHistory.HTTPMethod
    requestHistory.Header = baseRequestHistory.Header
    requestHistory.Parameters = baseRequestHistory.Parameters
    requestHistory.Body = baseRequestHistory.Body
    requestHistory.ResponseStatusCode = baseRequestHistory.ResponseStatusCode
    requestHistory.ResponseMessage = baseRequestHistory.ResponseMessage
    requestHistory.ResponseContentType = baseRequestHistory.ResponseContentType

    endpoint, err := service.GetEndpoint(baseRequestHistory.EndpointId)
    if err != nil {
        return err
    }

    requestHistory.Endpoint = *endpoint

    return nil
}

func (requestHistory *RequestHistory) Collapse() (*dbmodels.RequestHistory, error) {
    var collapsedRequestHistory = dbmodels.RequestHistory{
        Id:                  requestHistory.Id,
        EndpointId:          requestHistory.Endpoint.Id,
        RequestDate:         requestHistory.RequestDate,
        HTTPMethod:          requestHistory.HTTPMethod,
        Header:              requestHistory.Header,
        Parameters:          requestHistory.Parameters,
        Body:                requestHistory.Body,
        ResponseStatusCode:  requestHistory.ResponseStatusCode,
        ResponseMessage:     requestHistory.ResponseMessage,
        ResponseContentType: requestHistory.ResponseContentType,
    }

    return &collapsedRequestHistory, nil
}
