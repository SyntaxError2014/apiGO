package interfaces

type Expandable interface {
    Expand(obj Object) error
    Collapse() (*Object, error)
}
