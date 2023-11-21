package decodemessage

import "strconv"

func numberOfWaysToDecodeMessage(digits string) int {
	memo := map[int]int{}

	var dfs func(startIndex int) int
	dfs = func(startIndex int) int {
		if _, ok := memo[startIndex]; ok {
			return memo[startIndex]
		}

		if startIndex == len(digits) {
			return 1
		}

		ways := 0

		// can't decode string with leading 0
		if digits[startIndex] == '0' {
			return ways
		}

		// decode 1 digit
		ways += dfs(startIndex + 1)

		// decode 2 digits
		var digit int
		s := digits[startIndex : startIndex+2]

		if s == "" {
			digit = 0
		} else {
			d, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			digit = d
		}

		if startIndex+2 <= len(digits) && digit <= 26 {
			ways += dfs(startIndex + 2)
		}

		memo[startIndex] = ways
		return ways
	}

	return dfs(0)
}
