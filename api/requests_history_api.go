package api

import (
    "apiGO/dbmodels"
    "apiGO/filter"
    "apiGO/interfaces"
    "apiGO/models"
    "apiGO/service"
    "encoding/json"
    "net/http"
)

func (api *Api) GetRequestsHistory(vars *ApiVar, resp *ApiResponse) error {
    user, err := fetchUserUsingAuthToken(vars, resp)
    if user == nil {
        return nil
    }

    endpoint, err := fetchEndpointUsingEndpointPath(vars, resp)
    if endpoint == nil {
        return nil
    }

    requestsHistory, err := service.GetEntireRequestHistoryForEndpoint(endpoint.Id)
    if err != nil {
        internalServerError(resp, err.Error())
    }

    historyArray := expandRequestsHistoryArray(requestsHistory)

    historyJson, err := json.MarshalIndent(historyArray, interfaces.JsonPrefix, interfaces.JsonIndent)
    if err != nil {
        return internalServerError(resp, err.Error())
    }

    resp.StatusCode = http.StatusOK
    resp.Message = historyJson

    return nil
}

func fetchEndpointUsingEndpointPath(vars *ApiVar, resp *ApiResponse) (*dbmodels.Endpoint, error) {
    endpointPath, endpointPathError, endpointPathFound := filter.GetStringValueFromParams("endpointPath", vars.RequestForm)

    if !endpointPathFound {
        return nil, badRequest(resp, "No endpoint id was specified")
    }

    if endpointPathError != nil {
        return nil, badRequest(resp, endpointPathError.Error())
    }

    return service.GetEndpointByURLPath(endpointPath)
}

func expandRequestsHistoryArray(reqHistArr []dbmodels.RequestHistory) []models.RequestHistory {
    var expandedArray []models.RequestHistory
    expandedArray = make([]models.RequestHistory, len(reqHistArr))

    for i := 0; i < len(reqHistArr); i++ {
        expandedArray[i].Expand(reqHistArr[i])
    }

    return expandedArray
}
