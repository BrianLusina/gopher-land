package onlinestockspan

type StockSpanner struct {
	Stack [][]int
}

func Constructor() StockSpanner {
	return StockSpanner{[][]int{}}
}

func (this *StockSpanner) Next(price int) int {
	weight := 1

	for len(this.Stack) > 0 && price >= this.Stack[len(this.Stack)-1][0] {
		weight += this.Stack[len(this.Stack)-1][1]
		_ = this.Stack[len(this.Stack)-1]
		this.Stack = this.Stack[:len(this.Stack)-1]
	}

	this.Stack = append(this.Stack, []int{price, weight})

	return weight
}
