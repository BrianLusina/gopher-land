package cpucaches

import "testing"

func BenchmarkSumLinkedList(b *testing.B) {
	linkedList := node{
		value: 1,
		next: &node{
			value: 2,
			next: &node{
				value: 3,
				next: &node{
					value: 4,
					next: &node{
						value: 5,
						next: &node{
							value: 6,
							next: &node{
								value: 7,
								next: &node{
									value: 8,
								},
							},
						},
					},
				},
			},
		},
	}

	for i := 0; i < b.N; i++ {
		sumLinkedList(&linkedList)
	}
}

func BenchmarkSum2_with_linked(b *testing.B) {
	s := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		sum2(s)
	}
}
