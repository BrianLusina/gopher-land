package mergesortedarray

/*
	merge uses a 3 pointer approach to merge sorted arrays.

The first pointer i points to the last element of the subarray of size m in the nums1 array,
The second pointer j points to the last element of the nums2 array of size n.
The pointer k points to the end of the array nums1 of size m + n.
Until the pointer j reaches 0 (that is, until we reach the beginning of the nums2 array), we check the following conditions:
- If nums1[i] <= nums2[j] => write nums2[j] to nums1[k] and then decrease the value of j
- In all other cases, write nums1[i] to nums1[k] and then decrease the value of i

Complexity:
- Time complexity: O(n+m)
- Space complexity: O(1)
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1

	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] <= nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
	}

	if j >= 0 {
		copy(nums1[:k+1], nums2[:j+1])
	}
}
