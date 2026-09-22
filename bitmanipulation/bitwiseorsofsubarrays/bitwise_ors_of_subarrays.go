package bitwiseorsofsubarrays

func subarrayBitwiseORs(arr []int) int {
	// Map to store all distinct bitwise OR values from all subarrays
	all_ors := make(map[int]bool)
	// Map to store the bitwise OR values from the previous iteration
	previous_ors := make(map[int]bool)

	// Iterate through each element in the array
	for i := range arr {
		current_element := arr[i]
		// Initialize a new map to store the current OR values for this iteration
		current_ors := make(map[int]bool)

		// Add the current element to the current OR values
		current_ors[current_element] = true

		// Calculate OR with all subarrays ending at previous position
		for previous_values := range previous_ors {
			current_ors[previous_values|current_element] = true
		}

		// Update the map for the next iteration
		previous_ors = current_ors

		// Add all current OR values to the overall set of OR values
		for or_val := range current_ors {
			all_ors[or_val] = true
		}
	}
	return len(all_ors)
}
