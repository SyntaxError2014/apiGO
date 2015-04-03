package service

import (
    "apiGO/dbmodels"
    "gopkg.in/mgo.v2/bson"
)

func GetUserSessionByToken(token string) (*dbmodels.UserSession, error) {
    session, collection := Connect(UserSessionsCollectionName)
    defer session.Close()

    userSession := dbmodels.UserSession{}
    err := collection.Find(bson.M{"token": token}).One(&userSession)

    return &userSession, err
}
