package canplaceflowers

import (
	"fmt"
	"testing"
)

type testCase struct {
	flowerbed []int
	n         int
	expected  bool
}

var testCases = []testCase{
	{
		flowerbed: []int{1, 0, 0, 0, 1},
		n:         1,
		expected:  true,
	},
	{
		flowerbed: []int{1, 0, 0, 0, 1},
		n:         2,
		expected:  false,
	},
}

func TestCanPlaceFlowers(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Flowerbed=%v, Flowers=%d", tc.flowerbed, tc.n), func(t *testing.T) {
			actual := canPlaceFlowers(tc.flowerbed, tc.n)
			if actual != tc.expected {
				t.Fatalf("canPlaceFlowers(%v, %d)=%v, expected %v", tc.flowerbed, tc.n, actual, tc.expected)
			}
		})
	}
}

func TestCanPlaceFlowers2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Flowerbed=%v, Flowers=%d", tc.flowerbed, tc.n), func(t *testing.T) {
			actual := canPlaceFlowers2(tc.flowerbed, tc.n)
			if actual != tc.expected {
				t.Fatalf("canPlaceFlowers(%v, %d)=%v, expected %v", tc.flowerbed, tc.n, actual, tc.expected)
			}
		})
	}
}

func BenchmarkCanPlaceFlowers(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i <= b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("Flowerbed=%v, Flowers=%d", tc.flowerbed, tc.n), func(b *testing.B) {
				actual := canPlaceFlowers(tc.flowerbed, tc.n)
				if actual != tc.expected {
					b.Fatalf("canPlaceFlowers(%v, %d)=%v, expected %v", tc.flowerbed, tc.n, actual, tc.expected)
				}
			})
		}
	}
}

func BenchmarkCanPlaceFlowers2(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i <= b.N; i++ {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("Flowerbed=%v, Flowers=%d", tc.flowerbed, tc.n), func(b *testing.B) {
				actual := canPlaceFlowers2(tc.flowerbed, tc.n)
				if actual != tc.expected {
					b.Fatalf("canPlaceFlowers2(%v, %d)=%v, expected %v", tc.flowerbed, tc.n, actual, tc.expected)
				}
			})
		}
	}
}
