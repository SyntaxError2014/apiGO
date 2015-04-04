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
