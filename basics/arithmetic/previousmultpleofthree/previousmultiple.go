package previousmultpleofthree

import (
	"fmt"
	"strconv"
)

func PrevMultOfThree(n int) interface{} {
	var resp interface{}
	if n%3 == 0 {
		resp = n
	} else if n > 3 {
		digits := fmt.Sprintf("%d", n)
		remainingDigits := digits[:len(digits)-1]
		if d, err := strconv.Atoi(remainingDigits); err == nil {
			resp = PrevMultOfThree(d)
		} else {
			resp = nil
		}
	} else {
		resp = nil
	}

	return resp
}

func PrevMultOfThree2(n int) interface{} {
	for i := n; i > 0; i /= 10 {
		if i%3 == 0 {
			return i
		}
	}
	return nil
}
