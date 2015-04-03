package models

import (
    "apiGO/dbmodels"
    "apiGO/interfaces"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
)

type User struct {
    Id bson.ObjectId `json:"id"`
}

func (user *User) SerializeJson() ([]byte, error) {
    data, err := json.MarshalIndent(*user, interfaces.JsonPrefix, interfaces.JsonIndent)

    if err != nil {
        return nil, err
    }

    return data, nil
}

func (user *User) DeserializeJson(obj []byte) error {
    err := json.Unmarshal(obj, user)

    if err != nil {
        return err
    }

    return nil
}

func (user *User) Expand(baseUser dbmodels.User) error {
    user.Id = baseUser.Id
    return nil
}

func (user *User) Collapse() (*dbmodels.User, error) {
    var collapsedUser = dbmodels.User{
        Id: user.Id,
    }

    return &collapsedUser, nil
}
