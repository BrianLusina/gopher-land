package deadfish

func Parse(data string) (output []int) {
	value := 0
	output = make([]int, 0)

	for _, d := range data {
		switch d {
		case 'i':
			value++
		case 's':
			value *= value
		case 'd':
			value--
		case 'o':
			output = append(output, value)
		}
	}

	return
}
