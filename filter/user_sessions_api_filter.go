package filter

import (
    "apiGO/dbmodels"
    "apiGO/models"
    "apiGO/service"
)

// Checks if the user entity has all the mandatory
// fields populated with any kind of data
func CheckUserSessionIntegrity(userSession *models.UserSession) bool {
    switch {
    case len(userSession.Id) == 0:
        return false
    case len(userSession.Token) == 0:
        return false
    case userSession.User.Equal(dbmodels.User{}):
        return false
    }

    return true
}

// Checks if the session exists and if the user
// that is allocated for that session also exists
func CheckAuthToken(token string) bool {
    userSession, err := service.GetUserSessionByToken(token)

    if err != nil || userSession == nil {
        return false
    }

    user, err := service.GetUser(userSession.UserId)

    if err != nil || user == nil {
        return false
    }

    return true
}
