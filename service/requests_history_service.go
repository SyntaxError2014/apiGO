package service

import (
    "apiGO/dbmodels"
    "gopkg.in/mgo.v2/bson"
)

func GetEntireRequestHistoryForEndpoint(endpointId bson.ObjectId) ([]dbmodels.RequestHistory, error) {
    session, collection := Connect(RequestsHistoryCollectionName)
    defer session.Close()

    var requestsHistory []dbmodels.RequestHistory
    err := collection.Find(bson.M{"endpointId": endpointId}).All(&requestsHistory)

    return requestsHistory, err
}
