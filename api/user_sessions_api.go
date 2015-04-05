package api

import (
    "apiGO/dbmodels"
    "apiGO/filter"
    "apiGO/models"
    "apiGO/service"
    "net/http"
)

func (api *Api) GetUserSession(vars *ApiVar, resp *ApiResponse) error {
    user, err := fetchUserUsingAuthToken(vars, resp)

    if err != nil || user == nil {
        return nil
    }

    expandedUser := &models.User{}
    expandedUser.Expand(*user)
    userJson, _ := expandedUser.SerializeJson()

    resp.StatusCode = http.StatusOK
    resp.Message = userJson

    return nil
}

func (api *Api) PostUserSession(vars *ApiVar, resp *ApiResponse) error {
    // Get URL parameters
    username, userError, userWasFound := filter.GetStringValueFromParams("username", vars.RequestForm)
    password, passwordError, passwordWasFound := filter.GetStringValueFromParams("password", vars.RequestForm)

    if !userWasFound || !passwordWasFound {
        return badRequest(resp, "The username or password was not specified")
    }

    if userError != nil {
        return badRequest(resp, userError.Error())
    }
    if passwordError != nil {
        return badRequest(resp, passwordError.Error())
    }

    // Fetch User entity from database
    user, err := service.GetUserByUsernameAndPassword(username, password)
    if err != nil || user == nil {
        return badRequest(resp, "Username or password is incorrect")
    }

    // Delete all the existing sessions
    service.DeleteAllSessionsWithUserId(user.Id)

    // Generate a new user session
    userSession, err := service.GenerateAndInsertUserSession(user.Id)
    if err != nil || userSession == nil {
        return internalServerError(resp, err.Error())
    }

    resp.StatusCode = http.StatusCreated
    resp.Message = []byte(userSession.Token)

    return nil
}

func (api *Api) DeleteUserSession(vars *ApiVar, resp *ApiResponse) error {
    return nil
}

func fetchUserUsingAuthToken(vars *ApiVar, resp *ApiResponse) (*dbmodels.User, error) {
    token, err, found := filter.GetStringValueFromParams("token", vars.RequestForm)

    if !found {
        return nil, badRequest(resp, "Session token has not been specified")
    }

    if err != nil {
        return nil, badRequest(resp, err.Error())
    }

    userSession, err := service.GetUserSessionByToken(token)
    if err != nil || userSession == nil {
        return nil, notFound(resp, "There is no session with the specified token")
    }

    user, err := service.GetUser(userSession.UserId)
    if err != nil || user == nil {
        service.DeleteUserSession(userSession.Id)
        return nil, notFound(resp, "The user with the current session no longer exists")
    }

    return user, nil
}
