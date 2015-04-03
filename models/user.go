package models

import (
    "apiGO/dbmodels"
    "apiGO/interfaces"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
)

type User struct {
    Id  bson.ObjectId `json:"id"`

    Username   string `json:"username"`
    Password   string `json:"password"`
    FirstName  string `json:"firstName"`
    LastName   string `json:"lastName"`
    Email      string `json:"email"`
    FacebookId string `json:"facebookId"`
    GoogleId   string `json:"googleId"`
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
    user.Username = baseUser.Username
    user.Password = baseUser.Password
    user.FirstName = baseUser.FirstName
    user.LastName = baseUser.LastName
    user.Email = baseUser.Email
    user.FacebookId = baseUser.FacebookId
    user.GoogleId = baseUser.GoogleId
    return nil
}

func (user *User) Collapse() (*dbmodels.User, error) {
    var collapsedUser = dbmodels.User{
        Id:         user.Id,
        Username:   user.Username,
        Password:   user.Password,
        FirstName:  user.FirstName,
        LastName:   user.LastName,
        Email:      user.Email,
        FacebookId: user.FacebookId,
        GoogleId:   user.GoogleId,
    }

    return &collapsedUser, nil
}
