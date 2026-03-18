package mapsum

import (
	"gopherland/pkg/types"
	"testing"
)

type testCase struct {
	name       string
	operations []types.Pair[string, types.Pair[string, int]]
}

var testCases = []testCase{
	{
		name: "MapSum()->Insert('apple', 3)->Sum('apple')->Insert('apple', 2)->Sum('ap')",
		operations: []types.Pair[string, types.Pair[string, int]]{
			{
				First: "insert",
				Second: types.Pair[string, int]{
					First:  "apple",
					Second: 3,
				},
			},
			{
				First: "sum",
				Second: types.Pair[string, int]{
					First:  "apple",
					Second: 3,
				},
			},
			{
				First: "insert",
				Second: types.Pair[string, int]{
					First:  "apple",
					Second: 2,
				},
			},
			{
				First: "sum",
				Second: types.Pair[string, int]{
					First:  "ap",
					Second: 2,
				},
			},
		},
	},
	{
		name: "MapSum()->Insert('apple', 3)->Sum('ap')->Insert('ap', 2)->Sum('ap')",
		operations: []types.Pair[string, types.Pair[string, int]]{
			{
				First: "insert",
				Second: types.Pair[string, int]{
					First:  "apple",
					Second: 3,
				},
			},
			{
				First: "sum",
				Second: types.Pair[string, int]{
					First:  "ap",
					Second: 3,
				},
			},
			{
				First: "insert",
				Second: types.Pair[string, int]{
					First:  "ap",
					Second: 2,
				},
			},
			{
				First: "sum",
				Second: types.Pair[string, int]{
					First:  "ap",
					Second: 5,
				},
			},
		},
	},
	{
		name: "MapSum()->Insert('apple', 3)->Insert('apple', 5)->Sum('ap')->Insert('ap', 2)->Sum('ap')",
		operations: []types.Pair[string, types.Pair[string, int]]{
			{
				First: "insert",
				Second: types.Pair[string, int]{
					First:  "apple",
					Second: 3,
				},
			},
			{
				First: "insert",
				Second: types.Pair[string, int]{
					First:  "apple",
					Second: 5,
				},
			},
			{
				First: "sum",
				Second: types.Pair[string, int]{
					First:  "ap",
					Second: 5,
				},
			},
			{
				First: "insert",
				Second: types.Pair[string, int]{
					First:  "apricot",
					Second: 2,
				},
			},
			{
				First: "sum",
				Second: types.Pair[string, int]{
					First:  "ap",
					Second: 7,
				},
			},
		},
	},
	{
		name: "MapSum()->Insert('car', 3)->Insert('cat', 2)->Insert('cart', 4)->Sum('ca')->Sum('car')",
		operations: []types.Pair[string, types.Pair[string, int]]{
			{
				First: "insert",
				Second: types.Pair[string, int]{
					First:  "car",
					Second: 3,
				},
			},
			{
				First: "insert",
				Second: types.Pair[string, int]{
					First:  "cat",
					Second: 2,
				},
			},
			{
				First: "insert",
				Second: types.Pair[string, int]{
					First:  "cart",
					Second: 4,
				},
			},
			{
				First: "sum",
				Second: types.Pair[string, int]{
					First:  "ca",
					Second: 9,
				},
			},
			{
				First: "sum",
				Second: types.Pair[string, int]{
					First:  "car",
					Second: 7,
				},
			},
		},
	},
	{
		name: "MapSum()->Insert('dog', 5)->Insert('cat', 7)->Sum('z')",
		operations: []types.Pair[string, types.Pair[string, int]]{
			{
				First: "insert",
				Second: types.Pair[string, int]{
					First:  "dog",
					Second: 5,
				},
			},
			{
				First: "insert",
				Second: types.Pair[string, int]{
					First:  "cat",
					Second: 7,
				},
			},
			{
				First: "sum",
				Second: types.Pair[string, int]{
					First:  "z",
					Second: 0,
				},
			},
		},
	},
	{
		name: "MapSum()->Insert('a', 3)->Insert('apple', 2)->Sum('a')->Sum('app')",
		operations: []types.Pair[string, types.Pair[string, int]]{
			{
				First: "insert",
				Second: types.Pair[string, int]{
					First:  "a",
					Second: 3,
				},
			},
			{
				First: "insert",
				Second: types.Pair[string, int]{
					First:  "apple",
					Second: 2,
				},
			},
			{
				First: "sum",
				Second: types.Pair[string, int]{
					First:  "a",
					Second: 5,
				},
			},
			{
				First: "sum",
				Second: types.Pair[string, int]{
					First:  "app",
					Second: 2,
				},
			},
		},
	},
}

func TestMapSumPairsBruteForce(t *testing.T) {
	for _, tc := range testCases {
		mapSum := NewMapSumBruteForce()
		t.Run(tc.name, func(t *testing.T) {
			for _, operation := range tc.operations {
				cmd := operation.First
				params := operation.Second

				if cmd == "insert" {
					key, value := params.First, params.Second
					mapSum.Insert(key, value)
				}
				if cmd == "sum" {
					prefix, expected := params.First, params.Second
					actual := mapSum.Sum(prefix)
					if actual != expected {
						t.Errorf("expected %d to equal %d", actual, expected)
					}
				}
			}
		})
	}
}

func BenchmarkMapSumPairsBruteForce(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}
	for b.Loop() {
		for _, tc := range testCases {
			mapSum := NewMapSumBruteForce()

			for _, operation := range tc.operations {
				cmd := operation.First
				params := operation.Second

				if cmd == "insert" {
					key, value := params.First, params.Second
					mapSum.Insert(key, value)
				}
				if cmd == "sum" {
					prefix, _ := params.First, params.Second
					_ = mapSum.Sum(prefix)

				}
			}
		}
	}
}

func TestMapSumPairsPrefix(t *testing.T) {
	for _, tc := range testCases {
		mapSum := NewMapSumPrefix()
		t.Run(tc.name, func(t *testing.T) {
			for _, operation := range tc.operations {
				cmd := operation.First
				params := operation.Second

				if cmd == "insert" {
					key, value := params.First, params.Second
					mapSum.Insert(key, value)
				}
				if cmd == "sum" {
					prefix, expected := params.First, params.Second
					actual := mapSum.Sum(prefix)
					if actual != expected {
						t.Errorf("expected %d to equal %d", actual, expected)
					}
				}
			}
		})
	}
}

func BenchmarkMapSumPairsPrefix(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}
	for b.Loop() {
		for _, tc := range testCases {
			mapSum := NewMapSumPrefix()

			for _, operation := range tc.operations {
				cmd := operation.First
				params := operation.Second

				if cmd == "insert" {
					key, value := params.First, params.Second
					mapSum.Insert(key, value)
				}
				if cmd == "sum" {
					prefix, _ := params.First, params.Second
					_ = mapSum.Sum(prefix)

				}
			}
		}
	}
}

func TestMapSumPairsTrie(t *testing.T) {
	for _, tc := range testCases {
		mapSum := NewMapSumTrie()
		t.Run(tc.name, func(t *testing.T) {
			for _, operation := range tc.operations {
				cmd := operation.First
				params := operation.Second

				if cmd == "insert" {
					key, value := params.First, params.Second
					mapSum.Insert(key, value)
				}
				if cmd == "sum" {
					prefix, expected := params.First, params.Second
					actual := mapSum.Sum(prefix)
					if actual != expected {
						t.Errorf("expected %d to equal %d", actual, expected)
					}
				}
			}
		})
	}
}

func BenchmarkMapSumPairsTrie(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}
	for b.Loop() {
		for _, tc := range testCases {
			mapSum := NewMapSumTrie()

			for _, operation := range tc.operations {
				cmd := operation.First
				params := operation.Second

				if cmd == "insert" {
					key, value := params.First, params.Second
					mapSum.Insert(key, value)
				}
				if cmd == "sum" {
					prefix, _ := params.First, params.Second
					_ = mapSum.Sum(prefix)

				}
			}
		}
	}
}
