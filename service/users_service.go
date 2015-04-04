package service

import (
    "apiGO/dbmodels"
    "gopkg.in/mgo.v2/bson"
)

func GetUserByUsernameAndPassword(username, password string) (*dbmodels.User, error) {
    session, collection := Connect(UsersCollectionName)
    defer session.Close()

    user := dbmodels.User{}

    err := collection.Find(bson.M{
        "username": username,
        "password": password,
    }).One(&user)

    return &user, err
}
