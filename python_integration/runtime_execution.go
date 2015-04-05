package python_integration

import (
    "apiGO/interfaces"
    "encoding/json"
    "log"
    "net/url"
    "os/exec"
)

const scriptHead = "python"
const scriptFile = "script.py"

func ExecuteCommand(codeBody string, parameters url.Values) ([]byte, error) {
    params := make([]string, len(parameters))

    params = append(params, scriptFile, codeBody)
    log.Println(scriptHead, params)

    jsonArray, _ := json.MarshalIndent(parameters, interfaces.JsonPrefix, interfaces.JsonIndent)

    return exec.Command(scriptHead, scriptFile, codeBody, string(jsonArray)).Output()
}
