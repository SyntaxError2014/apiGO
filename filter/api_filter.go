package filter

// Checks whether the array of bytes parameter is either
// empty or has the string value "null" stored.
func CheckNotNull(data []byte) bool {
    if data == nil {
        return false
    }

    if string(data) == "null" {
        return false
    }

    return true
}
