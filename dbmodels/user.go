package dbmodels

import (
    "apiGO/interfaces"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
)

type User struct {
    Id  bson.ObjectId `bson:"_id" json:"id"`

    Username   string `bson:"username" json:"username"`
    Password   string `bson:"password" json:"password"`
    FirstName  string `bson:"firstName" json:"firstName"`
    LastName   string `bson:"lastName" json:"lastName"`
    Email      string `bson:"email" json:"email"`
    FacebookId string `bson:"facebookId" json:"facebookId"`
    GoogleId   string `bson:"googleId" json:"googleId"`
}

func (user *User) Equal(otherUser User) bool {
    switch {
    case user.Id != otherUser.Id:
        return false
    case user.Username != otherUser.Username:
        return false
    case user.Password != otherUser.Password:
        return false
    case user.FirstName != otherUser.FirstName:
        return false
    case user.LastName != otherUser.LastName:
        return false
    case user.Email != otherUser.Email:
        return false
    case user.FacebookId != otherUser.FacebookId:
        return false
    case user.GoogleId != otherUser.GoogleId:
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
