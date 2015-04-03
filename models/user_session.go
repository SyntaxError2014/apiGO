package models

import (
    "apiGO/dbmodels"
    "apiGO/interfaces"
    "apiGO/service"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
    "time"
)

type UserSession struct {
    Id  bson.ObjectId `bson:"_id" json:"id"`

    User  dbmodels.User `bson:"user" json:"user"`
    Token string        `bson:"token" json:"token"`
    Time  time.Time     `bson:"time" json:"time"`
}

func (userSession *UserSession) Equal(otherUserSession UserSession) bool {
    switch {
    case userSession.Id != otherUserSession.Id:
        return false
    case !userSession.User.Equal(otherUserSession.User):
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

func (userSession *UserSession) Expand(baseUserSession dbmodels.UserSession) error {
    userSession.Id = baseUserSession.Id
    userSession.Token = baseUserSession.Token
    userSession.Time = baseUserSession.Time

    user, err := service.GetUser(baseUserSession.UserId)
    if err != nil {
        return err
    }

    userSession.User = *user

    return nil
}

func (userSession *UserSession) Collapse() (*dbmodels.UserSession, error) {
    var collapsedUserSession = dbmodels.UserSession{
        Id:     userSession.Id,
        UserId: userSession.User.Id,
        Token:  userSession.Token,
        Time:   userSession.Time,
    }

    return &collapsedUserSession, nil
}
