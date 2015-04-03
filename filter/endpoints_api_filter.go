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
    case endpoint.GET.Equal(dbmodels.EndpointResponse{}) &&
        endpoint.POST.Equal(dbmodels.EndpointResponse{}) &&
        endpoint.PUT.Equal(dbmodels.EndpointResponse{}) &&
        endpoint.DELETE.Equal(dbmodels.EndpointResponse{}):
        return false
    }

    return true
}
