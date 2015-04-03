package dbmodels

import (
    "apiGO/interfaces"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
)

type User struct {
    Id bson.ObjectId `bson:"_id" json:"id"`
}

func (user *User) Equal(otherUser User) bool {
    switch {
    case user.Id != otherUser.Id:
        return false
    }

    return true
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
