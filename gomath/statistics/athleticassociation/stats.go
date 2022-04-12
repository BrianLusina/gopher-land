package athleticassociation

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func formatSeconds(seconds int) string {
	return fmt.Sprintf("%02d|%02d|%02d", seconds/3600, (seconds%3600)/60, seconds%60)
}

func getMedian(timestamps []int) int {
	midIdx := len(timestamps) / 2
	if len(timestamps)%2 == 0 {
		return (timestamps[midIdx-1] + timestamps[midIdx]) / 2
	}
	return timestamps[midIdx]
}

func Stati(str string) string {
	if len(str) == 0 {
		return ""
	}

	results := strings.Split(str, ",")

	totalTime := 0
	var timestamps []int

	for _, result := range results {
		r := strings.Trim(result, " ")
		timestamp := strings.Split(r, "|")
		h, m, s := timestamp[0], timestamp[1], timestamp[2]

		hoursInSecs, err := strconv.Atoi(h)
		if err != nil {
			continue
		}
		hoursInSecs = hoursInSecs * 60 * 60

		minutesInSecs, err := strconv.Atoi(m)
		if err != nil {
			continue
		}
		minutesInSecs = minutesInSecs * 60

		seconds, err := strconv.Atoi(s)
		if err != nil {
			continue
		}

		totalTimeInSeconds := hoursInSecs + minutesInSecs + seconds

		totalTime += totalTimeInSeconds

		timestamps = append(timestamps, totalTimeInSeconds)
	}

	sort.Ints(timestamps)

	diff := timestamps[len(timestamps)-1] - timestamps[0]
	mean := totalTime / len(timestamps)
	median := getMedian(timestamps)

	return fmt.Sprintf("Range: %s Average: %s Median: %s", formatSeconds(diff), formatSeconds(mean), formatSeconds(median))
}
