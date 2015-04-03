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

func (api *Api) GetUserSession(vars *ApiVar, resp *ApiResponse) error {
    userSessionId, err, found := filter.GetIdFromParams(vars.RequestForm)
    if found {
        if err != nil {
            return badRequest(resp, err.Error())
        }

        return getUserSession(vars, resp, userSessionId)
    }

    limit, err, found := filter.GetIntValueFromParams("limit", vars.RequestForm)
    if found {
        if err != nil {
            return badRequest(resp, err.Error())
        }

        return getAllUserSessions(vars, resp, limit)
    }

    return getAllUserSessions(vars, resp, -1)

}

func (api *Api) PostUserSession(vars *ApiVar, resp *ApiResponse) error {
    expandedUserSession := &models.UserSession{}

    err := expandedUserSession.DeserializeJson(vars.RequestBody)
    if err != nil {
        return badRequest(resp, "The entity was not in the correct format")
    }

    if !filter.CheckUserSessionIntegrity(expandedUserSession) {
        return badRequest(resp, "The entity doesn't comply to the integrity requirements")
    }

    userSession, err := expandedUserSession.Collapse()
    if err != nil {
        return internalServerError(resp, err.Error())
    }

    userSession, err = service.CreateUserSession(userSession)
    if err != nil || userSession == nil {
        return internalServerError(resp, "The entity could not be processed")
    }

    resp.StatusCode = http.StatusCreated
    resp.Message, _ = userSession.SerializeJson()

    return nil
}

func (api *Api) PutUserSession(vars *ApiVar, resp *ApiResponse) error {
    expandedUserSession := &models.UserSession{}

    err := expandedUserSession.DeserializeJson(vars.RequestBody)
    if err != nil {
        return badRequest(resp, "The entity was not in the correct format")
    }

    if expandedUserSession.Id == "" {
        return badRequest(resp, "No id was specified for the userSession to be updated")
    }

    if !filter.CheckUserSessionIntegrity(expandedUserSession) {
        return badRequest(resp, "The entity doesn't comply to the integrity requirements")
    }

    userSession, err := expandedUserSession.Collapse()
    if err != nil {
        return internalServerError(resp, err.Error())
    }

    err = service.UpdateUserSession(userSession)
    if err != nil {
        return notFound(resp, "The userSession with the specified id could not be found")
    }

    resp.StatusCode = http.StatusOK
    resp.Message, _ = userSession.SerializeJson()

    return nil
}

func (api *Api) DeleteUserSession(vars *ApiVar, resp *ApiResponse) error {
    userSessionId, err, found := filter.GetIdFromParams(vars.RequestForm)

    if found {
        if err != nil {
            return badRequest(resp, err.Error())
        }

        err = service.DeleteUserSession(userSessionId)
        if err != nil {
            return notFound(resp, err.Error())
        }

        resp.StatusCode = http.StatusOK
        return nil
    }

    return badRequest(resp, err.Error())
}

func getAllUserSessions(vars *ApiVar, resp *ApiResponse, limit int) error {
    var userSessions []dbmodels.UserSession
    var err error

    if limit == 0 {
        return badRequest(resp, "The limit cannot be 0. Use the value -1 for retrieving all the entities")
    }

    if limit != -1 {
        userSessions, err = service.GetAllUserSessionsLimited(limit)
    } else {
        userSessions, err = service.GetAllUserSessions()
    }

    if err != nil {
        return internalServerError(resp, err.Error())
    }

    var expandedUserSessions []models.UserSession
    expandedUserSessions = make([]models.UserSession, len(userSessions))

    for i := 0; i < len(userSessions); i++ {
        err = expandedUserSessions[i].Expand(userSessions[i])

        if err != nil {
            return internalServerError(resp, err.Error())
        }
    }

    userSessionsJson, err := json.MarshalIndent(expandedUserSessions, interfaces.JsonPrefix, interfaces.JsonIndent)

    if err != nil {
        return internalServerError(resp, err.Error())
    }

    resp.StatusCode = http.StatusOK
    resp.Message = userSessionsJson

    return nil
}

func getUserSession(vars *ApiVar, resp *ApiResponse, userSessionId bson.ObjectId) error {
    userSession, err := service.GetUserSession(userSessionId)

    if err != nil {
        if err.Error() == "not found" {
            return notFound(resp, err.Error())
        } else {
            return internalServerError(resp, err.Error())
        }
    }

    if userSession == nil {
        return notFound(resp, "No UserSession with the selected id was found")
    }

    var expandedUserSession models.UserSession
    err = expandedUserSession.Expand(*userSession)

    if err != nil {
        return internalServerError(resp, err.Error())
    }

    userSessionJson, err := expandedUserSession.SerializeJson()
    if err != nil {
        return internalServerError(resp, err.Error())
    }

    resp.StatusCode = http.StatusOK
    resp.Message = userSessionJson

    return nil
}
