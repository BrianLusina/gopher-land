package happynumber

import "gopherland/math/utils"

// IsHappyNumber checks if an unsigned integer is a happy number
// A happy number is defined by the following process:
//   - Replace the number by the sum of the squares of its digits.
//   - Repeat the process until the number becomes 1 (happy number)
//     or a cycle is detected (not a happy number).
//
// This uses a set to store already seen numbers which consumes extra space. But since the numbers may not be too
// large, using a set might be straightforward and easy to implement
//
// While this approach works well for small numbers, we might have to perform several computations for larger numbers
// to get the required result. So, it might get infeasible for such cases. The time complexity of this approach is
// O(log(n)). The space complexity is O(log(n)) since we're using additional space to store our calculated sums.
func IsHappyNumber(n int) bool {
	seen := make(map[int]bool)

	for n != 1 {
		if seen[n] {
			return false
		}
		seen[n] = true

		sum := 0
		for n > 0 {
			digit := n % 10
			sum += digit * digit
			n /= 10
		}
		n = sum
	}

	return true
}

// IsHappyNumberFastAndSlow determines if a number n is a happy number.
// The function implements the Floyd's Tortoise and Hare algorithm to detect
// a cycle in the sequence of numbers produced by repeatedly summing the
// squares of the digits of a number. The algorithm uses two pointers, slow and fast,
// that move at different speeds through the sequence, eventually meeting at the start
// of the cycle. The algorithm assumes that the sequence contains a cycle and
// that the numbers in the sequence are indices that point to other numbers in the
// sequence.
// To determine whether a number is a happy number, it is iteratively replaced by the sum of the squares of its digits,
// forming a sequence of numbers. This sequence either converges to 1 (if the number is happy) or forms a cycle
// (if the number is not happy). We use the fast and slow pointers technique to detect such cycles efficiently.
// This technique involves advancing two pointers through the sequence at different speeds: one moving one step at a
// time and the other two at a time.
//
// The pointer moving slower is initialized to the given number, and the faster one starts at the sum of the squared
// digits of the given number. Then, in each subsequent iteration, the slow pointer updates to the sum of squared
// digits of itself, while the fast pointer advances two steps ahead: first by updating to the sum of squared digits
// of itself and then to the sum of squared digits of this recently calculated sum. If the number is happy, the fast
// pointer will eventually reach 1.
//
// However, if the number is not happy, indicating the presence of a cycle in the sequence, both pointers will
// eventually meet. This is because, in the non-cyclic part of the sequence, the distance between the pointers
// increases by one number in each iteration. Once both pointers enter the cyclic part, the faster pointer starts
// closing the gap on the slower pointer, decreasing the distance by one number in each iteration until they meet.
// This way, we can efficiently determine whether a number is a happy number or not.
//
//	The time complexity for this algorithm is O(log(n)), where n is the input number.
//
//	The worst case time complexity of this algorithm is given by the case of a non-happy number, since it gets stuck in
//	a cycle, whereas a happy number quickly converges to 1. Let’s first calculate the time complexity of the Sum Digits
//	function. Since we are calculating the sum of all digits in a number, the time complexity of this function is
//	O(log(n)), because the number of digits in the number n log10n.
//
// Time Complexity: O(log(n))
// Space Complexity: O(1)
func IsHappyNumberFastAndSlow(n int) bool {
	slow := n
	fast := n

	for {
		slow = utils.SumOfSquaredDigits(slow)
		fast = utils.SumOfSquaredDigits(utils.SumOfSquaredDigits(fast))

		if slow == fast {
			break
		}
	}

	return fast == 1
}
