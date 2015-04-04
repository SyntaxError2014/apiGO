package api

import (
    "apiGO/config"
    "apiGO/filter"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "net/rpc"
)

// Performs a call to the RPC server using the appropriate data extracted from the HTTP request,
// which is stored in the specific API variable
func PerformClientCall(handlerName string, rw http.ResponseWriter, req *http.Request, route *config.Route) {
    client, err := rpc.DialHTTP(config.Protocol, config.RpcServerAddress)
    defer closeClient(client)

    if err != nil {
        GiveApiMessage(http.StatusInternalServerError, err.Error(), rw)
        return
    }

    vars := createApiVars(route.Pattern, req, rw)

    if vars == nil {
        return
    }

    resp := &ApiResponse{}

    err = client.Call(handlerName, vars, resp)

    log.Println(req.Method, route.Pattern, resp.StatusCode)
    if err != nil {
        GiveApiMessage(resp.StatusCode, err.Error(), rw)
    } else if resp.ErrorMessage != "" {
        GiveApiMessage(resp.StatusCode, resp.ErrorMessage, rw)
    } else {
        GiveApiResponse(resp.StatusCode, resp.Message, rw)
    }
}

// Close an opened RPC client
func closeClient(client *rpc.Client) {
    err := client.Close()
    if err != nil {
        log.Panic("Error in closing client connection: ", err.Error())
    }
}

// Create the specialized API variable filled with data that is
// extracted from the HTTP request made to the server
func createApiVars(requestURI string, r *http.Request, rw http.ResponseWriter) *ApiVar {
    err, statusCode := filter.CheckMethodAndParseContent(r)
    if err != nil {
        GiveApiMessage(statusCode, err.Error(), rw)
        return nil
    }

    body, err := convertBodyToReadableFormat(r.Body)
    if err != nil {
        GiveApiMessage(http.StatusBadRequest, err.Error(), rw)
        return nil
    }

    vars := &ApiVar{
        RequestURI:           requestURI,
        RequestHeader:        r.Header,
        RequestForm:          r.Form,
        RequestContentLength: r.ContentLength,
        RequestBody:          body,
    }

    return vars
}

// Converts the http request body from an io.ReadCloser to an array
// of bytes, which can then be serialized and deserialized
func convertBodyToReadableFormat(data io.ReadCloser) ([]byte, error) {
    body, err := ioutil.ReadAll(data)

    return body, err
}
