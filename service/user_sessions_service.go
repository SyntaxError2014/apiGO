package service

import (
    "apiGO/dbmodels"
    "apiGO/random"
    "gopkg.in/mgo.v2/bson"
    "time"
)

func GetUserSessionByToken(token string) (*dbmodels.UserSession, error) {
    session, collection := Connect(UserSessionsCollectionName)
    defer session.Close()

    userSession := dbmodels.UserSession{}
    err := collection.Find(bson.M{"token": token}).One(&userSession)

    return &userSession, err
}

func DeleteAllSessionsWithUserId(userId bson.ObjectId) error {
    session, collection := Connect(UserSessionsCollectionName)
    defer session.Close()

    _, err := collection.RemoveAll(bson.M{"userId": userId})

    return err
}

func GenerateAndInsertUserSession(userId bson.ObjectId) (*dbmodels.UserSession, error) {
    session, collection := Connect(UserSessionsCollectionName)
    defer session.Close()

    userSession := dbmodels.UserSession{
        Id:     bson.NewObjectId(),
        UserId: userId,
        Token:  random.RandomString(25),
        Time:   time.Now().Local(),
    }

    err := collection.Insert(userSession)

    return &userSession, err
}
