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

func (userSession *UserSession) Equal(otherUserSession UserSession) bool {
    switch {
    case userSession.Id != otherUserSession.Id:
        return false
    case userSession.UserId != otherUserSession.UserId:
        return false
    case userSession.Token != otherUserSession.Token:
        return false
    case !userSession.Time.Equal(otherUserSession.Time):
        return false
    }

    return true
}

func (userSession *UserSession) SerializeJson() ([]byte, error) {
    data, err := json.MarshalIndent(*userSession, interfaces.JsonPrefix, interfaces.JsonIndent)

    if err != nil {
        return nil, err
    }

    return data, nil
}

func (userSession *UserSession) DeserializeJson(obj []byte) error {
    err := json.Unmarshal(obj, userSession)

    if err != nil {
        return err
    }

    return nil
}
