package filter

import (
    "apiGO/models"
)

// Checks if the user entity has all the mandatory
// fields populated with any kind of data
func CheckUserIntegrity(user *models.User) bool {
    switch {
    case len(user.Id) == 0:
        return false
    }

    return true
}
