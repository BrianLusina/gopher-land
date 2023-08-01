package recentcounter

import "container/list"

type RecentCounter struct {
	slidingWindow *list.List
}

func New() RecentCounter {
	return RecentCounter{
		slidingWindow: list.New(),
	}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.slidingWindow.PushBack(t)

	for rc.slidingWindow.Front().Value.(int) < t-3000 {
		rc.slidingWindow.Remove(rc.slidingWindow.Front())
	}

	return rc.slidingWindow.Len()
}
