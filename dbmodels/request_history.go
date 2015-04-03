package dbmodels

import (
    "apiGO/interfaces"
    "bytes"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
    "time"
)

type RequestHistory struct {
    Id  bson.ObjectId `bson:"_id" json:"id"`

    EndpointId bson.ObjectId `bson:"endpointId,omitempty" json:"endpointId"`
    Time       time.Time     `bson:"time" json:"time"`
    HTTPMethod string        `bson:"httpMethod" json:"httpMethod"`
    Header     []byte        `bson:"header" json:"header"`
    Parameters []byte        `bson:"parameters" json:"parameters"`
    Body       []byte        `bson:"body" json:"body"`
}

func (requestHistory *RequestHistory) Equal(otherRequestHistory RequestHistory) bool {
    switch {
    case requestHistory.Id != otherRequestHistory.Id:
        return false
    case requestHistory.EndpointId != otherRequestHistory.EndpointId:
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
