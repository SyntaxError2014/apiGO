package python_integration

import (
    "net/url"
    "os/exec"
    "strings"
)

const scriptFile = "script.py"

func ExecuteCommand(codeBody string, parameters url.Values) ([]byte, error) {
    head := "script.py"
    params := make([]string, len(parameters))

    params = append(params, codeBody)

    for _, p := range parameters {
        params = append(params, strings.Join(p, " "))
    }

    return exec.Command(head, params...).Output()
}
