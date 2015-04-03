package service

import (
    "apiGO/dbmodels"
    "gopkg.in/mgo.v2/bson"
)

func CreateUser(user *dbmodels.User) (*dbmodels.User, error) {
    session, collection := Connect(UsersCollectionName)
    defer session.Close()

    if user.Id == "" {
        user.Id = bson.NewObjectId()
    }

    err := collection.Insert(user)

    return user, err
}

func UpdateUser(user *dbmodels.User) error {
    session, collection := Connect(UsersCollectionName)
    defer session.Close()

    err := collection.UpdateId(user.Id, user)

    return err
}

func DeleteUser(userId bson.ObjectId) error {
    session, collection := Connect(UsersCollectionName)
    defer session.Close()

    err := collection.RemoveId(userId)

    return err
}

func GetUser(userId bson.ObjectId) (*dbmodels.User, error) {
    session, collection := Connect(UsersCollectionName)
    defer session.Close()

    user := dbmodels.User{}
    err := collection.FindId(userId).One(&user)

    return &user, err
}

func GetAllUsers() ([]dbmodels.User, error) {
    session, collection := Connect(UsersCollectionName)
    defer session.Close()

    var users []dbmodels.User
    err := collection.Find(bson.M{}).All(&users)

    return users, err
}

func GetAllUsersLimited(limit int) ([]dbmodels.User, error) {
    session, collection := Connect(UsersCollectionName)
    defer session.Close()

    var users []dbmodels.User
    err := collection.Find(bson.M{}).Limit(limit).All(&users)

    return users, err
}
