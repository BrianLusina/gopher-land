package longesthappystrings

import (
	"container/heap"
	"strings"
)

type charCount struct {
	count int
	ch    string
}

// Max Head implementation for Go
type maxHeap []charCount

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(charCount))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func longestDiverseString(a, b, c int) string {
	if a == 0 && b == 0 && c == 0 {
		return ""
	}

	maxHeap := &maxHeap{}
	heap.Init(maxHeap)

	if a > 0 {
		heap.Push(maxHeap, charCount{count: a, ch: "a"})
	}
	if b > 0 {
		heap.Push(maxHeap, charCount{count: b, ch: "b"})
	}
	if c > 0 {
		heap.Push(maxHeap, charCount{count: c, ch: "c"})
	}

	result := []string{}
	for maxHeap.Len() > 0 {
		// Pop the firstCharacter with the highest remaining count/frequency
		firstCharacter := heap.Pop(maxHeap).(charCount)

		// Check if adding this character violates the "no three consecutive" rule
		if (len(result)) >= 2 && result[len(result)-1] == firstCharacter.ch && result[len(result)-2] == firstCharacter.ch {
			// If the rule is violated and no alternative character exists, break the loop
			if maxHeap.Len() == 0 {
				break
			}
			// Use the next most frequent character temporarily
			secondCharacter := heap.Pop(maxHeap).(charCount)
			// Add the alternative character to the result
			result = append(result, secondCharacter.ch)

			// Push the alternative character back with its updated count
			secondCharacter.count--
			if secondCharacter.count > 0 {
				heap.Push(maxHeap, secondCharacter)
			}
			// Push the original character back to the heap to try adding it later
			heap.Push(maxHeap, firstCharacter)
		} else {
			// If no violation, add the current character to the result
			firstCharacter.count--
			result = append(result, firstCharacter.ch)
			if firstCharacter.count > 0 {
				heap.Push(maxHeap, firstCharacter)
			}
		}
	}

	return strings.Join(result, "")
}
