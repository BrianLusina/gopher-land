package meetingrooms

import (
	"fmt"
	"testing"
)

type minMeetingRoomsTestCase struct {
	intervals [][]int
	expected  int
}

var minMeetingRoomsTestCases = []minMeetingRoomsTestCase{
	{
		intervals: [][]int{},
		expected:  0,
	},
	{
		intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
		expected:  2,
	},
	{
		intervals: [][]int{{7, 10}, {2, 4}},
		expected:  1,
	},
	{
		intervals: [][]int{{1, 5}, {8, 9}, {8, 9}},
		expected:  2,
	},
	{
		intervals: [][]int{{1, 18}, {18, 23}, {15, 29}, {4, 15}, {2, 11}, {5, 13}},
		expected:  4,
	},
	{
		intervals: [][]int{{2, 8}, {3, 4}, {3, 9}, {5, 11}, {8, 20}, {11, 15}},
		expected:  3,
	},
	{
		intervals: [][]int{{1, 7}, {2, 6}, {3, 7}, {4, 8}, {5, 8}, {2, 9}, {1, 8}},
		expected:  7,
	},
	{
		intervals: [][]int{{1, 6}, {4, 8}, {1, 5}, {6, 8}, {8, 11}, {8, 9}, {5, 10}},
		expected:  3,
	},
	{
		intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {9, 15}, {12, 14}},
		expected:  2,
	},
	{
		intervals: [][]int{{1, 2}, {4, 6}, {3, 4}, {7, 8}},
		expected:  1,
	},
	{
		intervals: [][]int{{1, 7}, {2, 6}, {3, 7}, {4, 8}, {5, 8}},
		expected:  5,
	},
	{
		intervals: [][]int{{1, 2}, {1, 2}, {1, 2}},
		expected:  3,
	},
}

func TestMinMeetingRooms(t *testing.T) {
	for _, tc := range minMeetingRoomsTestCases {
		t.Run(fmt.Sprintf("minMeetingRooms(%v)", tc.intervals), func(t *testing.T) {
			actual := minMeetingRooms(tc.intervals)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkMinMeetingRooms(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for b.Loop() {
		for _, tc := range minMeetingRoomsTestCases {
			minMeetingRooms(tc.intervals)
		}
	}
}

func TestMinMeetingRoomsPriorityQueue(t *testing.T) {
	for _, tc := range minMeetingRoomsTestCases {
		t.Run(fmt.Sprintf("minMeetingRoomsPriorityQueue(%v)", tc.intervals), func(t *testing.T) {
			actual := minMeetingRoomsPriorityQueue(tc.intervals)
			if actual != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkMinMeetingRoomsPriorityQueue(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for b.Loop() {
		for _, tc := range minMeetingRoomsTestCases {
			minMeetingRoomsPriorityQueue(tc.intervals)
		}
	}
}
