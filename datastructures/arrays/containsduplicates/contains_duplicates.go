// containsduplicates package provides a function to check if an array contains duplicate elements.
package containsduplicates

// ContainsDuplicate returns true if an array contains duplicate elements.
// Time complexity: O(n) and space complexity: O(n) where n is the length of the array.
func ContainsDuplicate(nums []int) bool {
	// using a struct to store the number and its frequency as struct occupies 0 bytes while a bool occupies 1 byte
	// setting the size/capacity of the map to be the length of the array avoids map resizing
	set := make(map[int]struct{}, len(nums))

	for _, num := range nums {
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}
