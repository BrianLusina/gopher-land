package meetingrooms

import (
	"container/heap"
	"slices"
)

// minMeetingRooms takes in an array of intervals and returns the minimum number of rooms required to hold all of the meetings.
// The intervals are given as pairs of start and end times where start <= end.
// The function returns the minimum number of rooms required to hold all of the meetings without any double booking.
// Complexity:
// Time Complexity: O(n*log(n)) as we sort the 2 lists
// Space Complexity: O(n) we store the meetings in a list
func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	startTimes := make([]int, len(intervals))
	endTimes := make([]int, len(intervals))

	for i, interval := range intervals {
		startTimes[i] = interval[0]
		endTimes[i] = interval[1]
	}

	slices.Sort(startTimes)
	slices.Sort(endTimes)

	usedRooms := 0
	j := 0

	for i := range startTimes {
		if startTimes[i] < endTimes[j] {
			usedRooms++
		} else {
			j++
		}
	}

	return usedRooms
}

// minMeetingRoomsPriorityQueue takes in an array of intervals and returns the minimum number of rooms required to hold all of the meetings.
// The intervals are given as pairs of start and end times where start <= end.
// The function returns the minimum number of rooms required to hold all of the meetings without any double booking.
// Complexity:
// Time Complexity: O(n*log(n)) as we sort the 2 lists
// Space Complexity: O(n) we store the meetings in a list
func minMeetingRoomsPriorityQueue(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// sort meetings in ascending order by start time
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	rooms := &priorityQueue{}
	heap.Init(rooms)
	heap.Push(rooms, intervals[0][1])

	for _, interval := range intervals[1:] {
		if rooms.Peek().(int) <= interval[0] {
			heap.Pop(rooms)
		}
		heap.Push(rooms, interval[1])
	}

	return rooms.Len()
}
