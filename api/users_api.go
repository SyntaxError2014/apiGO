package api

import (
    "apiGO/filter"
    "apiGO/models"
    "apiGO/service"
    "net/http"
)

func (api *Api) PostUser(vars *ApiVar, resp *ApiResponse) error {
    expandedUser := &models.User{}

    err := expandedUser.DeserializeJson(vars.RequestBody)
    if err != nil {
        return badRequest(resp, "The entity was not in the correct format")
    }

    if !filter.CheckUserIntegrity(expandedUser) {
        return badRequest(resp, "The entity doesn't comply to the integrity requirements")
    }

    if filter.CheckUserExists(expandedUser) {
        return badRequest(resp, "The user already exists")
    }

    user, err := expandedUser.Collapse()
    if err != nil {
        return internalServerError(resp, err.Error())
    }

    user, err = service.CreateUser(user)
    if err != nil || user == nil {
        return internalServerError(resp, "The entity could not be processed")
    }

    resp.StatusCode = http.StatusCreated
    resp.Message, _ = user.SerializeJson()

    return nil
}
