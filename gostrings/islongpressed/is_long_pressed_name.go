package islongpressed

func isLongPressedName(name string, typed string) bool {
	i := 0

	for idx := range typed {
		if i < len(name) && name[i] == typed[idx] {
			i++
		} else if idx == 0 || typed[idx] != typed[idx-1] {
			return false
		}
	}

	return i == len(name)
}
