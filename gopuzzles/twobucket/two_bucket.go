package twobucket

import (
	"errors"
	"fmt"
)

type bucketNum int
type stepNum int

// buckets
const (
	bucketOne bucketNum = iota
	bucketTwo
)

// steps
const (
	emptyOne stepNum = iota
	emptyTwo
	fillOne
	fillTwo
	pourOneToTwo
	pourTwoToOne
)

const (
	FirstBucketName  = "one"
	SecondBucketName = "two"
)

type problem struct {
	capacity [2]int
	goal     int
	start    bucketNum
}

type state struct {
	level        [2]int
	previousStep stepNum
	numSteps     int
}

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	if err := validateParams(sizeBucketOne, sizeBucketTwo, goalAmount, startBucket); err != nil {
		return "", 0, 0, err
	}

	p := problem{
		capacity: [2]int{sizeBucketOne, sizeBucketTwo},
		goal:     goalAmount,
	}

	var currentState state

	if startBucket == FirstBucketName {
		p.start = bucketOne
		performStep(p, &currentState, fillOne)
	} else {
		p.start = bucketTwo
		performStep(p, &currentState, fillTwo)
	}

	if !isSolution(p, currentState) {
		currentState = findGoal(p, currentState)
	}
	switch {
	case currentState.level[bucketOne] == p.goal:
		return FirstBucketName, currentState.numSteps, currentState.level[bucketTwo], nil
	case currentState.level[bucketTwo] == p.goal:
		return SecondBucketName, currentState.numSteps, currentState.level[bucketOne], nil
	}
	return "", 0, 0, errors.New("no solution")
}

func validateParams(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) error {
	if sizeBucketOne <= 0 {
		return fmt.Errorf("sizeBucketOne %d is invalid", sizeBucketOne)
	}

	if sizeBucketTwo <= 0 {
		return fmt.Errorf("sizeBucketTwo %d is invalid", sizeBucketTwo)
	}

	if goalAmount <= 0 {
		return fmt.Errorf("goalAmount %d is invalid", goalAmount)
	}

	if startBucket != FirstBucketName && startBucket != SecondBucketName {
		return fmt.Errorf("startBucket %s is invalid", startBucket)
	}

	return nil
}

func performStep(p problem, currentState *state, step stepNum) {
	switch step {
	case emptyOne:
		currentState.level[bucketOne] = 0
	case emptyTwo:
		currentState.level[bucketTwo] = 0
	case fillOne:
		currentState.level[bucketOne] = p.capacity[bucketOne]
	case fillTwo:
		currentState.level[bucketTwo] = p.capacity[bucketTwo]
	case pourOneToTwo:
		pour(p, currentState, bucketOne, bucketTwo)
	case pourTwoToOne:
		pour(p, currentState, bucketTwo, bucketOne)
	}

	currentState.numSteps++
	currentState.previousStep = step
}

func pour(p problem, currentState *state, bOne, bTwo bucketNum) {
	amount := p.capacity[bTwo] - currentState.level[bTwo]
	if amount > currentState.level[bOne] {
		amount = currentState.level[bOne]
	}
	currentState.level[bTwo] += amount
	currentState.level[bOne] -= amount
}

func isSolution(p problem, s state) bool {
	return s.level[bucketOne] == p.goal || s.level[bucketTwo] == p.goal
}

func findGoal(p problem, s state) (g state) {
	searchList := make([]state, 1)
	searchList[0] = s

	// Use breadth-first search to find the goal, tracking any previously visited states.
	visited := map[[2]int]bool{}

	// Mark as already visited two invalid bucket levels: 0,0 and the reverse starting position.
	visited[[2]int{0, 0}] = true
	if p.start == bucketOne {
		visited[[2]int{0, p.capacity[bucketTwo]}] = true
	} else {
		visited[[2]int{p.capacity[bucketOne], 0}] = true
	}

	for len(searchList) != 0 {
		// Pop one item from the searchList each pass.
		current := searchList[0]
		searchList = searchList[1:]

		for _, x := range getPossibleSteps(p, current) {
			next := current
			performStep(p, &next, x)
			if isSolution(p, next) {
				return next
			}
			if !visited[next.level] {
				searchList = append(searchList, next)
				visited[next.level] = true
			}
		}
	}
	return state{}
}

func getPossibleSteps(p problem, s state) (list []stepNum) {
	for x := emptyOne; x <= pourTwoToOne; x++ {
		if canPerformStep(p, s, x) {
			list = append(list, x)
		}
	}
	return list
}

func canPerformStep(p problem, s state, x stepNum) bool {
	switch x {
	case emptyOne:
		if s.previousStep == fillOne || s.previousStep == pourOneToTwo {
			return false
		}
		return s.level[bucketOne] != 0
	case emptyTwo:
		if s.previousStep == fillTwo || s.previousStep == pourTwoToOne {
			return false
		}
		return s.level[bucketTwo] != 0
	case fillOne:
		if s.previousStep == emptyOne || s.level[bucketOne] == p.capacity[bucketOne] {
			return false
		}
		return true
	case fillTwo:
		if s.previousStep == emptyTwo || s.level[bucketTwo] == p.capacity[bucketTwo] {
			return false
		}
		return true
	case pourOneToTwo:
		return s.level[bucketOne] != 0 && s.level[bucketTwo] < p.capacity[bucketTwo]
	case pourTwoToOne:
		return s.level[bucketTwo] != 0 && s.level[bucketOne] < p.capacity[bucketOne]
	}
	return false
}
