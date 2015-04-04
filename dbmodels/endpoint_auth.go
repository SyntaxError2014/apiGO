package dbmodels

import ()

type EndpointAuth struct {
    Username string `bson:"username" json:"username"`
    Password string `bson:"password" json:"password"`
}

func (endpointAuth EndpointAuth) Equal(otherEndpointAuth EndpointAuth) bool {
    switch {
    case endpointAuth.Username != otherEndpointAuth.Username:
        return false
    case endpointAuth.Password != otherEndpointAuth.Password:
        return false
    }

    return true
}
