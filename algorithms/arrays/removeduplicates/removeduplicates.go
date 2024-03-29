package removeduplicates

/*
*
RemoveDuplicates removes duplicates from a sorted slice of integers.

Removes duplicates from a sorted list of integers and returns the number of the unique elements in the list. This
modifies the list in place so no extra space is used, effectively the Space complexity is O(1). However, the time
complexity is O(n) where n is the length of the list.

This uses 2 pointers to keep track of the unique and non-unique elements in the list of integers. i pointer is at
the element before and j pointer is at the next element. These are moved along the elements in the list as below:

1. as long as j is less than the length of the list we shall iterate until j reaches the end of the list, i.e. the
last element in the list

2. if the element at position i is not equal to the element at position j(at this point js is ahead of i), we
increase the i index by 1 and assign the index at i the value of the index at j and increment the index at j, i.e.
moving the pointer j and i forward by 1

3. if the elements are equal, then that means we have encountered a duplicate & therefore we move the pointer j
forward by 1 until we encounter a unique element.

4. When the loop is complete we return i +1. This is because the pointer i is now pointing to the last unique
element in the list.

The other elements after the pointer i are the duplicate elements. This modifies the sorted list in place retaining
"some" of the duplicated elements to the right of i. Note that this means we will not know how many times an
element occurred in the original sorted list of elements.
*
*/
func RemoveDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	slowPointer := 0
	fastPointer := 1

	for fastPointer < len(nums) {
		if nums[slowPointer] != nums[fastPointer] {
			slowPointer += 1
			nums[slowPointer] = nums[fastPointer]
		}
		fastPointer++
	}

	return slowPointer + 1
}

func removeDuplicatesTwo(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	// Initialize an integer k that updates the kth index of the array...
	// only when the current element does not match either of the two previous indexes...

	k := 2

	for i := 2; i < len(nums); i++ {
		// If the index does not match the (k-1)th and (k-2)th elements, count that element...
		if nums[i] != nums[k-2] || nums[i] != nums[k-1] {
			nums[k] = nums[i]
			k++
		}
		// If the index matches the (k-1)th and (k-2)th elements, we skip it...
	}

	// Return k after placing the final result in the first k slots of nums...
	return k
}
