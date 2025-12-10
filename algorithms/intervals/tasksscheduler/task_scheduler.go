package tasksscheduler

import mathUtils "gopherland/math/utils"

func leastIntervalsMathematical(tasks []string, n int) int {
	totalTasks := len(tasks)
	if n == 0 || totalTasks <= 1 {
		return totalTasks
	}

	// get all frequencies of tasks
	taskCountMap := make(map[string]int)
	for _, task := range tasks {
		taskCount, ok := taskCountMap[task]
		if !ok {
			taskCountMap[task] = 1
		} else {
			taskCountMap[task] = taskCount + 1
		}
	}

	// get the maximum frequency
	maxFreq := 0
	for _, count := range taskCountMap {
		maxFreq = mathUtils.Max(count, maxFreq)
	}

	// calculate how many tasks share the maximum frequency
	// This is slightly more efficient than looping through the full dictionary
	//  because it stops checking once frequencies drop below max_freq

	maxFreqCount := 0
	for _, count := range taskCountMap {
		if count == maxFreq {
			maxFreqCount++
		}
	}

	result := (maxFreq-1)*(n+1) + maxFreqCount

	return mathUtils.Max(totalTasks, result)
}
