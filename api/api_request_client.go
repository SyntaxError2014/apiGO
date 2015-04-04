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

    vars := createApiVars(route, req, rw)

    if vars == nil {
        return
    }

    resp := &ApiResponse{}

    err = client.Call(handlerName, vars, resp)

    // Log requests and an eventual error messages
    log.Println(req.Method, route.Pattern, resp.StatusCode, resp.ErrorMessage)

    // Treat RPC call error separately
    if err != nil {
        GiveApiMessage(resp.StatusCode, err.Error(), rw)
        return
    }

    // If a content type is specified, write it to the response header
    if len(resp.ContentType) > 0 {
        rw.Header().Add("Content-type", resp.ContentType)
    }

    // Give appripriate http response depending on the status of the processed request
    if resp.ErrorMessage != "" {
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
func createApiVars(route *config.Route, req *http.Request, rw http.ResponseWriter) *ApiVar {
    err, statusCode := filter.CheckMethodAndParseContent(req)
    if err != nil {
        GiveApiMessage(statusCode, err.Error(), rw)
        return nil
    }

    body, err := convertBodyToReadableFormat(req.Body)
    if err != nil {
        GiveApiMessage(http.StatusBadRequest, err.Error(), rw)
        return nil
    }

    vars := &ApiVar{
        Route:                *route,
        RequestMethod:        req.Method,
        RequestHeader:        req.Header,
        RequestForm:          req.Form,
        RequestContentLength: req.ContentLength,
        RequestBody:          body,
        BasicAuth:            parseBasicAuthData(req),
    }

    return vars
}

// Converts the http request body from an io.ReadCloser to an array
// of bytes, which can then be serialized and deserialized
func convertBodyToReadableFormat(data io.ReadCloser) ([]byte, error) {
    body, err := ioutil.ReadAll(data)

    return body, err
}

// Prepare the Basic Authentication details into a data transport type
func parseBasicAuthData(req *http.Request) BasicAuthentication {
    username, password, ok := req.BasicAuth()

    basicAuth := BasicAuthentication{
        OK:       ok,
        Username: username,
        Password: password,
    }

    return basicAuth
}
