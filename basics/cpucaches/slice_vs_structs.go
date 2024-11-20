package cpucaches

type Foo struct {
	a int
	b int
}

func sumFoo(foos []Foo) int {
	var total int
	for i := 0; i < len(foos); i++ {
		total += foos[i].a
	}
	return total
}

type Bar struct {
	a []int
	b []int
}

func sumBar(bar Bar) int {
	var total int
	for i := 0; i < len(bar.a); i++ {
		total += bar.a[i]
	}
	return total
}
