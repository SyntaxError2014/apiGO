package service

import (
    "apiGO/dbmodels"
    "gopkg.in/mgo.v2/bson"
)

func GetEndpointByURLPath(urlPath string) (*dbmodels.Endpoint, error) {
    session, collection := Connect(EndpointsCollectionName)
    defer session.Close()

    var endpoint dbmodels.Endpoint
    err := collection.Find(bson.M{"urlPath": urlPath}).One(&endpoint)

    return &endpoint, err
}

func GetAllEndpointsForUser(userId bson.ObjectId) ([]dbmodels.Endpoint, error) {
    session, collection := Connect(EndpointsCollectionName)
    defer session.Close()

    var endpoints []dbmodels.Endpoint
    err := collection.Find(bson.M{"userId": userId}).All(&endpoints)

    return endpoints, err
}
