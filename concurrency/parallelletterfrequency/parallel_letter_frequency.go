package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	ch := make(chan FreqMap, 10)
	m := FreqMap{}

	var wg sync.WaitGroup
	wg.Add(len(l))

	for _, text := range l {
		go func(t string) {
			ch <- Frequency(t)
			wg.Done()
		}(text)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for freqMap := range ch {
		for k, v := range freqMap {
			m[k] += v
		}
	}
	return m
}
