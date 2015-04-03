package service

import (
    "apiGO/dbmodels"
    "gopkg.in/mgo.v2/bson"
)

const EndpointsCollectionName = "endpoints"

func CreateEndpoint(endpoint *dbmodels.Endpoint) (*dbmodels.Endpoint, error) {
    session, collection := Connect(EndpointsCollectionName)
    defer session.Close()

    if endpoint.Id == "" {
        endpoint.Id = bson.NewObjectId()
    }

    err := collection.Insert(endpoint)

    return endpoint, err
}

func UpdateEndpoint(endpoint *dbmodels.Endpoint) error {
    session, collection := Connect(EndpointsCollectionName)
    defer session.Close()

    err := collection.UpdateId(endpoint.Id, endpoint)

    return err
}

func DeleteEndpoint(endpointId bson.ObjectId) error {
    session, collection := Connect(EndpointsCollectionName)
    defer session.Close()

    err := collection.RemoveId(endpointId)

    return err
}

func GetEndpoint(endpointId bson.ObjectId) (*dbmodels.Endpoint, error) {
    session, collection := Connect(EndpointsCollectionName)
    defer session.Close()

    endpoint := dbmodels.Endpoint{}
    err := collection.FindId(endpointId).One(&endpoint)

    return &endpoint, err
}

func GetAllEndpoints() ([]dbmodels.Endpoint, error) {
    session, collection := Connect(EndpointsCollectionName)
    defer session.Close()

    var endpoints []dbmodels.Endpoint
    err := collection.Find(bson.M{}).All(&endpoints)

    return endpoints, err
}

func GetAllEndpointsLimited(limit int) ([]dbmodels.Endpoint, error) {
    session, collection := Connect(EndpointsCollectionName)
    defer session.Close()

    var endpoints []dbmodels.Endpoint
    err := collection.Find(bson.M{}).Limit(limit).All(&endpoints)

    return endpoints, err
}
