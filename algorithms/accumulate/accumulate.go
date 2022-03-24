package accumulate

// Accumulate takes in a collection & applies function to each item in collection
func Accumulate(collection []string, function func(string) string) (output []string) {
	for _, item := range collection {
		output = append(output, function(item))
	}
	return output
}
