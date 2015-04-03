package filter

import (
    "errors"
    "net/http"
)

// Function for filtering HTTP requests.
// This checks if POSTs and PUTs have any kind of content (contentLength != 0).
// If yes, then it tries parsing Form/PostForm in order for the Form, MultipartForm and
// Body to be readable. If this fails, an error is returned
//
// This returns and error and a status code.
// A nil error and -1 status code means that the parsing went just fine.
func CheckMethodAndParseContent(r *http.Request) (error, int) {
    if r.ContentLength == 0 {
        if r.Method == "POST" || r.Method == "PUT" {
            return errors.New("No content has been received"), http.StatusBadRequest
        }
    }

    err := r.ParseForm()

    if err != nil {
        return errors.New("The request form has an invalid format"), http.StatusBadRequest
    }

    return nil, -1
}
