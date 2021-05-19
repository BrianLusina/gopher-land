package climbstairsmincost

import "math"

// func minCostClimbingStairs(cost []int) int {

// 	var minimumCost func(int) int

// 	miniumCost := func(i int) int {
// 		memo := map[int]int{}
// 		if i <= 1 {
// 			return 0
// 		}

// 		res, ok := memo[i]

// 		if ok {
// 			return res
// 		}

// 		// encounters nil pointer dereference? hmm
// 		downOne := cost[i-1] + minimumCost(i-1)
// 		downTwo := cost[i-2] + minimumCost(i-2)
// 		memo[i] = int(math.Min(float64(downOne), float64(downTwo)))
// 		return memo[i]
// 	}

// 	return miniumCost(len(cost))
// }

func minCostClimbingStairs(cost []int) int {
	/**
	Initialize two variables, downOne and downTwo, that represent the minimum cost to reach one step and two steps below the current step, respectively.
	We will start iteration from step 2, which means these variables will initially represent the minimum cost to reach steps 0 and 1,
	so we will initialize each of them to 0.
	*/
	downOne, downTwo := 0, 0

	/**
	Iterate over the array, again with 1 extra iteration at the end to treat the top floor as the final "step". At each iteration,
	simulate moving 1 step up. This means downOne will now refer to the current step, so apply our recurrence relation to update downOne.
	downTwo will be whatever downOne was prior to the update, so let's use a temporary variable to help with the update.
	**/
	for i := 2; i <= len(cost); i++ {
		temp := downOne
		downOne = int(math.Min(float64(downOne+cost[i-1]), float64(downTwo+cost[i-2])))
		downTwo = temp
	}

	// In the end, since we treated the top floor as a step, downOne will refer to the minimum cost to reach the top floor. Return downOne.
	return downOne
}
