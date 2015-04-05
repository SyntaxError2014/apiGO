package dbmodels

import (
    "apiGO/interfaces"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
    "net/url"
    "time"
)

type RequestHistory struct {
    Id  bson.ObjectId `bson:"_id" json:"id"`

    EndpointId          bson.ObjectId       `bson:"endpointId,omitempty" json:"endpointId"`
    RequestDate         time.Time           `bson:"requestDate" json:"requestDate"`
    HTTPMethod          string              `bson:"httpMethod" json:"httpMethod"`
    Header              map[string][]string `bson:"header" json:"header"`
    Parameters          url.Values          `bson:"parameters" json:"parameters"`
    Body                string              `bson:"body" json:"body"`
    ResponseStatusCode  int                 `bson:"responseStatusCode" json:"responseStatusCode"`
    ResponseMessage     string              `bson:"responseMessage" json:"responseMessage"`
    ResponseContentType string              `bson:"responseContentType" json:"responseContentType"`
}

func (requestHistory *RequestHistory) Equal(otherRequestHistory RequestHistory) bool {
    switch {
    case requestHistory.Id != otherRequestHistory.Id:
        return false
    case requestHistory.EndpointId != otherRequestHistory.EndpointId:
        return false
    case !requestHistory.RequestDate.Equal(otherRequestHistory.RequestDate):
        return false
    case requestHistory.Body != otherRequestHistory.Body:
        return false
    case requestHistory.ResponseStatusCode != otherRequestHistory.ResponseStatusCode:
        return false
    case requestHistory.ResponseMessage != otherRequestHistory.ResponseMessage:
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
