package drop

// Value provides methods to inspect attached interface{} object
// (Go representation of arbitrary JSON value) and cast it to
// concrete type.
type Value struct {
	source interface{}
}

func (v Value) Object() *Object {
	return &Object{
		data: v.source.(map[string]interface{}),
	}
}
