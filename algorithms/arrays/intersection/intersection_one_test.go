package intersection

import (
	"reflect"
	"testing"
)

func TestIntersectionOfTwoArrays(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			name:     "should return [2] for nums1 = [1, 2, 2, 1] and nums2 = [2,2]",
			nums1:    []int{1, 2, 2, 1},
			nums2:    []int{2, 2},
			expected: []int{2},
		},
		{
			name:     "should return [9,4] for nums1 = [4,9,5] and nums2 = [9, 4, 9, 8, 4]",
			nums1:    []int{4, 9, 5},
			nums2:    []int{9, 4, 9, 8, 4},
			expected: []int{9, 4},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := intersection(test.nums1, test.nums2)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf(
					"intersection(%v, %v) = %v, want %v",
					test.nums1,
					test.nums2,
					actual,
					test.expected,
				)
			}
		})
	}
}
