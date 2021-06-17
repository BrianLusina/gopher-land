package orderedstream

type OrderedStream struct {
	Stream  []string
	Pointer int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		Pointer: 0,
		Stream:  make([]string, n),
	}
}

func (orderedStream *OrderedStream) Insert(idKey int, value string) []string {
	orderedStream.Stream[idKey-1] = value

	if orderedStream.Stream[orderedStream.Pointer] == "" {
		return []string{}
	}

	for idx, val := range orderedStream.Stream[orderedStream.Pointer:] {
		if val == "" {
			tmp := orderedStream.Pointer
			orderedStream.Pointer += idx
			return orderedStream.Stream[tmp:orderedStream.Pointer]
		}
	}
	return orderedStream.Stream[orderedStream.Pointer:]
}
