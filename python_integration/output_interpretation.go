package python_integration

import (
    "strconv"
    "strings"
)

func ParseOutput(output []byte) (int, string) {
    stringOutput := string(output)

    data := strings.Split(stringOutput, "\n")
    statusCode, _ := strconv.Atoi(data[0])

    return statusCode, data[1]
}
