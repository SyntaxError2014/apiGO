package interfaces

// Object that can be serialized and deserialized using JSON data type
type Serializable interface {
    SerializeJson() ([]byte, error)
    DeserializeJson(obj []byte) error
}

// Constants used for serializations
const (
    JsonPrefix = ""
    JsonIndent = "  "
)
