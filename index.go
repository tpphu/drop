package drop

// Load returns a new Value given a reporter used to report failures
// and value to be inspected.
func Load(source interface{}) *Value {
	return &Value{
		source: source,
	}
}
