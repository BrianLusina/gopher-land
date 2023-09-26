package timemap

type pair struct {
	timestamp int
	value     string
}

type TimeMap struct {
	histories map[string][]pair
}

func New() TimeMap {
	histories := make(map[string][]pair)

	return TimeMap{
		histories: histories,
	}
}

func (tm *TimeMap) Set(key string, value string, timestamp int) {
	timeToValue := pair{timestamp: timestamp, value: value}

	if _, ok := tm.histories[key]; !ok {
		tm.histories[key] = make([]pair, 0)
	}

	tm.histories[key] = append(tm.histories[key], timeToValue)
}

func (tm *TimeMap) Get(key string, timestamp int) string {
	if _, ok := tm.histories[key]; !ok {
		return ""
	}

	keyHistoricalValues := tm.histories[key]

	if keyHistoricalValues[0].timestamp > timestamp {
		return ""
	}

	left, right, pos := 0, len(keyHistoricalValues)-1, -1

	for left <= right {
		mid := left + (right-left)/2

		timestampValueAtMid := keyHistoricalValues[mid]
		timestampForValue := timestampValueAtMid.timestamp

		if timestampForValue == timestamp {
			return timestampValueAtMid.value
		}

		if timestampForValue <= timestamp {
			left = mid + 1
			pos = mid
		} else {
			right = mid - 1
		}
	}

	if pos == -1 {
		return ""
	}

	return keyHistoricalValues[pos].value
}
