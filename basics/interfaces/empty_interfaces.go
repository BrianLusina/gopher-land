// interfaces This showcases empty interfaces whereby an Empty Interface can be of any type & this shows how we can
// check which type it is before use. In this case, we are checking if the Empty Interface is of type string
package interfaces

var EmptyInterface interface{}

func Empty() interface{} {
	EmptyInterface = "こんにちは"

	// check if EmptyInterface is of type string.
	if s, ok := EmptyInterface.(string); ok {
		return s
	}

	return "Not a string"
}
