package interfaces

// Interface used to expand JSON entites that contain different relations.
// In short, entity IDs are transformed in those entities in full
// inside the json objects that implement this interface.
//
// Each database model should have an expanded representation
// that is used for receiving/returning data from requests
type Expandable interface {
    Expand(obj Object) error
    Collapse() (*Object, error)
}
