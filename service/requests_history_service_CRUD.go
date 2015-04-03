package service

import (
    "apiGO/dbmodels"
    "gopkg.in/mgo.v2/bson"
)

const RequestHistorysCollectionName = "requests_history"

func CreateRequestHistory(requestHistory *dbmodels.RequestHistory) (*dbmodels.RequestHistory, error) {
    session, collection := Connect(RequestHistorysCollectionName)
    defer session.Close()

    if requestHistory.Id == "" {
        requestHistory.Id = bson.NewObjectId()
    }

    err := collection.Insert(requestHistory)

    return requestHistory, err
}

func UpdateRequestHistory(requestHistory *dbmodels.RequestHistory) error {
    session, collection := Connect(RequestHistorysCollectionName)
    defer session.Close()

    err := collection.UpdateId(requestHistory.Id, requestHistory)

    return err
}

func DeleteRequestHistory(requestHistoryId bson.ObjectId) error {
    session, collection := Connect(RequestHistorysCollectionName)
    defer session.Close()

    err := collection.RemoveId(requestHistoryId)

    return err
}

func GetRequestHistory(requestHistoryId bson.ObjectId) (*dbmodels.RequestHistory, error) {
    session, collection := Connect(RequestHistorysCollectionName)
    defer session.Close()

    requestHistory := dbmodels.RequestHistory{}
    err := collection.FindId(requestHistoryId).One(&requestHistory)

    return &requestHistory, err
}

func GetAllRequestHistorys() ([]dbmodels.RequestHistory, error) {
    session, collection := Connect(RequestHistorysCollectionName)
    defer session.Close()

    var requestsHistories []dbmodels.RequestHistory
    err := collection.Find(bson.M{}).All(&requestsHistories)

    return requestsHistories, err
}

func GetAllRequestHistorysLimited(limit int) ([]dbmodels.RequestHistory, error) {
    session, collection := Connect(RequestHistorysCollectionName)
    defer session.Close()

    var requestsHistories []dbmodels.RequestHistory
    err := collection.Find(bson.M{}).Limit(limit).All(&requestsHistories)

    return requestsHistories, err
}
