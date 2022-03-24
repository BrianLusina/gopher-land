package interfaces

type Miner interface {
	Len() int
	ElemIx(ix int) interface{}
	Less(i, j int) bool
}

func Min(data Miner) interface{} {
	minimum := data.ElemIx(0)
	minIdx := 0
	for i := 1; i < data.Len(); i++ {
		if data.Less(i, minIdx) {
			minimum = data.ElemIx(i)
			minIdx = i
		}
	}

	return minimum
}

type IntArray []int

func (p IntArray) Len() int {
	return len(p)
}
func (p IntArray) ElemIx(ix int) interface{} {
	return p[ix]
}
func (p IntArray) Less(i, j int) bool {
	return p[i] < p[j]
}

type StringArray []string

func (p StringArray) Len() int {
	return len(p)
}
func (p StringArray) ElemIx(ix int) interface{} {
	return p[ix]
}
func (p StringArray) Less(i, j int) bool {
	return p[i] < p[j]
}
