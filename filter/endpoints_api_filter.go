package filter

import (
    "apiGO/dbmodels"
    "apiGO/models"
)

// Checks if the user entity has all the mandatory
// fields populated with any kind of data
func CheckEndpointIntegrity(endpoint *models.Endpoint) bool {
    switch {
    case len(endpoint.Id) == 0:
        return false
    case len(endpoint.URLPath) == 0:
        return false
    case endpoint.User.Equal(dbmodels.User{}):
        return false
    }

    return true
}
