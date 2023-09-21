// Package snapshotarray
//
// We wish to find the pos for the most recent value at the time we took the snapshot with the given snap_id, we are
//
// trying to find the rightmost index in history=histories[i] such that the snap_id at history[pos] is less or equal
// to the target snap_id (a[i][0] <= snap_id). This means that the feasible function is a[i][0] <= snap_id, whenever
// this is true, we must check the positions on its right to find the rightmost position that makes this condition hold.
//
// Complexity Analysis, Let n be the maximum number of calls and k = length.
//
// Time complexity: O(nlog(n)+k)
//   - We initialize histories with size k.
//   - In the worst-case scenario, the number of calls to get, set, and snap are all O(n).
//   - For each call to get(index, snap_id), we will perform a binary search over the list of records of nums[index].
//     Since a list contains at most O(n) records, a binary search takes O(logn) time on average. Thus it requires
//     O(nlogn) time.
//   - Each call to snap takes O(1) time.
//   - Each call to set(index, snap_id) appends a pair to the historical record of nums[index], which takes O(1) time,
//
// Space complexity: O(n+k)
//   - We initialize historyRecords with size k.
//   - We add one pair (snap_id, val) for each call to set, thus there are at most nnn pairs saved in histories
package snapshotarray

type SnapshotArray struct {
	snapId    int
	histories [][][]int
}

func New(length int) SnapshotArray {
	return SnapshotArray{
		snapId:    0,
		histories: make([][][]int, length),
	}
}

func (sa *SnapshotArray) Set(index int, val int) {
	if len(sa.histories[index]) == 0 || sa.histories[index][len(sa.histories[index])-1][0] != sa.snapId {
		sa.histories[index] = append(sa.histories[index], []int{sa.snapId, val})
		return
	}
	sa.histories[index][len(sa.histories[index])-1][1] = val
}

func (sa *SnapshotArray) Snap() int {
	sa.snapId++
	return sa.snapId - 1
}

func (sa *SnapshotArray) Get(index int, snapId int) int {
	data := sa.histories[index]

	if len(data) == 0 {
		return 0
	}

	if data[0][0] > snapId {
		return 0
	}

	if data[len(data)-1][0] <= snapId {
		return data[len(data)-1][1]
	}

	left, right := 0, len(data)-1

	for left <= right {
		mid := (left + right) / 2

		if data[mid][0] == snapId {
			return data[mid][1]
		}

		if data[mid][0] < snapId {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return data[right][1]
}
