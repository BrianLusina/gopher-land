package generateparenthesis

import "strings"

func generateParenthesis(n int) []string {
	stack := []string{}
	res := []string{}

	backtrack(0, 0, n, &stack, &res)
	return res
}

func backtrack(openCount, closedCount, n int, stack, res *[]string) {
	if openCount == closedCount && openCount == n && closedCount == n {
		*res = append(*res, strings.Join(*stack, ""))
		return
	}

	if openCount < n {
		*stack = append(*stack, "(")
		backtrack(openCount+1, closedCount, n, stack, res)
		_ = (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
	}

	if closedCount < openCount {
		*stack = append(*stack, ")")
		backtrack(openCount, closedCount+1, n, stack, res)
		_ = (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
	}
}
