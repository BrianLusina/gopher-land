package twosum

type TwoSum struct {
	counter map[int]int
}

func New() TwoSum {
	return TwoSum{
		counter: make(map[int]int),
	}
}

func (ts *TwoSum) Add(number int) {
	ts.counter[number]++
}

func (ts *TwoSum) Find(value int) bool {
	for num := range ts.counter {
		complement := value - num
		if _, ok := ts.counter[complement]; ok {
			if complement != num || ts.counter[num] > 1 {
				return true
			}
		}
	}
	return false
}
