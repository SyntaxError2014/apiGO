package api

import (
    "apiGO/dbmodels"
    "apiGO/filter"
    "apiGO/interfaces"
    "apiGO/models"
    "apiGO/service"
    "encoding/json"
    "gopkg.in/mgo.v2/bson"
    "net/http"
)

func (api *Api) GetUser(vars *ApiVar, resp *ApiResponse) error {
    userId, err, found := filter.GetIdFromParams(vars.RequestForm)
    if found {
        if err != nil {
            return badRequest(resp, err.Error())
        }

        return getUser(vars, resp, userId)
    }

    limit, err, found := filter.GetIntValueFromParams("limit", vars.RequestForm)
    if found {
        if err != nil {
            return badRequest(resp, err.Error())
        }

        return getAllUsers(vars, resp, limit)
    }

    return getAllUsers(vars, resp, -1)

}

func (api *Api) PostUser(vars *ApiVar, resp *ApiResponse) error {
    expandedUser := &models.User{}

    err := expandedUser.DeserializeJson(vars.RequestBody)
    if err != nil {
        return badRequest(resp, "The entity was not in the correct format")
    }

    if !filter.CheckUserIntegrity(expandedUser) {
        return badRequest(resp, "The entity doesn't comply to the integrity requirements")
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

func (api *Api) PutUser(vars *ApiVar, resp *ApiResponse) error {
    expandedUser := &models.User{}

    err := expandedUser.DeserializeJson(vars.RequestBody)
    if err != nil {
        return badRequest(resp, "The entity was not in the correct format")
    }

    if expandedUser.Id == "" {
        return badRequest(resp, "No id was specified for the user to be updated")
    }

    if !filter.CheckUserIntegrity(expandedUser) {
        return badRequest(resp, "The entity doesn't comply to the integrity requirements")
    }

    user, err := expandedUser.Collapse()
    if err != nil {
        return internalServerError(resp, err.Error())
    }

    err = service.UpdateUser(user)
    if err != nil {
        return notFound(resp, "The user with the specified id could not be found")
    }

    resp.StatusCode = http.StatusOK
    resp.Message, _ = user.SerializeJson()

    return nil
}

func (api *Api) DeleteUser(vars *ApiVar, resp *ApiResponse) error {
    userId, err, found := filter.GetIdFromParams(vars.RequestForm)

    if found {
        if err != nil {
            return badRequest(resp, err.Error())
        }

        err = service.DeleteUser(userId)
        if err != nil {
            return notFound(resp, err.Error())
        }

        resp.StatusCode = http.StatusOK
        return nil
    }

    return badRequest(resp, err.Error())
}

func getAllUsers(vars *ApiVar, resp *ApiResponse, limit int) error {
    var users []dbmodels.User
    var err error

    if limit == 0 {
        return badRequest(resp, "The limit cannot be 0. Use the value -1 for retrieving all the entities")
    }

    if limit != -1 {
        users, err = service.GetAllUsersLimited(limit)
    } else {
        users, err = service.GetAllUsers()
    }

    if err != nil {
        return internalServerError(resp, err.Error())
    }

    var expandedUsers []models.User
    expandedUsers = make([]models.User, len(users))

    for i := 0; i < len(users); i++ {
        err = expandedUsers[i].Expand(users[i])

        if err != nil {
            return internalServerError(resp, err.Error())
        }
    }

    usersJson, err := json.MarshalIndent(expandedUsers, interfaces.JsonPrefix, interfaces.JsonIndent)

    if err != nil {
        return internalServerError(resp, err.Error())
    }

    resp.StatusCode = http.StatusOK
    resp.Message = usersJson

    return nil
}

func getUser(vars *ApiVar, resp *ApiResponse, userId bson.ObjectId) error {
    user, err := service.GetUser(userId)

    if err != nil {
        if err.Error() == "not found" {
            return notFound(resp, err.Error())
        } else {
            return internalServerError(resp, err.Error())
        }
    }

    if user == nil {
        return notFound(resp, "No User with the selected id was found")
    }

    var expandedUser models.User
    err = expandedUser.Expand(*user)

    if err != nil {
        return internalServerError(resp, err.Error())
    }

    userJson, err := expandedUser.SerializeJson()
    if err != nil {
        return internalServerError(resp, err.Error())
    }

    resp.StatusCode = http.StatusOK
    resp.Message = userJson

    return nil
}
