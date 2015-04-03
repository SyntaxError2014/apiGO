package interfaces

// Interface for defining the models as Objects which
// can be compared with other objects.
//
// This interface is mainly used in order to interpret
// unit test results
type Object interface {
    Equal(obj Object) bool
}
