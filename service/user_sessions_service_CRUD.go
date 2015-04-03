package service

import (
    "apiGO/dbmodels"
    "gopkg.in/mgo.v2/bson"
)

func CreateUserSession(userSession *dbmodels.UserSession) (*dbmodels.UserSession, error) {
    session, collection := Connect(UserSessionsCollectionName)
    defer session.Close()

    if userSession.Id == "" {
        userSession.Id = bson.NewObjectId()
    }

    err := collection.Insert(userSession)

    return userSession, err
}

func UpdateUserSession(userSession *dbmodels.UserSession) error {
    session, collection := Connect(UserSessionsCollectionName)
    defer session.Close()

    err := collection.UpdateId(userSession.Id, userSession)

    return err
}

func DeleteUserSession(userSessionId bson.ObjectId) error {
    session, collection := Connect(UserSessionsCollectionName)
    defer session.Close()

    err := collection.RemoveId(userSessionId)

    return err
}

func GetUserSession(userSessionId bson.ObjectId) (*dbmodels.UserSession, error) {
    session, collection := Connect(UserSessionsCollectionName)
    defer session.Close()

    userSession := dbmodels.UserSession{}
    err := collection.FindId(userSessionId).One(&userSession)

    return &userSession, err
}

func GetAllUserSessions() ([]dbmodels.UserSession, error) {
    session, collection := Connect(UserSessionsCollectionName)
    defer session.Close()

    var userSessions []dbmodels.UserSession
    err := collection.Find(bson.M{}).All(&userSessions)

    return userSessions, err
}

func GetAllUserSessionsLimited(limit int) ([]dbmodels.UserSession, error) {
    session, collection := Connect(UserSessionsCollectionName)
    defer session.Close()

    var userSessions []dbmodels.UserSession
    err := collection.Find(bson.M{}).Limit(limit).All(&userSessions)

    return userSessions, err
}
